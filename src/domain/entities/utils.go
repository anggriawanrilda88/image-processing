package entities

import (
	"image/jpeg"
	"io"
	"os"

	"gocv.io/x/gocv"
)

var osCreateInput = os.Create
var osCreateOutput = os.Create
var osOpen = os.Open

var ioCopy = io.Copy
var ioReadAll = io.ReadAll

var gocvIMRead = gocv.IMRead
var gocvIMWrite = gocv.IMWrite

var jpegEncode = jpeg.Encode
