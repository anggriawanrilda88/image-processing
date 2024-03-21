# Ubernap Test Documentation

This Image Processing application is created to complete the technical test at Ubersnap Company. The code architecture design in this repository follows the clean architecture using the principles of DDD (Domain Driven Design), which places the business logic in the domain layer. Therefore, all business logic related to image processing can be found in `/src/domain/entities/image_processing_entity.go`.

## APIs

This repository provides several APIs:

Converts an image from PNG format to JPG format.

**Example Request:**
```bash
curl --location 'http://192.168.190.3:8011/api/v1/images/convert-png-to-jpg' \
--form 'image=@"/C:/Users/anggr/Desktop/git.png"'
```

Resizes an image to the specified dimensions.

**Example Request:**
```bash
curl --location 'http://192.168.190.3:8011/api/v1/images/resize-image' \
--form 'image=@"/C:/Users/anggr/Desktop/git.png"' \
--form 'width="100"' \
--form 'height="100"'
```

Compresses an image while maintaining reasonable quality.

**Example Request:**
```bash
curl --location 'http://192.168.190.3:8011/api/v1/images/compress-image' \
--form 'image=@"/C:/Users/anggr/Desktop/linkedin.png"' \
--form 'imageQuality="high"'
```

Downloads the converted image using the provided link.

**Example Request:**
```bash
curl --location '192.168.190.3:8011/api/v1/images/download/65302fb4-59bf-4762-8e41-5cb6116a.jpg'
```

## Quick Setup
To run this Golang application, follow these steps:

<br>Prepare the .env file according to the example in the .env.example file and adjust the variables according to the condition of your local computer.
<br>Run the command `go mod tidy` to download the required external modules.
<br>Run the command `make setup-tools` to download the necessary files needed for development auto run such as the 'air' binary.
<br>Install the OpenCV library on your local machine. To install it, follow the steps on the following link: [OpenCV Installation Guide](https://gocv.io/).
<br>To start the application, run the command `make run-dev` and the application will run with auto reload when there are file changes.
<br>To run unit testing, use the command `go test -cover github.com/image-processing/src/domain/entities`. Unit testing is specifically focused only on the entities package because the business logic functions are placed in this file.