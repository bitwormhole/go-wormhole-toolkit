package autoconfig

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// LoadAST 加载抽象语法树
func (inst *SourceFile) LoadAST() error {

	file := inst.Path
	fmt.Println("load AST from " + file.Path())
	src, err := file.GetIO().ReadText()
	if err != nil {
		return err
	}

	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		return err
	}

	// Print the AST.
	ast.Print(fset, f)
	return nil
}
