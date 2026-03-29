// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode_test

// Import Go standard library package testing as well as lpcode and tserr
import (
	"testing" // testing

	"github.com/thorsphere/lpcode" // lpcode
	"github.com/thorsphere/tserr"  // tserr
)

// TestTestVariablesNil tests Testvariables to return nil in case
// *Code is nil. The test fails if Testvariables does not return nil.
func TestTestVariablesNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Testvariables does not return nil.
	if n := c.Testvariables(&lpcode.Testvars{}); n != nil {
		t.Error(tserr.NilExpected("Testvariables"))
	}
}

// TestShortVarDeclNil tests ShortVarDecl to return nil in case
// *Code is nil. The test fails if ShortVarDecl does not return nil.
func TestShortVarDeclNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if ShortVarDecl does not return nil.
	if n := c.ShortVarDecl(&lpcode.ShortVarDeclArgs{}); n != nil {
		t.Error(tserr.NilExpected("ShortVarDecl"))
	}
}

// TestIdentNil tests Ident to return nil in case
// *Code is nil. The test fails if Ident does not return nil.
func TestIdentNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Ident does not return nil.
	if n := c.Ident(""); n != nil {
		t.Error(tserr.NilExpected("Ident"))
	}
}

// TestListNil tests List to return nil in case
// *Code is nil. The test fails if List does not return nil.
func TestListNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if List does not return nil.
	if n := c.List(); n != nil {
		t.Error(tserr.NilExpected("List"))
	}
}

// TestListlnNil tests Listln to return nil in case
// *Code is nil. The test fails if Listln does not return nil.
func TestListlnNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Listln does not return nil.
	if n := c.Listln(); n != nil {
		t.Error(tserr.NilExpected("Listln"))
	}
}

// TestParamEndlnNil tests ParamEndln to return nil in case
// *Code is nil. The test fails if ParamEndln does not return nil.
func TestParamEndlnNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if ParamEndln does not return nil.
	if n := c.ParamEndln(); n != nil {
		t.Error(tserr.NilExpected("ParamEndln"))
	}
}

// TestCallNil tests Call to return nil in case
// *Code is nil. The test fails if Call does not return nil.
func TestCallNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Call does not return nil.
	if n := c.Call(""); n != nil {
		t.Error(tserr.NilExpected("Call"))
	}
}

// TestLineCommentNil tests LineComment to return nil in case
// *Code is nil. The test fails if LineComment does not return nil.
func TestLineCommentNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if LineComment does not return nil.
	if n := c.LineComment(""); n != nil {
		t.Error(tserr.NilExpected("LineComment"))
	}
}

// TestTypeStructNil tests TypeStruct to return nil in case
// *Code is nil. The test fails if TypeStruct does not return nil.
func TestTypeStructNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if TypeStruct does not return nil.
	if n := c.TypeStruct(""); n != nil {
		t.Error(tserr.NilExpected("TypeStruct"))
	}
}

// TestStringNil tests String to return an empty string in case
// *Code is nil. The test fails if String does not return an empty string.
func TestStringNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Ident does not return an empty string.
	if n := c.String(); n != "" {
		t.Error(tserr.NotEqual(&tserr.NotEqualArgs{X: n, Y: ""}))
	}
}

// TestParamEndNil tests ParamEnd to return nil in case
// *Code is nil. The test fails if ParamEnd does not return nil.
func TestParamEndNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if ParamEnd does not return nil.
	if n := c.ParamEnd(); n != nil {
		t.Error(tserr.NilExpected("ParamEnd"))
	}
}

// TestFuncEndNil tests FuncEnd to return nil in case
// *Code is nil. The test fails if FuncEnd does not return nil.
func TestFuncEndNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if FuncEnd does not return nil.
	if n := c.FuncEnd(); n != nil {
		t.Error(tserr.NilExpected("FuncEnd"))
	}
}

