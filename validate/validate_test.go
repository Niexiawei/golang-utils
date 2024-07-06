package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func init() {
	v := validator.New(validator.WithRequiredStructEnabled())
	Setup(v)
}

func TestSetupValid(t *testing.T) {
	fmt.Println(validatorInstance)
}

type PointTest struct {
	Location string `json:"location" validate:"required,location"`
}

func TestLocationValidator(t *testing.T) {
	param := PointTest{
		Location: "11836",
	}
	err := validatorInstance.Struct(param)
	if err != nil {
		t.Error(FormatError(err))
	}
}

func TestScan(t *testing.T) {
	filePath := "./method/In_validator.go"

	// 创建一个标志集
	fset := token.NewFileSet()

	// 解析文件
	f, parseErr := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if parseErr != nil {
		log.Fatalf("Error parsing %s: %v", filePath, parseErr)
	}

	// 遍历文件中的所有声明
	for _, decl := range f.Decls {
		//fmt.Println(decl)
		// 判断是否是常量声明
		ff, ok := decl.(*ast.GenDecl)
		if !ok || ff.Tok != token.TYPE {
			continue
		}
	}
}
