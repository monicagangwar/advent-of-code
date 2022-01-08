package input

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func ReadInput(filepath string) []byte {
	patternFilePath := path.Join(path.Dir(filepath), "input.txt")

	content, err := ioutil.ReadFile(patternFilePath)
	if err != nil {
		log.Fatalf("unable to open pattern file due to error: %s", err)
	}
	return content
}

func GetFileMarker(filepath string) *os.File {
	file, err := os.Open(path.Join(path.Dir(filepath), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	return file
}