// TestBlockEndNil tests BlockEnd to return nil in case
// *Code is nil. The test fails if BlockEnd does not return nil.
func TestBlockEndNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if BlockEnd does not return nil.
	if n := c.BlockEnd(); n != nil {
		t.Error(tserr.NilExpected("BlockEnd"))
	}
}

// TestKeyedElementNil tests KeyedElement to return nil in case
// *Code is nil. The test fails if KeyedElement does not return nil.
func TestKeyedElementNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if KeyedElement does not return nil.
	if n := c.KeyedElement(&lpcode.KeyedElementArgs{}); n != nil {
		t.Error(tserr.NilExpected("KeyedElement"))
	}
}

// TestVarSpecNil tests VarSpec to return nil in case
// *Code is nil. The test fails if VarSpec does not return nil.
func TestVarSpecNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if VarSpec does not return nil.
	if n := c.VarSpec(&lpcode.VarSpecArgs{}); n != nil {
		t.Error(tserr.NilExpected("VarSpec"))
	}
}

// TestSelFieldNil tests SelField to return nil in case
// *Code is nil. The test fails if SelField does not return nil.
func TestSelFieldNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if SelField does not return nil.
	if n := c.SelField(&lpcode.SelArgs{}); n != nil {
		t.Error(tserr.NilExpected("SelField"))
	}
}

// TestSelMethodNil tests SelMethod to return nil in case
// *Code is nil. The test fails if SelMethod does not return nil.
func TestSelMethodNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if SelField does not return nil.
	if n := c.SelMethod(&lpcode.SelArgs{}); n != nil {
		t.Error(tserr.NilExpected("SelMethod"))
	}
}

// TestTestVariablesNil2 tests Testvariables to return nil in case
// t is nil. The test fails if Testvariables does not return nil.
func TestTestVariablesNil2(t *testing.T) {
	// The test fails if Testvariables does not return nil.
	if n := lpcode.NewCode().Testvariables(nil); n != nil {
		t.Error(tserr.NilExpected("Testvariables"))
	}
}

// TestShortVarDeclNil2 tests ShortVarDecl to return nil in case
// a is nil. The test fails if ShortVarDecl dows not return nil.
func TestShortVarDeclNil2(t *testing.T) {
	// The test fails if ShortVarDecl does not return nil.
	if n := lpcode.NewCode().ShortVarDecl(nil); n != nil {
		t.Error(tserr.NilExpected("ShortVarDecl"))
	}
}

// TestSelFieldNil2 tests SelField to return nil in case
// a is nil. The test fails if SelField dows not return nil.
func TestSelFieldNil2(t *testing.T) {
	// The test fails if SelField does not return nil.
	if n := lpcode.NewCode().SelField(nil); n != nil {
		t.Error(tserr.NilExpected("SelField"))
	}
}

// TestSelMethodNil2 tests SelMethod to return nil in case
// a is nil. The test fails if SelMethod dows not return nil.
func TestSelMethodNil2(t *testing.T) {
	// The test fails if SelMethod does not return nil.
	if n := lpcode.NewCode().SelMethod(nil); n != nil {
		t.Error(tserr.NilExpected("SelMethod"))
	}
}

// TestKeyedElementNil2 tests KeyedElement to return nil in case
// a is nil. The test fails if KeyedElement does not return nil.
func TestKeyedElementNil2(t *testing.T) {
	// The test fails if KeyedElement does not return nil.
	if n := lpcode.NewCode().KeyedElement(nil); n != nil {
		t.Error(tserr.NilExpected("KeyedElement"))
	}
}

// TestFunc1 tests Func1 to return nil in case
// a is nil. The test fails if Func1 does not return nil.
func TestFunc1Nil(t *testing.T) {
	// The test fails if Func1 does not return nil.
	if n := lpcode.NewCode().Func1(nil); n != nil {
		t.Error(tserr.NilExpected("Func1"))
	}
}

// TestFunc1Nil2 tests Func1 to return nil in case
// *Code is nil. The test fails if Func1 does not return nil.
func TestFunc1Nil2(t *testing.T) {
	// Declare c as type *Code and assign nil.
	var c *lpcode.Code = nil
	// The test fails if Func1 does not return nil.
	if n := c.Func1(&lpcode.Func1Args{}); n != nil {
		t.Error(tserr.NilExpected("Func1"))
	}
}

