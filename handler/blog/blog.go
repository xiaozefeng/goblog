package blog

import (
	"io/ioutil"
	"strings"
)

func ListBlog(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, f := range files {
		if index := strings.LastIndex(f.Name(), `.`); index != -1 {
			result = append(result, f.Name()[:index])
		} else {
			result = append(result, f.Name())
		}

	}
	return result, nil
}
