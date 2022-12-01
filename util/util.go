package util

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// ReadFile is a wrapper over io/ioutil.ReadFile but also determines the dynamic
// absolute path to the file.
//
// Deprecated in favor of go:embed, refer to scripts/skeleton/tmpls
func ReadFile(pathFromCaller string) string {
	// Docs: https://golang.org/pkg/runtime/#Caller
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}

	// parse directory with pathFromCaller (which could be relative to Directory)
	absolutePath := path.Join(path.Dir(filename), pathFromCaller)

	// read the entire file & return the byte slice as a string
	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}
	// trim off new lines and tabs at end of input files
	strContent := string(content)
	return strings.TrimRight(strContent, "\n")
}

// Dirname is a port of __dirname in node
func Dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(filename)
}

// ToInt will case a given arg into an int type.
// Supported types are:
//    - string
func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}