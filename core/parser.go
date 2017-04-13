package core

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/parser"
)

//parseObjectList will marshal our objects into warpcore format
func parseObjectList(list *ast.ObjectList) error {

	config := new(Config)

	if vars := list.Filter("resource"); len(vars.Items) > 0 {

		item := vars.Items[0]

		if len(item.Keys) > 0 {
			item = &ast.ObjectItem{
				Val: &ast.ObjectType{
					List: &ast.ObjectList{
						Items: []*ast.ObjectItem{item},
					},
				},
			}

			if _, ok := item.Val.(*ast.ObjectType); ok {
				log.Println("Ok")
			} else {
				return fmt.Errorf("Incorrect item type")
			}
			if err := hcl.DecodeObject(&config, item.Val); err != nil {
				return fmt.Errorf(
					"Error reading config: %s",
					err)
			}
		}
		spew.Dump(config)
	}

	return nil
}

//ParseFiles will decode HCL
func ParseFiles(files []string) (*ast.File, error) {
	if len(files) == 0 {
		return nil, fmt.Errorf("Empty file list")
	}
	for _, fname := range files {
		str, _ := ioutil.ReadFile(fname)

		parsed, err := parser.Parse(str)
		if err != nil {
			return nil, fmt.Errorf("Could not parse file")
		}

		list, ok := parsed.Node.(*ast.ObjectList)
		if !ok {
			return nil, fmt.Errorf("Could not read the object list")
		}

		parseObjectList(list)
	}

	return nil, fmt.Errorf("Nothing to parse")
}
