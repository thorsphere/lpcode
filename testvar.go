// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode

// Import Go standard library packages and tserr
import (
	"fmt" // fmt

	"github.com/thorsphere/tserr" // tserr
)

// TestVar is a struct that represents a test variable with its type T, name N and value V.
type TestVar struct {
	T string // type of testvariable
	N string // name of testvariable
	V string // value of testvariable
}

// TestvarDecl generates variable declarations for test variables.
// The test variables are generated based on the array tv of type Testvar.
// It returns nil if code is nil or if tv is nil.
// The generated variable declarations are added to code.
// The test variables are generated based on the type, name and value specified in tv. If no test variables are generated, an empty string is returned. The generated code should be validated by go/format or other tools. The test variables are generated in a var block. Each test variable is declared with its name, type and value, followed by a comment indicating the type of the test variable. The generated code can be used for unit tests to provide consistent test variables across different test cases.
func (code *Code) TestVarDecl(tv []TestVar) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case t is nil
	if tv == nil {
		return nil
	}
	// Initialize text
	text := ""

	// Iterate over testvars and add variable declarations to code
	for _, v := range tv {
		text += fmt.Sprintf("%v %v = %v // test variable of type %v\n", v.N, v.T, v.V, v.T)
	}

	// If test variables exist, add them to the variable declaration
	if text != "" {
		code.LineComment("test variables")
		code.c += "var (\n"
		code.c += text
		code.c += ")\n\n"
	}
	// Return generated code
	return code
}

// findTestVar is a helper function that takes t as a string representing a type, iterates
// through the array of TestVar structs tv, and checks if the type of any TestVar matches t.
// It returns a pointer to the corresponding TestVar struct for that type.
// It returns nil and an error if the type is not found in the predefined test cases
// or if there is an issue with the input.
func FindTestVar(t string, tv []TestVar) (*TestVar, error) {
	// Return nil and an error in case tv is nil
	if tv == nil {
		return nil, tserr.NilPtr()
	}
	// Iterate over testvars and return the one that matches the type t
	for _, v := range tv {
		// Check if the type of the current TestVar matches t
		if v.T == t {
			// If a match is found, return a pointer to the corresponding TestVar struct
			return &v, nil
		}
	}
	// Return nil and an error if the type is not found in the predefined test variables
	return nil, tserr.NotExistent(fmt.Sprintf("test variable of type %v", t))
}
