// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode

// Import Go standard library packages and tserr
import (
	"fmt"       // fmt
	"go/format" // format

	"github.com/thorsphere/tserr" // tserr
)

// Code contains the source code as string. The stored source code is amended by using
// its methods. The source code can be retrieved with String and formatted
// with Format.
type Code struct {
	c string // the source code
}

// NewCode returns a pointer to a new Code instance.
// The source code in the new Code instance is empty.
func NewCode() *Code {
	// Return a new Code instance
	return &Code{}
}

// String returns the source code in code as string.
// It returns an empty string if code is nil.
func (code *Code) String() string {
	// Return an empty string if code is nil
	if code == nil {
		// Return an empty string
		return ""
	}
	// Return the source code as string
	return code.c
}

// LineComment adds a line comment and a new line to code: // c\n. The comment is provided by argument c.
func (code *Code) LineComment(c string) *Code {
	// Return nil if code is nil
	if code == nil {
		return nil
	}
	// Add a line comment and a new line to code
	code.c += fmt.Sprintf("// %v\n", c)
	// Return code
	return code
}

// FuncEnd adds a block end and two new lines to code: }\n\n.
func (code *Code) FuncEnd() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a block end and two new lines to code
	code.c += "}\n\n"
	// Return code
	return code
}

// BlockEnd adds a block ending to code: }\n.
func (code *Code) BlockEnd() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a block ending to code
	code.c += "}\n"
	// Return code
	return code
}

// Call adds a function call to code: n(. The function name is
// provided by n.
func (code *Code) Call(n string) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a function call to code
	code.c += fmt.Sprintf("%v(", n)
	// Return code
	return code
}

// ParamEndln adds a parameters ending and a new line to code: )\n.
func (code *Code) ParamEndln() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a parameters ending and a new line to code
	code.c += ")\n"
	// Return code
	return code
}

// ParamEnd adds a parameters ending to code: ).
func (code *Code) ParamEnd() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add parameters ending to code
	code.c += ")"
	// Return code
	return code
}

// Func1Args contains the configuration for the generation of a function declaration with Func1.
type Func1Args struct {
	Name, Var, Type, Return string
}

// Func1 adds a function declaration with one parameter to code: func Name(Var Type) Return {\n.
// The function name, parameter variable, parameter type and return type are provided by a.
func (code *Code) Func1(a *Func1Args) *Code {
	// Return nil in case code is nil.
	if code == nil {
		return nil
	}
	// Return nil in case a is nil.
	if a == nil {
		return nil
	}
	// Add a function declaration with one parameter to code.
	code.c += fmt.Sprintf("func %v(%v %v) %v {\n", a.Name, a.Var, a.Type, a.Return)
	// Return code
	return code
}

// TypeStruct adds a type declaration for a struct type to code: type n struct {\n.
// The name of the type is provided with n.
func (code *Code) TypeStruct(n string) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a type declaration for a struct type to code
	code.c += fmt.Sprintf("type %v struct {\n", n)
	// Return code
	return code
}

// VarSpecArgs contains the identifier Ident and type Type to generate
// a variable specification with VarSpec
type VarSpecArgs struct {
	Ident, Type string
}

// VarSpec adds a variable specification to code: Ident Type\n. The identifier and type
// is provided by a.
func (code *Code) VarSpec(a *VarSpecArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a variable specification to code
	code.c += fmt.Sprintf("%v %v\n", a.Ident, a.Type)
	// Return code
	return code
}

// List adds an identifier list to code: , .
func (code *Code) List() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add an identifier list to code
	code.c += ", "
	// Return code
	return code
}

// Listln adds an identifier list and a new line to code: ,\n.
func (code *Code) Listln() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add an identifier list and a new line to code
	code.c += ",\n"
	// Return code
	return code
}

type SelArgs struct {
	Val, Sel string
}

// SelField adds a field selector to code: val.sel. The value val and selector sel are
// provided by a. It returns nil if a is nil.
func (code *Code) SelField(a *SelArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add a field selector to code
	code.c += fmt.Sprintf("%v.%v", a.Val, a.Sel)
	// Return code
	return code
}

