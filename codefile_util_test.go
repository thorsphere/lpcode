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

// rmHF removes the header file and the footer file from the directory dn. In case of an error
// the execution stops.
func rmHF(t *testing.T, dn tsfio.Directory) {
	// Retrieve path to header file and footer file
	th, tf := pathHF(t, dn)

	// Remove the header file and the footer file. In case of an error execution stops.
	rm(t, th)
	rm(t, tf)
}

// pathHF retrieves the paths to the header file and the footer file in the directory dn.
func pathHF(t *testing.T, dn tsfio.Directory) (th, tf tsfio.Filename) {
	// Retrieve path to header file
	th = path(t, dn, testcfheader)
	// Retrieve path to footer file
	tf = path(t, dn, testcffooter)
	// Return the paths to the header file and the footer file
	return th, tf
}

// copyHF copies the header and footer files to the directory dn. In case of an error
// the execution stops.
func copyHF(t *testing.T, dn tsfio.Directory) {
	// Retrieve path to header file and footer file
	th, tf := pathHF(t, dn)

	// Copy the header file to the directory
	if err := tsfio.CopyFile(
		&tsfio.Copy{
			Src: testcfheader,
			Dst: th,
		}); err != nil {
		// The test fails if the header file cannot be copied to the directory
		t.Fatal(err)
	}

	// Copy the footer file to the directory
	if err := tsfio.CopyFile(
		&tsfio.Copy{
			Src: testcffooter,
			Dst: tf,
		}); err != nil {
		// The test fails if the header file cannot be copied to the directory
		t.Fatal(err)
	}
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
