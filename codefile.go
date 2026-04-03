// Package lpcode provides a simple and fluent API for programmatic Go code generation.
// It consists of two main components:
//
//   - Code: A builder for constructing Go source code strings using method chaining.
//     The Code type offers a wide range of methods for generating common Go syntax
//     including function declarations, struct types, control flow statements, variable
//     declarations, composite literals, and more. It supports formatting via go/format.
//
//   - Codefile: A utility for managing the lifecycle of code files. It handles creating
//     new files, appending header/footer templates, writing Code content, and formatting
//     the final output. It integrates with the tsfio package for file operations and
//     tserr for consistent error handling.
//
// The package is designed primarily to support the tserrgen code generator but can be
// used for other code generation tasks. It does not perform syntax validation; the
// generated code should be validated by go/format or other tools.
//
// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode

// Import standard library packages, tserr and tsfio.
import (
	"fmt"
	"go/format" // format

	"github.com/thorsphere/tserr" // tserr
	"github.com/thorsphere/tsfio" // tsfio
)

// Codefile represents a code file with a filename fn and filepath fp.
// It provides methods to start, write and finish the code file.
type Codefile struct {
	fn tsfio.Filename // filename
	fp tsfio.Filename // filepath which includes the filename and the directory path
}

// Constants for the header and footer file suffixes.
const (
	headerSuffix = ".header"
	footerSuffix = ".footer"
)

// NewCodefile creates a new Codefile with the given directory name dn and filename fn.
// It returns a pointer to the created Codefile and an error if any occurs during the creation
// process otherwise nil.
func NewCodefile(dn tsfio.Directory, fn tsfio.Filename) (*Codefile, error) {
	// Create the file path by joining the directory name and filename using tsfio.Path.
	f, err := tsfio.Path(dn, fn)
	// Return nil and an error if the file path cannot be created.
	if err != nil {
		return nil, tserr.Op(&tserr.OpArgs{Op: "Path", Fn: string(dn) + string(fn), Err: err})
	}
	// Create a new Codefile with the created file path and the given filename.
	cf := &Codefile{fp: f, fn: fn}
	// Return the created Codefile and nil.
	return cf, nil
}

// Filepath returns the file path of the Codefile. It returns an empty string
// if the Codefile is nil.
func (cf *Codefile) Filepath() tsfio.Filename {
	// Return an empty string if the Codefile is nil.
	if cf == nil {
		return ""
	}
	// Return the file path of the Codefile.
	return cf.fp
}

// StartFile starts the code file by resetting the file and appending the header file.
// It checks if the header file exists in the codefile path.
// If it does not exist in the codefile path, it checks if it exists in the current directory.
// If the header file does not exist in either location, it returns an error.
// Finally, it appends the header file to the code file using tsfio.AppendFile
// It returns an error if any occurs during the process otherwise nil.
func (cf *Codefile) StartFile() error {
	// Return an error if the Codefile is nil.
	if cf == nil {
		return tserr.NilPtr()
	}
	// Reset the file using tsfio.ResetFile. If an error occurs, return an error.
	if e := tsfio.ResetFile(cf.fp); e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "ResetFile", Fn: string(cf.fp), Err: e})
	}

	// Find the header file using the findFile method. If an error occurs, return an error.
	fh, err := cf.findFile(headerSuffix)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "findFile", Fn: string(cf.fn + headerSuffix), Err: err})
	}

	// Append the header file to the code file using tsfio.AppendFile.
	if e := tsfio.AppendFile(&tsfio.Append{FileA: cf.fp, FileI: fh}); e != nil {
		// If an error occurs, return an error.
		return tserr.Op(&tserr.OpArgs{Op: "AppendFile", Fn: string(cf.fp), Err: e})
	}
	// Return nil if the file is started successfully.
	return nil
}

// WriteCode writes the source code of c to the code file.
// It returns an error if any occurs during the process otherwise nil.
func (cf *Codefile) WriteCode(c *Code) error {
	// Return an error if the Codefile is nil.
	if cf == nil {
		return tserr.NilPtr()
	}
	// Return an error if the Code is nil.
	if c == nil {
		return tserr.NilPtr()
	}
	// Write the source code of c to the code file using tsfio.WriteStr.
	// If an error occurs, return the error of WriteStr.
	return tsfio.WriteStr(cf.fp, c.String())
}

