package util

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

type SnapLoader struct {
	folderPath string
	fileinfos  []os.FileInfo
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
	log.WithFields(log.Fields{"DataPath": s.folderPath}).Info()
	// Load egonodes
	fileinfos, err := ioutil.ReadDir(s.folderPath)
	if err != nil {
		return nil, err
	}
	s.fileinfos = fileinfos
	log.WithFields(log.Fields{"Files": len(s.fileinfos)}).Debug()
	//
	return s, nil
}

func (sn SnapLoader) Test() {
	log.Info("SnapLoader Test")
}
