// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode_test

// Import Go packages lpcode, tserr and tsfio
import (
	// testing

	"github.com/thorsphere/lpcode" // lpcode
	"github.com/thorsphere/tserr"  // tserr
	"github.com/thorsphere/tsfio"  // tsfio
)

// evalCode evaluates the source code of c by comparing the source code
// with the associated golden file for test case tc. The function returns
// an error if the source code does not match the contents of the golden file.
func evalCode(
	c *lpcode.Code,
	tc string,
) error {
	// Return an error if c is nil
	if c == nil {
		return tserr.NilPtr()
	}
	// Format the retrieved code
	if e := c.Format(); e != nil {
		// The test fails if Format returns an error
		return e
	}
	// Evaluate the retrieved code with the golden file
	if e := tsfio.EvalGoldenFile(&tsfio.Testcase{Name: tc, Data: c.String()}); e != nil {
		// The test fails if the retrieved code does not match the contents of the golden file
		return e
	}
	// Return nil
	return nil
}

// testIfErr generates a code snippet that demonstrates the usage of the IfErr method in lpcode.
// It returns the generated code snippet as a pointer to lpcode.Code.
// The generated code snippet includes a struct declaration, variable specifications,
// a function declaration, an if statement for error handling, another if statement,
// a composite literal with keyed elements, and return statements.
func testIfErr() *lpcode.Code {
	// Retrieve a new code.
	return lpcode.NewCode().
		// type struct declaration with TypeStruct
		TypeStruct("foo").
		// variable specification with VarSpec
		VarSpec(&lpcode.VarSpecArgs{Ident: "A", Type: "int"}).
		VarSpec(&lpcode.VarSpecArgs{Ident: "B", Type: "string"}).BlockEnd().
		// function declaration with Func1
		Func1(&lpcode.Func1Args{Name: "Example", Var: "x", Type: "int", Return: "error"}).
		// if statement for error handling with IfErr
		IfErr(&lpcode.IfErrArgs{Method: "doSomething()", Operator: "!="}).Return().Ident("err").BlockEnd().
		// if statement with If
		If(&lpcode.IfArgs{ExprLeft: "x", Operator: ">", ExprRight: "10"}).
		// composite literal with CompositeLit
		CompositeLit("foo").
		// keyed element with KeyedElement
		KeyedElement(&lpcode.KeyedElementArgs{Key: "A", Elem: "1"}).
		KeyedElement(&lpcode.KeyedElementArgs{Key: `B`, Elem: `"hello"`}).
		// return statement with Return, an identifier with Ident, a block ending with BlockEnd and a function ending with FuncEnd
		BlockEnd().BlockEnd().Return().Ident("nil").FuncEnd()
}