// FinishFile finishes the code file by appending the footer file and formatting the code file.
// It returns an error if any occurs during the process otherwise nil.
func (cf *Codefile) FinishFile() error {
	// Return an error if the Codefile is nil.
	if cf == nil {
		return tserr.NilPtr()
	}

	// Find the footer file using the findFile method. If an error occurs, return an error.
	fe, err := cf.findFile(footerSuffix)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "findFile", Fn: string(cf.fn + footerSuffix), Err: err})
	}

	// Append the footer file to the code file using tsfio.AppendFile.
	if e := tsfio.AppendFile(&tsfio.Append{FileA: cf.fp, FileI: fe}); e != nil {
		// If an error occurs, return an error.
		return tserr.Op(&tserr.OpArgs{Op: "AppendFile", Fn: string(cf.fp), Err: e})
	}
	// Format the code file.
	if e := cf.Format(); e != nil {
		// If an error occurs, return an error.
		return tserr.Op(&tserr.OpArgs{Op: "format", Fn: string(cf.fp), Err: e})
	}
	// Return nil if the file is finished successfully.
	return nil
}

// findFile checks if a file with the given suffix exists in the codefile path or in the current directory.
// It returns the file path of the found file and an error if any occurs during the process otherwise nil.
func (cf *Codefile) findFile(suffix tsfio.Filename) (tsfio.Filename, error) {
	// Return an error if the Codefile is nil.
	if cf == nil {
		return "", tserr.NilPtr()
	}

	var (
		exists bool  = false
		e      error = nil
	)

	// Append the header file suffix to the codefile path.
	fn := cf.fp + suffix

	// Check if the header file exists in the codefile path using tsfio.ExistsFile.
	if exists, e = tsfio.ExistsFile(fn); e != nil {
		// Return an error if any error occurs during the check.
		return "", tserr.Op(&tserr.OpArgs{Op: "ExistsFile", Fn: string(fn), Err: e})
	}

	// If it exists in the codefile path, return the file path.
	if exists {
		return fn, nil
	}

	// If it does not exist, append the header file suffix to the code filename.
	fn = cf.fn + suffix

	// Check if the header file exists in the current directory using tsfio.ExistsFile.
	if exists, e = tsfio.ExistsFile(fn); e != nil {
		// Return an error if any error occurs during the check.
		return "", tserr.Op(&tserr.OpArgs{Op: "ExistsFile", Fn: string(fn), Err: e})
	}

	// If it does not exist in the current directory, return an error.
	if !exists {
		return "", tserr.NotExistent(fmt.Sprintf("%s or %s", string(fn), string(cf.fp+suffix)))
	}

	// If it exists in the current directory, return the file path.
	return fn, nil
}

// String returns the contents of the code file as a string. It returns an empty string
// if the Codefile is nil or if any error occurs during reading the file.
func (cf *Codefile) String() string {
	// Return an error if the Codefile is nil.
	if cf == nil {
		return ""
	}

	// Read the contents of the code file using tsfio.ReadFile.
	i, e := tsfio.ReadFile(cf.fp)

	// If an error occurs, return the error.
	if e != nil {
		return ""
	}

	// Return the contents of the code file as a string.
	return string(i)
}

// Format formats the code file using the go/format package.
// It returns an error if any occurs during the formatting process otherwise nil.
func (cf *Codefile) Format() error {
	// Return an error if the Codefile is nil.
	if cf == nil {
		return tserr.NilPtr()
	}
	// Read the contents of the code file using tsfio.ReadFile.
	i, e := tsfio.ReadFile(cf.fp)
	// If an error occurs, return the error.
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(cf.fp), Err: e})
	}
	// Format the contents of the code file using format.Source.
	o, e := format.Source(i)
	// If an error occurs, return the error.
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "Source", Fn: string(cf.fp), Err: e})
	}
	// Write the formatted contents back to the code file using tsfio.WriteSingleStr.
	if e := tsfio.WriteSingleStr(cf.fp, string(o)); e != nil {
		// If an error occurs, return the error.
		return tserr.Op(&tserr.OpArgs{Op: "WriteSingleStr", Fn: string(cf.fp), Err: e})
	}
	// Return nil if the code file is formatted successfully.
	return nil
}
