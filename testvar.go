// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode

import "fmt"

type Testvar struct {
	T string // type of testvariable
	N string // name of testvariable
	V string // value of testvariable
}

// TestvarDecl generates variable declarations for test variables.
// The test variables are generated based on the array tv of type Testvar.
// It returns nil if code is nil or if tv is nil.
// The generated variable declarations are added to code.
// The test variables are generated based on the type, name and value specified in tv. If no test variables are generated, an empty string is returned. The generated code should be validated by go/format or other tools. The test variables are generated in a var block. Each test variable is declared with its name, type and value, followed by a comment indicating the type of the test variable. The generated code can be used for unit tests to provide consistent test variables across different test cases.
func (code *Code) TestVarDecl(tv []Testvar) *Code {
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
		text += fmt.Sprintf("%v %v = %v //test variable type %v\n", v.N, v.T, v.V, v.T)
	}

	// If test variables exist, add them to the variable declaration
	if text != "" {
		code.c += "var (\n"
		code.c += text
		code.c += ")\n\n"
	}
	// Return generated code
	return code
}