// TestAssignmentNil tests Assignment to return nil in case
// *Code is nil. The test fails if Assignment does not return nil.
func TestAssignmentNil(t *testing.T) {
	// Declare c as type *Code and assign nil.
	var c *lpcode.Code = nil
	// The test fails if Assignment does not return nil in case code is nil.
	if n := c.Assignment(&lpcode.AssignmentArgs{}); n != nil {
		t.Error(tserr.NilExpected("Assignment"))
	}
}

// TestAssignmentNil2 tests Assignment to return nil in case
// a is nil. The test fails if Assignment does not return nil in case a is nil.
func TestAssignmentNil2(t *testing.T) {
	// The test fails if Assignment does not return nil in case a is nil.
	if n := lpcode.NewCode().Assignment(nil); n != nil {
		t.Error(tserr.NilExpected("Assignment"))
	}
}

// TestReturnNil tests Return to return nil in case
// *Code is nil. The test fails if Return does not return nil.
func TestReturnNil(t *testing.T) {
	// Declare c as type *Code and assign nil.
	var c *lpcode.Code = nil
	// The test fails if Return does not return nil.
	if n := c.Return(); n != nil {
		t.Error(tserr.NilExpected("Return"))
	}
}

// TestIfNil tests If to return an empty string in case
// *Code is nil. The test fails if Code does not return nil.
func TestIfNil(t *testing.T) {
	// Declare c as type *Code and assign nil.
	var c *lpcode.Code = nil
	// The test fails if If does not return nil.
	if n := c.If(&lpcode.IfArgs{}); n != nil {
		t.Error(tserr.NilExpected("If"))
	}
}

// TestIfNil2 tests If to return nil in case
// a is nil. The test fails if If does not return nil.
func TestIfNil2(t *testing.T) {
	// The test fails if If does not return nil.
	if n := lpcode.NewCode().If(nil); n != nil {
		t.Error(tserr.NilExpected("If"))
	}
}

// TestIfErrNil2 tests If to return nil in case
// a is nil. The test fails if If does not return nil.
func TestIfErrNil2(t *testing.T) {
	// The test fails if If does not return nil.
	if n := lpcode.NewCode().IfErr(nil); n != nil {
		t.Error(tserr.NilExpected("IfErr"))
	}
}

// TestIfErrNil tests IfErr to return nil in case
// *Code is nil. The test fails if IfErr does not return nil.
func TestIfErrNil(t *testing.T) {
	// Declare c as type *Code and assign nil.
	var c *lpcode.Code = nil
	// The test fails if IfErr does not return nil.
	if n := c.IfErr(&lpcode.IfErrArgs{}); n != nil {
		t.Error(tserr.NilExpected("IfErr"))
	}
}

// TestCompositeLitNil tests CompositeLit to return nil in case
// *Code is nil. The test fails if CompositeLit does not return nil.
func TestCompositeLitNil(t *testing.T) {
	// Declare c as type *Code and assign nil.
	var c *lpcode.Code = nil
	// The test fails if CompositeLit does not return nil.
	if n := c.CompositeLit(""); n != nil {
		t.Error(tserr.NilExpected("CompositeLit"))
	}
}

// TestAddrNil tests Addr to return nil in case
// *Code is nil. The test fails if Addr does not return nil.
func TestAddrNil(t *testing.T) {
	// Declare c as type *Code and assign nil.
	var c *lpcode.Code = nil
	// The test fails if Addr does not return nil.
	if n := c.Addr(); n != nil {
		t.Error(tserr.NilExpected("Addr"))
	}
}

// TestFormatNil tests Format to return nil in case
// *Code is nil. The test fails if Format does not return nil.
func TestFormatNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Format does not return nil.
	if n := c.Format(); n == nil {
		t.Error(tserr.NilFailed("Format"))
	}
}
