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

A Codefile represents a code file with an interface that provides methods to start, write and finish the code file. A utility for managing the lifecycle of code files.

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

A Code contains the source code as string. A builder for constructing Go source code strings using method chaining. The Code type offers a range of methods for generating common Go syntax. The stored source code is amended by using its methods. The source code can be retrieved with String and formatted with Format.

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
func (code *Code) Testvariables(t *Testvars) *Code
func (code *Code) Format() error
func (code *Code) String() string
```

## Example

## Links

[Godoc](https://pkg.go.dev/github.com/thorsphere/lpcode)

[Go Report Card](https://goreportcard.com/report/github.com/thorsphere/lpcode)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Flpcode)

