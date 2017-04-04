package core

import (
	"log"
	"os"
	"path"
)

//Context data structure
type Context struct {
	Cwd string
}

//NewContext entrypoint
func NewContext(filepath string) *Context {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := path.Dir(ex)

	files := SearchFiles(filepath, ".warp")

	log.Println(files)

	_, err = ParseFiles(files)

	return &Context{exPath}
}
