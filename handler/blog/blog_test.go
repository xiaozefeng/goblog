package blog

import (
	"fmt"
	"strings"
	"testing"
)

func TestListBlog(t *testing.T) {
	fileNames, err := ListBlog("../../assets/static/")
	if err != nil {
		t.Error(err)
	}

	for _, f := range fileNames {
		fmt.Println(f)
	}

}

func TestStr(t *testing.T) {
	str := "index.html"
	index := strings.Index(str, `.`)

	r := str[:index]
	fmt.Println(r)
}
