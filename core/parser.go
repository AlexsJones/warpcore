package core

import (
	"fmt"
	"go/ast"
	"io/ioutil"
	"log"

	hclParser "github.com/hashicorp/hcl/hcl/parser"
)

//ParseFiles will decode HCL
func ParseFiles(files []string) (*ast.File, error) {
	for _, fname := range files {
		str, _ := ioutil.ReadFile(fname)

		log.Println(string(str))
		parsed, err := hclParser.Parse(str)
		if err != nil {
			return nil, fmt.Errorf("Could not parse file")
		}
		log.Println(parsed)
	}

	return nil, fmt.Errorf("Nothing to parse")
}
