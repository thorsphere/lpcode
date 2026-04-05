// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode_test

// Import Go standard library package testing as well as lpcode, tserr and tsfio
import (
	"fmt"
	"testing" // testing

	"github.com/thorsphere/lpcode" // lpcode
	"github.com/thorsphere/tserr"  // tserr
	"github.com/thorsphere/tsfio"  // tsfio
)

var (
	testKey    string = "lothlorien" // test key
	testElem   string = "ithilien"   // test element
	testIdent  string = "fangorn"    // test identifier
	testExpr   string = "trollshaws" // test expression
	testCall   string = "brethil"    // test function call
	testStruct string = "mirkwood"   // test struct
	testType   string = "int"        // test type
	testInt    int    = 9            // test integer
	// test comment
	testComment string = "Lorem ipsum dolor sit amet, " +
		"consectetur adipisici elit, " +
		"sed eiusmod tempor incidunt ut labore " +
		"et dolore magna aliqua. Ut enim ad minim veniam, " +
		"quis nostrud exercitation ullamco laboris nisi ut " +
		"aliquid ex ea commodi consequat. Quis aute iure " +
		"reprehenderit in voluptate velit esse cillum dolore " +
		"eu fugiat nulla pariatur. Excepteur sint obcaecat " +
		"cupiditat non proident, sunt in culpa qui officia " +
		"deserunt mollit anim id est laborum."
)

