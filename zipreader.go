package odfgo

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
)

type zipreader struct {
	fileName string
}

func (zr zipreader) searchZip(filePath string) []byte {

	var content []byte

	// Open a zip archive for reading.
	r, err := zip.OpenReader(zr.fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {

		if f.Name == filePath && !f.FileInfo().IsDir() {

			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			content := make([]byte, f.FileInfo().Size())

			_, err = io.ReadFull(rc, content)
			if err != nil {
				log.Fatal(err)
			}

			rc.Close()

		}
	}

	r.Close()

	return content

}

func (zr zipreader) readZip() {

	// Open a zip archive for reading.
	r, err := zip.OpenReader(zr.fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
	}

	r.Close()

}