package odfgo

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type xmlreader struct {
	fileName string
}

func (zr xmlreader) searchZip(filePath string) string {

	var content string = ""

	// Open a zip archive for reading.
	r, err := zip.OpenReader(zr.fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {

		if f.Name == filePath {

			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			content, err := ioutil.ReadFile(rc)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.CopyN(os.Stdout, rc, 68)
			if err != nil {
				log.Fatal(err)
			}

			rc.Close()

		}
	}

	return content

}

func (zr xmlreader) readZip() {

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

}