# lpcode

[![Go Report Card](https://goreportcard.com/badge/github.com/thorsphere/lpcode)](https://goreportcard.com/report/github.com/thorsphere/lpcode)
[![CodeFactor](https://www.codefactor.io/repository/github/thorsphere/lpcode/badge)](https://www.codefactor.io/repository/github/thorsphere/lpcode)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorsphere/lpcode)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorsphere/lpcode)](https://pkg.go.dev/mod/github.com/thorsphere/lpcode)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorsphere/lpcode)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorsphere/lpcode)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorsphere/lpcode)
![GitHub last commit](https://img.shields.io/github/last-commit/thorsphere/lpcode)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorsphere/lpcode)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorsphere/lpcode)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorsphere/lpcode)
![GitHub](https://img.shields.io/github/license/thorsphere/lpcode)

[Go](https://go.dev/) package for programmatic Go code generation that tries to keep it simple ([KISS principle](https://en.wikipedia.org/wiki/KISS_principle)). The package is designed primarily to support the [tserrgen](https://github.com/thorsphere/tserrgen) code generator.

- **Simple**: A simple and fluent API with method chaining
- **Formatted**: Automatic gofmt integration for clean output
- **Tested**: Well-tested with unit tests with high code coverage
- **Dependencies**: Minimal dependencies and only depends on [Go Standard Library](https://pkg.go.dev/std), [tsfio](https://pkg.go.dev/github.com/thorsphere/tsfio) and [tserr](https://pkg.go.dev/github.com/thorsphere/tserr)

## Usage

In the Go app, the package is imported with

```go
import "github.com/thorsphere/lpcode"
```

Codefile manages the lifecycle of a code file. StartFile initializes the file with a header template, WriteCode writes generated code content, and FinishFile appends a footer template and formats the result. Header and footer files (with suffixes ".header" and ".footer") are required and are searched for in the codefile directory or the current directory. Header and footer files allow you to include package declarations, imports, licenses, or other boilerplate that should appear at the beginning or end of generated files.

```go
func NewCodefile(dn tsfio.Directory, fn tsfio.Filename) (*Codefile, error)
func (cf *Codefile) StartFile() error
func (cf *Codefile) WriteCode(c *Code) error
func (cf *Codefile) FinishFile() error
```

There are helper functions for formatting, retrieving the code file filepath and for retrieving its contents as a string.

```go
func (cf *Codefile) Format() error
func (cf *Codefile) Filepath() tsfio.Filename
func (cf *Codefile) String() string
```

A Code contains source code as string. A builder for constructing Go source code strings using method chaining. The Code type offers a range of methods for generating common Go syntax. The stored source code is amended by calling its methods. The source code can be retrieved with String and formatted with Format.

```go
func NewCode() *Code
func (code *Code) LineComment(c string) *Code
func (code *Code) FuncEnd() *Code
func (code *Code) BlockEnd() *Code
func (code *Code) Call(n string) *Code
func (code *Code) ParamEndln() *Code
func (code *Code) ParamEnd() *Code
func (code *Code) Func1(a *Func1Args) *Code
func (code *Code) TypeStruct(n string) *Code
func (code *Code) VarSpec(a *VarSpecArgs) *Code
func (code *Code) List() *Code
func (code *Code) Listln() *Code
func (code *Code) SelField(a *SelArgs) *Code
func (code *Code) SelMethod(a *SelArgs) *Code
func (code *Code) If(a *IfArgs) *Code
func (code *Code) IfErr(a *IfErrArgs) *Code
func (code *Code) Return() *Code
func (code *Code) Addr() *Code
func (code *Code) Ident(n string) *Code
func (code *Code) Assignment(a *AssignmentArgs) *Code
func (code *Code) CompositeLit(LiteralType string) *Code
func (code *Code) ShortVarDecl(a *ShortVarDeclArgs) *Code
func (code *Code) KeyedElement(a *KeyedElementArgs) *Code
func (code *Code) Format() error
func (code *Code) String() string
```

The package includes support for generating test variable declarations through the TestVarDecl method. Test variables are predefined values with associated types and names that can be reused across multiple test cases to ensure consistency. The FindTestVar helper function allows you to retrieve a specific test variable by its type from an array of TestVar structs.

```go
func (code *Code) TestVarDecl(tv []TestVar) *Code
func FindTestVar(t string, tv []TestVar) (*TestVar, error)
```

## Example

```go
package main

import (
	"fmt"

	"github.com/thorsphere/lpcode"
)

func main() {
	c := lpcode.NewCode().
		TypeStruct("foo").
		VarSpec(&lpcode.VarSpecArgs{Ident: "A", Type: "int"}).
		VarSpec(&lpcode.VarSpecArgs{Ident: "B", Type: "string"}).BlockEnd().
		Func1(&lpcode.Func1Args{Name: "Example", Var: "x", Type: "int", Return: "error"}).
		IfErr(&lpcode.IfErrArgs{Method: "doSomething()", Operator: "!="}).Return().Ident("err").BlockEnd().
		If(&lpcode.IfArgs{ExprLeft: "x", Operator: ">", ExprRight: "10"}).
		CompositeLit("foo").
		KeyedElement(&lpcode.KeyedElementArgs{Key: "A", Elem: "1"}).
		KeyedElement(&lpcode.KeyedElementArgs{Key: `B`, Elem: `"hello"`}).
		BlockEnd().BlockEnd().Return().Ident("nil").FuncEnd()
	c.Format()
	fmt.Print(c)
}
```

[Run in Go Playground](https://go.dev/play/p/pR7uG4qtamD)

Output:

```go
type foo struct {
	A int
	B string
}

func Example(x int) error {
	if err := doSomething(); err != nil {
		return err
	}
	if x > 10 {
		foo{A: 1,
			B: "hello",
		}
	}
	return nil
}
```

## Limitations

It does not perform syntax validation; the generated code should be validated by go/format or other tools.

## Links

[Godoc](https://pkg.go.dev/github.com/thorsphere/lpcode)

[Go Report Card](https://goreportcard.com/report/github.com/thorsphere/lpcode)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Flpcode)

