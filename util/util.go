package util

import (
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// SnapLoader TODO:WriteDoc
type SnapLoader struct {
	folderPath  string
	fileinfos   []os.FileInfo
	filenames   []string
	fileendings []string
	egonodes    []string
}

// NewSnapLoader creates a new SnapLoader after validating input
func NewSnapLoader(path string) (*SnapLoader, error) {
	s := new(SnapLoader)
	// Constants
	s.fileendings = []string{"edges", "egofeat", "feat", "featnames"}
	// Validate path
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	s.folderPath = path
	log.WithFields(log.Fields{"DataPath": s.folderPath}).Info("NewSnapLoader")
	// Load egonodes
	s.fileinfos, err = ioutil.ReadDir(s.folderPath)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"Files": len(s.fileinfos)}).Info()
	// Extract egonodes and filenames
	s.filenames = make([]string, len(s.fileinfos))
	for i, f := range s.fileinfos {
		// egonode
		egonode := strings.Split(f.Name(), ".")[0]
		if len(s.egonodes) == 0 || egonode != s.egonodes[len(s.egonodes)-1] {
			s.egonodes = append(s.egonodes, egonode)
			log.WithFields(log.Fields{"Egonode": egonode}).Debug()
		}
		// filename
		// TODO: Are filenames really needed ?
		s.filenames[i] = f.Name()
		log.WithFields(log.Fields{"Filename": s.filenames[i]}).Debug()
	}
	log.WithField("Egonodes", s.egonodes).Info("NewSnapLoader Done")
	//
	return s, nil
}

func (s SnapLoader) Test() {
	log.Info("SnapLoader Test")
}
