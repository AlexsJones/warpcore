package core

import (
    "os"
    "log"
    "path"
)

type context struct {

  Cwd string
}

func NewContext() *context {

  ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := path.Dir(ex)
  log.Println("Starting here ", exPath)
  return &context{ exPath }
}
