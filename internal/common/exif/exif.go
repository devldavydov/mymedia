package exif

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	goexif "github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func RenameFile(dir, fileName string) error {
	f, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		return err
	}

	goexif.RegisterParsers(mknote.All...)

	var x *goexif.Exif
	if x, err = func() (*goexif.Exif, error) {
		defer f.Close()
		return goexif.Decode(f)
	}(); err != nil {
		return err
	}

	tm, _ := x.DateTime()
	dstPath := filepath.Join(dir, tm.Format("20060102_150405.jpg"))
	duplNumber := 0
	for isFileExists(dstPath) {
		duplNumber += 1
		dstPath = filepath.Join(dir, fmt.Sprintf("%s_%d.jpg", tm.Format("20060102_150405"), duplNumber))
	}

	return os.Rename(filepath.Join(dir, fileName), dstPath)
}

func isFileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