// TestTestVariables tests generated test variables by Testvariables. The test fails if
// the generated test variables do not match the contents of the golden file.
func TestTestVariables(t *testing.T) {
	// Configure Testvars to generate test variables of types string, error, int and float
	vars := []lpcode.TestVar{
		{T: "string", N: "strFoo", V: "\"foobar\""},
		{T: "error", N: "errFoo", V: "fmt.Errorf(\"an error occurred\")"},
		{T: "int64", N: "intFoo", V: "1234"},
		{T: "float64", N: "floatFoo", V: "5678"},
	}
	// Retrieve test variables with Testvariables
	c := lpcode.NewCode().TestVarDecl(vars)
	// Evaluate the retrieved source code
	if e := evalCode(c, "testvariables"); e != nil {
		// The test fails if the retrieved code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestKeyedElement tests retrieved source code using the short variable declaration by ShortVarDecl,
// a keyed element by KeyedElement and function ending by FuncEnd. The test fails if the retrieved
// source code does not match the contents of the golden file.
func TestKeyedElement(t *testing.T) {
	// Retrieve the short variable declaration with ShortVarDecl
	c := lpcode.NewCode().ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: testIdent, Expr: testExpr + "{"})
	// Retrieve the keyed element with KeyedElement and the function ending with FuncEnd
	c.KeyedElement(&lpcode.KeyedElementArgs{Key: testKey, Elem: testElem}).FuncEnd()
	// Evaluate the retrieved source code
	if e := evalCode(c, "keyedelement"); e != nil {
		// The test fails if the generated source code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestTypeStruct tests retrieved source code using the type declaration for a struct type by TypeStruct,
// a variable specification by VarSpec and a block ending with BlockEnd. The test fails if the retrieved
// source code does not match the contents of the golden file.
func TestTypeStruct(t *testing.T) {
	// Retrieve the type declaration for a struct type with TypeStruct
	c := lpcode.NewCode().TypeStruct(testStruct)
	// Retrieve variable specifications with VarSpec and a block ending with BlockEnd
	c.VarSpec(&lpcode.VarSpecArgs{Ident: testIdent, Type: testType}).VarSpec(&lpcode.VarSpecArgs{Ident: testKey, Type: testType}).BlockEnd()
	// Evaluate the retrieved source code
	if e := evalCode(c, "typestruct"); e != nil {
		// The test fails if the generated source code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestSel tests retrieved source code using the field selector with SelField, the method selector with SelMethod and
// a parameters ending with ParamEnd. The test fails if the retrieved source code does not match the contents of the golden file.
func TestSel(t *testing.T) {
	// Retrieve the field selector with SelField, the method selector with SelMethod and a parameters ending with ParamEnd
	c := lpcode.NewCode().SelMethod(&lpcode.SelArgs{Val: testKey, Sel: testCall}).SelField(&lpcode.SelArgs{Val: testIdent, Sel: testType}).ParamEnd()
	// Evaluate the retrieved source code
	if e := evalCode(c, "testsel"); e != nil {
		// The test fails if the generated source code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestCall tests retrieved source code using a function call from Call, identifiers from Ident and an identifier list with List. The test fails if the
// retrieved source code does not match the contents of the golden file.
func TestCall(t *testing.T) {
	// Retrieve the function call from Call, identifiers from Ident and an identifier list with List
	c := lpcode.NewCode().Call(testCall).Ident(testIdent).List().Ident(testKey).ParamEndln()
	// Evaluate the retrieved source code
	if e := evalCode(c, "call"); e != nil {
		// The test fails if the generated source code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestAssignment tests retrieved source code using an assignment from Assignment,
// identifiers from Ident and a short variable declaration with ShortVarDecl.
// The test fails if the retrieved source code does not match the contents of the golden file.
func TestAssignment(t *testing.T) {
	// Create a new identifier by appending "2" to the test identifier
	testIdent2 := testIdent + "2"
	// Retrieve the function declaration with Func1.
	c := lpcode.NewCode().Func1(&lpcode.Func1Args{Name: testCall, Var: testIdent, Type: testType, Return: ""})
	// Retrieve the short variable declaration with ShortVarDecl
	c.ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: testIdent2, Expr: fmt.Sprintf("%v", testInt)})
	// Retrieve the assignment with Assignment
	c.Assignment(&lpcode.AssignmentArgs{ExprLeft: testIdent, ExprRight: testIdent2}).FuncEnd()
	// Format the retrieved source code.
	if e := c.Format(); e != nil {
		// The test fails if Format returns an error.
		t.Error(e)
	}
	// Evaluate the retrieved source code
	if e := evalCode(c, "assignment"); e != nil {
		// The test fails if the generated source code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestCall2 tests retrieved source code using a function call from Call, identifiers from Ident and an identifier list with a line ending with Listln.
// The test fails if the retrieved source code does not match the contents of the golden file.
func TestCall2(t *testing.T) {
	// Retrieve the function call from Call, identifiers from Ident and
	// an identifier list with a line ending with Listln
	c := lpcode.NewCode().Call(testCall).Ident(testIdent).Listln().Ident(testKey).Listln().ParamEndln()
	// Evaluate the retrieved source code
	if e := evalCode(c, "call2"); e != nil {
		// The test fails if the generated source code does not
		// match the contents of the golden file.
		t.Error(e)
	}
}

// TestFunc1 tests retrieved source code using a function declaration with Func1,
// a return statement with Return, an identifier with Ident and a function ending with FuncEnd.
// The test fails if the retrieved source code does not match the contents of the golden file.
func TestFunc1(t *testing.T) {
	// Retrieve the function declaration with Func1, a return statement with Return,
	// an identifier with Ident and a function ending with FuncEnd.
	c := lpcode.NewCode().Func1(&lpcode.Func1Args{
		Name:   testCall,
		Var:    testIdent,
		Type:   testType,
		Return: "*" + testType,
	}).Return().Addr().Ident(testIdent).FuncEnd()
	// Format the retrieved source code.
	if e := c.Format(); e != nil {
		// The test fails if Format returns an error.
		t.Error(e)
	}
	// Evaluate the retrieved source code
	if e := evalCode(c, "func1"); e != nil {
		// The test fails if the generated source code does not
		// match the contents of the golden file.
		t.Error(e)
	}
}

// TestIfErr tests retrieved source code using an type struct declaration with TypeStruct,
// a variable specification with VarSpec, a function declaration with Func1,
// a composite literal with CompositeLit, a keyed element with KeyedElement,
// a return statement with Return, an identifier with Ident,
// a block ending with BlockEnd, a function ending with FuncEnd,
// an if statement with If and an if statement for error handling with IfErr.
// The test fails if the retrieved source code does not match the contents of the golden file.
func TestIfErr(t *testing.T) {
	// Retrieve a new code.
	c := testIfErr()
	// Format the retrieved source code.
	if e := c.Format(); e != nil {
		// The test fails if Format returns an error.
		t.Error(e)
	}
	// Evaluate the retrieved source code
	if e := evalCode(c, "iferr"); e != nil {
		t.Error(e)
	}
}

// TestLineComment tests generated source code using a line comment by LineComment. The test fails
// if the generated source code does not match the contents of the golden file.
func TestLineComment(t *testing.T) {
	// Retrieve the line comment with LineComment
	c := lpcode.NewCode().LineComment(testComment)
	// Evaluate the retrieved source code
	if e := evalCode(c, "linecomment"); e != nil {
		// The test fails if the generated source code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestTestVariablesEmpty tests Testvariables to return an empty string in case
// no test variables are configured. The test fails if the returned string is
// not empty.
func TestTestVariablesEmpty(t *testing.T) {
	// // Configure Testvars to generate no test variables
	vars := []lpcode.TestVar{}
	// Retrieve test variables with Testvariables
	c := lpcode.NewCode().TestVarDecl(vars)
	// The test fails in case the returned string is not empty
	if c.String() != "" {
		t.Error(tserr.Empty("code"))
	}
}

// TestFormatEmpty tests Format to return nil for an empty code.
// The test fails if Format returns an error.
func TestFormatEmpty(t *testing.T) {
	// Retrieve the return value of Format for an empty code
	if e := lpcode.NewCode().Format(); e != nil {
		// The test fails if Format returns an error
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Format", Fn: "empty source code", Err: e}))
	}
}

// TestFormatNoCode tests Format to return an error in case code contains a
// placeholder text which is not valid Go source code. The test fails if Format returns nil.
func TestFormatNoCode(t *testing.T) {
	// Declare testcase loremipsum
	tc := "loremipsum"
	// Retrieve golden file path for the testcase
	fn, err := tsfio.GoldenFilePath(tc)
	// The test fails if GoldenFilePath returns an error
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "GoldenFilePath", Fn: tc, Err: err}))
	}
	// Retrieve the contents of the golden file for the test case
	li, err := tsfio.ReadFile(fn)
	// The test fails if ReadFile returns an error
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	// Retrieve a new Code instance with the contents of the golden file
	c := lpcode.NewCode().Ident(string(li))
	// The test fails if Format returns nil
	if e := c.Format(); e == nil {
		t.Error(tserr.NilFailed("format code"))
	}
}
