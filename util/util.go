package util

import (
	"io/ioutil"
	"path"
	"strings"
)

func ListDir(dirPath, suffix string) (files []string, err error) {
	files = make([]string, 0, 20)
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}

		if suffix != "" {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				files = append(files, path.Join(dirPath, fi.Name()))
			}
		} else {
			files = append(files, path.Join(dirPath, fi.Name()))
		}
	}
	return files, nil
}
