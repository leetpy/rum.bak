package util

import (
	"io/ioutil"
	"path"
	"strings"
)

func ListDir(dirPath string, suffix []string) (files []string, err error) {
	files = make([]string, 0, 20)
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}

		if suffix != nil {
			s := strings.Split(fi.Name(), ".")
			if len(s) <= 1 {
				continue
			}
			for _, v := range suffix {
				if strings.ToUpper(v) == strings.ToUpper(s[len(s)-1]) {
					files = append(files, path.Join(dirPath, fi.Name()))
					break
				}
			}
		} else {
			files = append(files, path.Join(dirPath, fi.Name()))
		}
	}
	return files, nil
}