// SelMethod adds a method selector to code: val.sel(. The value val and selector sel are
// provided by a. It returns nil if a is nil.
func (code *Code) SelMethod(a *SelArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add a method selector to code
	code.c += fmt.Sprintf("%v.%v(", a.Val, a.Sel)
	// Return code
	return code
}

// IfArgs contains the left and right hand side expressions and the operator to generate
// an if statement with If.
type IfArgs struct {
	ExprLeft, ExprRight, Operator string
}

// If adds an if statement to code: if ExprLeft Operator ExprRight {\n.
// The left and right hand side expressions and the operator are provided by a.
// It returns nil if a is nil.
func (code *Code) If(a *IfArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add an if statement to code
	code.c += fmt.Sprintf("if %v %v %v {\n", a.ExprLeft, a.Operator, a.ExprRight)
	// Return code
	return code
}

type IfErrArgs struct {
	Method, Operator string
}

// IfErr adds an if statement for error handling to code: if err := Method; err Operator nil {\n.
// The method and operator are provided by a. It returns nil if a is nil.
func (code *Code) IfErr(a *IfErrArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add an if statement for error handling to code
	code.c += fmt.Sprintf("if err := %v; err %v nil {\n", a.Method, a.Operator)
	// Return code
	return code
}

// Return adds the return statement to code: return . It returns nil if code is nil.
func (code *Code) Return() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add the return statement to code
	code.c += "return "
	// Return code
	return code
}

// Addr adds an address operator to code: &. It returns nil if code is nil.
func (code *Code) Addr() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add an address operator to code
	code.c += "&"
	// Return code
	return code
}

// Ident adds an identifier to code: n. The identifier is provided by argument n.
func (code *Code) Ident(n string) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add identifier n to code
	code.c += n
	// Return code
	return code
}

// AssignmentArgs contains the left and right hand side expressions
// to generate an assignment with Assignment
type AssignmentArgs struct {
	ExprLeft, ExprRight string
}

// Assignment adds an assignment to code: ExprLeft = ExprRight.
// The left and right hand side expressions are provided by a.
// It returns nil if code is nil.
func (code *Code) Assignment(a *AssignmentArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add an assignment to code
	code.c += fmt.Sprintf("%v = %v", a.ExprLeft, a.ExprRight)
	// Return code
	return code
}

// CompositeLit adds the beginning of a composite literal to code: LiteralType{\n.
// The literal type is provided by argument LiteralType.
func (code *Code) CompositeLit(LiteralType string) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add the beginning of a composite literal to code
	code.c += fmt.Sprintf("%v{", LiteralType)
	// Return code
	return code
}

// ShortVarDeclArgs contains the identifier Ident and expression Expr to generate
// a short variable declaration with ShortVarDecl
type ShortVarDeclArgs struct {
	Ident, Expr string
}

// ShortVarDecl generates a short variable declaration: Ident := Expr\n. The identifier
// and expression is provided by a. It returns nil if a is nil.
func (code *Code) ShortVarDecl(a *ShortVarDeclArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add a short variable declaration to code
	code.c += fmt.Sprintf("%v := %v\n", a.Ident, a.Expr)
	// Return code
	return code
}

// KeyedElementArgs contains the key Key and element Elem to generate
// a keyed element of a composite literal with KeyedElement
type KeyedElementArgs struct {
	Key, Elem string
}

// KeyedElement generates a keyed element of a composite literal: Key: Element,\n. The key and element
// is provided by a. It returns nil if a is nil.
func (code *Code) KeyedElement(a *KeyedElementArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add keyed element of a composite literal to code
	code.c += fmt.Sprintf("%v: %v,\n", a.Key, a.Elem)
	// Return code
	return code
}

// Format formats the source code in code in canonical gfmt style.
// It uses Source from the go/format package. Format returns an error
// if code is nil or if go/format returns an error.
func (code *Code) Format() error {
	// Return an error in cae code is nil
	if code == nil {
		return tserr.NilPtr()
	}
	// Convert source code into a slice of bytes
	b := []byte(code.c)
	// Format the source code using Source from the go/format package
	o, e := format.Source(b)
	// Return an error in case Source fails
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "format source", Fn: "code", Err: e})
	}
	// Convert the formatted source code to string and store it in code
	code.c = string(o)
	// Return nil
	return nil
}
