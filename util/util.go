package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Test() {
	fmt.Println("Test")
}

type SnapLoader struct {
	folderPath string
}

// NewSnapLoader creates a new SnapLoader after validating input
func NewSnapLoader(path string) (*SnapLoader, error) {
	s := new(SnapLoader)
	// Validate path
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	s.folderPath = path
	// Load egonodes
	files, err := ioutil.ReadDir(s.folderPath)
	if err != nil {
		return nil, err
	}

	fmt.Println("Files: " + str(len(files)))
	return s, nil
}

func (sn SnapLoader) Test() {
	fmt.Println("SnapLoader test")
}
