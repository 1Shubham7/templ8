package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(zipPath, fileUrl string) error {
	response, err := http.Get(fileUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// create file
	file, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// write the response body to file
	_, err = io.Copy(file, response.Body)
	return err
}

func unzip(source, destination string) ([]string, error) {
	var filenames []string

	// OpenReader will open the Zip file and return a ReadCloser.
	reader, err := zip.OpenReader(source)
	if err != nil {
		return filenames, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		// will use this later on
		filePath := filepath.Join(destination, file.Name)
		fmt.Println("The filepath is: ", filePath)

		filenames = append(filenames, filePath)

		// if it itself is a dir, then make folder
		if file.FileInfo().IsDir() {
			// os.ModePern is 0777
			os.Mkdir(filePath, os.ModePerm)
			continue
		}

		// for all dirs inside zip file
		// filepath.Dir() does this:
		// /home/user/documents/file.txt -> /home/user/documents
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return filenames, err
		}

		outputFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		// os.O_WRONLY: Open the file in write-only mode.
		// os.O_CREATE: Create the file if it does not exist.
		// os.O_TRUNC: Truncate (empty) the file if it already exists.
		// these flags can be combined using bitwise OR '|'.
		// file.Mode() ~ 0755

		if err != nil {
			return filenames, err
		}

		readCloser, err := file.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outputFile, readCloser)

		// not using defer because we want to close this file before we go to next file in loop
		outputFile.Close()
		readCloser.Close()

		if err != nil {
			return filenames, err
		}
	}

	return filenames, nil
}
