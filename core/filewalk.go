package core
import (
  "path/filepath"
  "os"
  "log"
  "fmt")

func printFiles(ignoreDirs []string, outfiles []string, token string)filepath.WalkFunc {
  return func(path string, info os.FileInfo, err error) error {
    if err != nil {
               log.Print(err)
               return nil
           }
           if info.IsDir() {
               dir := filepath.Base(path)
               for _, d := range ignoreDirs {
                   if d == dir {
                       return filepath.SkipDir
                   }
               }
           }
           fmt.Println(path)
           return nil
       }
  }

func SearchFiles(token string) []string {

  ignoreDirs:= make([]string,0,5)

  fileSlice := make([]string,0,100)

  printFiles(ignoreDirs,fileSlice, token)

  return fileSlice
}
