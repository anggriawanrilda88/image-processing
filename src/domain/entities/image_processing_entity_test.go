package entities

import (
	"errors"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gocv.io/x/gocv"
)

func createDummyFile(fileName string) (*multipart.FileHeader, *os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, nil, err
	}

	fileHeader := &multipart.FileHeader{
		Filename: fileInfo.Name(),
		Size:     fileInfo.Size(),
	}

	return fileHeader, file, nil
}

func Test_ImageProcessing(t *testing.T) {
	add := ImageProcessing{
		outputPath:  "sds",
		fileContent: []byte{},
	}

	assert.Equal(t, add.GetOutputPath(), "sds")
	assert.Equal(t, add.GetFileContent(), []byte{})
}

func TestCreateConverPNGtoJPEG_Success(t *testing.T) {
	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = CreateConverPNGtoJPEG(file, fileHeader, "http://example.com")
	if err != nil {
		t.Errorf("Error while creating converting PNG to JPEG: %v", err)
	}
}

func TestCreateConverPNGtoJPEG_Err_Extension(t *testing.T) {
	fileHeader, file, err := createDummyFile("../../../mock/test.jpg")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = CreateConverPNGtoJPEG(file, fileHeader, "http://example.com")
	assert.Equal(t, "invalid file format. only PNG files are supported", err.Error())
}

func TestCreateConverPNGtoJPEG_Err_CreateInputFile(t *testing.T) {
	osCreateInput = func(_ string) (*os.File, error) {
		return nil, errors.New("file not found")
	}

	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = CreateConverPNGtoJPEG(file, fileHeader, "http://example.com")
	assert.Equal(t, "failed to save uploaded file", err.Error())
}

func TestCreateConverPNGtoJPEG_Err_CreateFileCopy(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = func(_ io.Writer, _ io.Reader) (written int64, err error) {
		return 0, errors.New("errors")
	}

	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = CreateConverPNGtoJPEG(file, fileHeader, "http://example.com")
	assert.Equal(t, "failed to save uploaded file", err.Error())
}

func TestCreateConverPNGtoJPEG_Err_ReadFile(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy

	gocvIMRead = func(_ string, _ gocv.IMReadFlag) gocv.Mat {
		return gocv.IMRead("test-error.jpg", gocv.IMWritePxmBinary)
	}

	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = CreateConverPNGtoJPEG(file, fileHeader, "http://example.com")
	assert.Error(t, err)
}

func TestCreateConverPNGtoJPEG_Err_CreateOutputFile(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = func(_ string) (*os.File, error) {
		return nil, errors.New("file not found")
	}

	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = CreateConverPNGtoJPEG(file, fileHeader, "http://example.com")
	assert.Equal(t, "failed to save uploaded file", err.Error())
}

func TestCreateConverPNGtoJPEG_Err_JPG_Encode(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create
	jpegEncode = func(_ io.Writer, _ image.Image, _ *jpeg.Options) error {
		return errors.New("error encode")
	}

	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = CreateConverPNGtoJPEG(file, fileHeader, "http://example.com")
	assert.Equal(t, "error encode", err.Error())
}

func TestResizeSpecificImage_Success(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ResizeSpecificImage(file, fileHeader, "100", "100", "http://example.com")
	if err != nil {
		t.Errorf("Error while creating converting PNG to JPEG: %v", err)
	}
}

func TestResizeSpecificImage_Err_Width_Not_Int(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ResizeSpecificImage(file, fileHeader, "xxx", "100", "http://example.com")
	assert.Equal(t, "invalid width parameter", err.Error())
}

func TestResizeSpecificImage_Err_Height_Not_Int(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ResizeSpecificImage(file, fileHeader, "100", "xxx", "http://example.com")
	assert.Equal(t, "invalid height parameter", err.Error())
}

func TestResizeSpecificImage_Err_CreateFile(t *testing.T) {
	osCreateOutput = func(_ string) (*os.File, error) {
		return nil, errors.New("file not found")
	}
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ResizeSpecificImage(file, fileHeader, "100", "100", "http://example.com")
	assert.Equal(t, "failed to save uploaded file", err.Error())
}

func TestResizeSpecificImage_Err_CopyFile(t *testing.T) {
	osCreateOutput = os.Create
	ioCopy = func(_ io.Writer, _ io.Reader) (written int64, err error) {
		return 0, errors.New("error copy")
	}
	gocvIMRead = gocv.IMRead

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ResizeSpecificImage(file, fileHeader, "100", "100", "http://example.com")
	assert.Equal(t, "failed to save uploaded file", err.Error())
}

