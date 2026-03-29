// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode_test

import (
	"os"      // os
	"testing" // testing

	"github.com/thorsphere/tserr" // tserr
	"github.com/thorsphere/tsfio" // tsfio
)

// The testcases use these tokens
const (
	testprefix   string          = "lpcode_*"      // used as prefix for temporary files or directories
	testcfheader tsfio.Filename  = "testcf.header" // the header file for codefile tests
	testcffooter tsfio.Filename  = "testcf.footer" // the footer file for codefile tests
	testcfname   string          = "testcf"        // the filename for the golden file and the code file
	testdatadir  tsfio.Directory = "testdata/"     // the directory for test data
)

// tmpDir creates a new temporary directory in the default directory for temporary files
// with the prefix testprefix and a random string to the end. In case of an error
// the execution stops.
func tmpDir(t *testing.T) tsfio.Directory {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// Create the temporary directory
	d, err := os.MkdirTemp("", testprefix)
	// Stop execution in case of an error
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "create temp dir", Fn: d, Err: err}))
	}
	// Return the temporary Directory
	return tsfio.Directory(d)
}

// tmpFile creates a new tmporary file in the default directory for temporary files
// with the prefix testprefix and a random string to the end. In case of an error
// the execution stops.
func tmpFile(t *testing.T, dn tsfio.Directory) tsfio.Filename {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// Create the temporary file
	f, err := os.CreateTemp(string(dn), testprefix)
	// Stop execution in case of an error
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "create temp file", Fn: f.Name(), Err: err}))
	}
	// Retrieve the name of the temporary file
	fn := tsfio.Filename(f.Name())
	// Close the temporary file
	if e := f.Close(); e != nil {
		// The test fails if Close returns an error
		t.Error(tserr.Op(&tserr.OpArgs{Op: "close temp file", Fn: string(fn), Err: e}))
	}
	// Return the temporary Filename
	return fn
}

// rm removes file named Filename a or empty directory Directory a. In case of an error
// execution stops.
func rm[T tsfio.Fio](t *testing.T, a T) {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// Remove file or empty directory
	if err := os.Remove(string(a)); err != nil {
		// Stop execution in case of an error
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "Remove", Fn: string(a), Err: err}))
	}
}

// path creates a file path by joining the directory name dn and filename fn using tsfio.Path.
// In case of an error execution stops.
func path(t *testing.T, dn tsfio.Directory, fn tsfio.Filename) tsfio.Filename {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// Create the file path by joining the directory name and filename using tsfio.Path.
	pn, err := tsfio.Path(dn, fn)
	// Stop execution in case of an error
	if err != nil {
		t.Fatal(err)
	}
	// Return the file path
	return pn
}
