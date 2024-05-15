package handlers

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
)

func OptionFormHandler(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Get string fields from the form
	bitString := form.Value["values"][0]

	// Get file from the form
	file, err := c.FormFile("files")
	if err != nil {
		return err
	}

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create destination file
	// dst, err := os.Create("./fileStorage/" + file.Filename)
	dst, err := os.Create("./fileStorage/" + file.Filename)
	if err != nil {
		log.Println("error creating file", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer dst.Close()

	// Copy the file content from source to destination
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Here, you can do whatever processing you need with the form data
	// For this example, let's just echo back the received data

	responseData := map[string]interface{}{
		"bit_string": bitString,
		"filename":   file.Filename,
	}

	return c.JSON(http.StatusOK, responseData)
}

func OptionFormHandler2(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Get string fields from the form
	bitString := form.Value["values"][0]

	// Get file from the form
	file, err := c.FormFile("files")
	if err != nil {
		return err
	}

	// Make folder from zip file name
	srcName := filenameWithoutExtension(file.Filename)
	err = os.Mkdir(filepath.Join("fileStorage", srcName), os.ModePerm)
	if err != nil {
		return err
	}

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create a temporary file to store the uploaded zip file
	tempFile, err := os.Create("./fileStorage/uploaded.zip")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// defer os.Remove(tempFile.Name()) // Remove the temporary file after function execution
	defer tempFile.Close()

	// Copy the file data to the temporary file
	_, err = io.Copy(tempFile, src)
	if err != nil {
		return err
	}

	// Open the temporary zip file
	r, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return err
	}
	defer r.Close()

	// Extract each file in the zip archive
	for _, f := range r.File {
		// Open the file in the zip archive
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Create the file to store the extracted content
		extractedFile, err := os.Create("./fileStorage/" + srcName + "/" + f.Name)
		if err != nil {
			return err
		}
		defer extractedFile.Close()

		// Copy the content from the zip file to the extracted file
		_, err = io.Copy(extractedFile, rc)
		if err != nil {
			return err
		}
	}

	responseData := map[string]interface{}{
		"bit_string": bitString,
		"filename":   file.Filename,
	}

	return c.JSON(http.StatusOK, responseData)
}

func filenameWithoutExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