func TestResizeSpecificImage_Err_ReadFile(t *testing.T) {
	osCreateOutput = os.Create
	ioCopy = io.Copy
	gocvIMRead = func(_ string, _ gocv.IMReadFlag) gocv.Mat {
		return gocv.IMRead("test-error.jpg", gocv.IMWritePxmBinary)
	}

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ResizeSpecificImage(file, fileHeader, "100", "100", "http://example.com")
	assert.Error(t, err)
}

func TestResizeSpecificImage_Err_WriteFile(t *testing.T) {
	osCreateOutput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	gocvIMWrite = func(_ string, _ gocv.Mat) bool {
		return false
	}

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ResizeSpecificImage(file, fileHeader, "100", "100", "http://example.com")
	assert.Error(t, err)
}

func TestImageCompress_Success_LowQuality(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create
	jpegEncode = jpeg.Encode

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "low", "http://example.com")
	if err != nil {
		t.Errorf("Error while creating converting PNG to JPEG: %v", err)
	}
}

func TestImageCompress_Success_MediumQuality(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create
	jpegEncode = jpeg.Encode

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "medium", "http://example.com")
	if err != nil {
		t.Errorf("Error while creating converting PNG to JPEG: %v", err)
	}
}

func TestImageCompress_Success_HighQuality(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create
	jpegEncode = jpeg.Encode

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "high", "http://example.com")
	if err != nil {
		t.Errorf("Error while creating converting PNG to JPEG: %v", err)
	}
}

func TestImageCompress_Err_ReqQuality(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create
	jpegEncode = jpeg.Encode

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "xxx", "http://example.com")
	assert.Equal(t, "invalid quality type. fill with this type 'high', 'medium', 'low'", err.Error())
}

func TestImageCompress_Err_CreateFile(t *testing.T) {
	osCreateInput = func(_ string) (*os.File, error) {
		return nil, errors.New("error create")
	}
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create
	jpegEncode = jpeg.Encode

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "low", "http://example.com")
	assert.Equal(t, "failed to save uploaded file", err.Error())
}

func TestImageCompress_Err_CopyFile(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = func(_ io.Writer, _ io.Reader) (written int64, err error) {
		return 0, errors.New("errpr copy")
	}
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create
	jpegEncode = jpeg.Encode

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "low", "http://example.com")
	assert.Equal(t, "failed to save uploaded file", err.Error())
}

func TestImageCompress_Err_ReadFile(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = func(_ string, _ gocv.IMReadFlag) gocv.Mat {
		return gocv.IMRead("test-error.jpg", gocv.IMWritePxmBinary)
	}
	osCreateOutput = os.Create
	jpegEncode = jpeg.Encode

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "low", "http://example.com")
	assert.Error(t, err)
}

func TestImageCompress_Err_OutputCreateFile(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = func(_ string) (*os.File, error) {
		return nil, errors.New("error output create")
	}
	jpegEncode = jpeg.Encode

	imagePathFolder = "../../../tmp"
	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "low", "http://example.com")
	assert.Equal(t, "failed to save uploaded file", err.Error())
}

func TestImageCompress_Err_JPG_Encode(t *testing.T) {
	osCreateInput = os.Create
	ioCopy = io.Copy
	gocvIMRead = gocv.IMRead
	osCreateOutput = os.Create
	jpegEncode = func(_ io.Writer, _ image.Image, _ *jpeg.Options) error {
		return errors.New("error encode")
	}

	fileHeader, file, err := createDummyFile("../../../mock/test.png")
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer file.Close()

	_, err = ImageCompress(file, fileHeader, "low", "http://example.com")
	assert.Equal(t, "error encode", err.Error())
}

func TestDownload_Success(t *testing.T) {
	imagePathFolder = "../../../mock"
	_, err := Download("test.jpg")
	if err != nil {
		t.Errorf("Error while creating converting PNG to JPEG: %v", err)
	}
}

func TestDownload_Err_OpenFile(t *testing.T) {
	imagePathFolder = "../../../mock"
	osOpen = func(_ string) (*os.File, error) {
		return nil, errors.New("error open")
	}

	_, err := Download("test.jpg")
	assert.Equal(t, "error open", err.Error())
}

func TestDownload_Err_ReadAllFile(t *testing.T) {
	imagePathFolder = "../../../mock"
	osOpen = os.Open
	ioReadAll = func(_ io.Reader) ([]byte, error) {
		return nil, errors.New("error read all")
	}

	_, err := Download("test.jpg")
	assert.Equal(t, "error read all", err.Error())
}
