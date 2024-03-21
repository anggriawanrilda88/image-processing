package entities

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gocv.io/x/gocv"
)

type ImageProcessing struct {
	outputPath  string
	fileContent []byte
}

var imagePathFolder = "./tmp"

const apiPathV1 = "/api/v1/images/download/"
const imageQuality = 100

func (v *ImageProcessing) GetOutputPath() string  { return v.outputPath }
func (v *ImageProcessing) GetFileContent() []byte { return v.fileContent }

func CreateConverPNGtoJPEG(file multipart.File, header *multipart.FileHeader, host string) (*ImageProcessing, error) {
	// read file extension
	ext := filepath.Ext(header.Filename)
	if ext != ".png" {
		return nil, errors.New("invalid file format. only PNG files are supported")
	}

	// create directory
	tempDir := imagePathFolder
	os.Mkdir(tempDir, os.ModePerm)

	// create input output file
	randomName := uuid.New().String()[:32]
	outputFileName := randomName + ".jpg"
	inputPath := filepath.Join(tempDir, outputFileName)
	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".jpg"

	// save file to local disk
	input, err := osCreateInput(inputPath)
	if err != nil {
		return nil, errors.New("failed to save uploaded file")
	}
	defer input.Close()
	_, err = ioCopy(input, file)
	if err != nil {
		return nil, errors.New("failed to save uploaded file")
	}

	// read from png file
	img := gocvIMRead(outputPath, gocv.IMReadColor)
	if img.Empty() {
		return nil, fmt.Errorf("failed to read image: %s", outputPath)
	}
	defer img.Close()

	// save for change jpg file to local disk
	output, err := osCreateOutput(outputPath)
	if err != nil {
		return nil, errors.New("failed to save uploaded file")
	}
	defer output.Close()

	// save as jpeg
	jpegOptions := &jpeg.Options{Quality: imageQuality} // set image quality
	image, _ := img.ToImage()
	if err := jpegEncode(output, image, jpegOptions); err != nil {
		return nil, err
	}

	// set host to jpg image download
	downloadLink := host + apiPathV1 + outputPath
	downloadLink = strings.Replace(downloadLink, "/tmp", "", 1)

	return &ImageProcessing{
		outputPath: downloadLink,
	}, nil
}

func ResizeSpecificImage(
	file multipart.File,
	header *multipart.FileHeader,
	widthStr string,
	heightStr string,
	host string,
) (*ImageProcessing, error) {
	// read file extension
	ext := filepath.Ext(header.Filename)

	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return nil, errors.New("invalid width parameter")
	}
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return nil, errors.New("invalid height parameter")
	}

	// create directory
	tempDir := imagePathFolder
	os.Mkdir(tempDir, os.ModePerm)

	// create input output file
	randomName := uuid.New().String()[:32]
	outputFileName := randomName + ext
	inputPath := filepath.Join(tempDir, outputFileName)
	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ext

	// save file to local disk
	out, err := osCreateOutput(outputPath)
	if err != nil {
		return nil, errors.New("failed to save uploaded file")
	}
	defer out.Close()
	_, err = ioCopy(out, file)
	if err != nil {
		return nil, errors.New("failed to save uploaded file")
	}

	// read from png file
	img := gocvIMRead(outputPath, gocv.IMWritePxmBinary)
	if img.Empty() {
		return nil, fmt.Errorf("failed to read image: %s", outputPath)
	}
	defer img.Close()

	// set new size
	resized := gocv.NewMat()
	gocv.Resize(img, &resized, image.Point{X: width, Y: height}, 0, 0, gocv.InterpolationArea)

	// write image with new size
	if !gocvIMWrite(outputPath, resized) {
		return nil, fmt.Errorf("failed to write resized image: %s", outputPath)
	}

	// set host to jpg image download
	downloadLink := host + apiPathV1 + outputPath
	downloadLink = strings.Replace(downloadLink, "/tmp", "", 1)

	return &ImageProcessing{
		outputPath: downloadLink,
	}, nil
}

func ImageCompress(file multipart.File, header *multipart.FileHeader, imageQuality string, host string) (*ImageProcessing, error) {
	quality := 30
	if imageQuality == "high" {
		quality = 100
	} else if imageQuality == "medium" {
		quality = 55
	} else if imageQuality == "low" {
		quality = 10
	} else {
		return nil, errors.New("invalid quality type. fill with this type 'high', 'medium', 'low'")
	}

	// read file extension
	ext := filepath.Ext(header.Filename)

	// create directory
	tempDir := imagePathFolder
	os.Mkdir(tempDir, os.ModePerm)

	// create input output file
	randomName := uuid.New().String()[:32]
	outputFileName := randomName + ext
	inputPath := filepath.Join(tempDir, outputFileName)
	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ext

	// save file to local disk
	input, err := osCreateInput(inputPath)
	if err != nil {
		return nil, errors.New("failed to save uploaded file")
	}
	defer input.Close()
	_, err = ioCopy(input, file)
	if err != nil {
		return nil, errors.New("failed to save uploaded file")
	}

	// read from png file
	img := gocvIMRead(outputPath, gocv.IMReadColor)
	if img.Empty() {
		return nil, fmt.Errorf("failed to read image: %s", outputPath)
	}
	defer img.Close()

	// save file to local disk
	output, err := osCreateOutput(outputPath)
	if err != nil {
		return nil, errors.New("failed to save uploaded file")
	}
	defer output.Close()

	// save as jpeg
	jpegOptions := &jpeg.Options{Quality: quality} // set image quality
	image, _ := img.ToImage()
	if err := jpegEncode(output, image, jpegOptions); err != nil {
		return nil, err
	}

	// set host to jpg image download
	downloadLink := host + apiPathV1 + outputPath
	downloadLink = strings.Replace(downloadLink, "/tmp", "", 1)

	return &ImageProcessing{
		outputPath: downloadLink,
	}, nil
}

func Download(image string) (*ImageProcessing, error) {
	tempDir := imagePathFolder
	filePath := filepath.Join(tempDir, image)

	// read file
	file, err := osOpen(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read file content
	fileContent, err := ioReadAll(file)
	if err != nil {
		return nil, err
	}

	return &ImageProcessing{
		fileContent: fileContent,
	}, nil
}
