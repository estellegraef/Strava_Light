package testStorage

import (
	"path"
	"runtime"
)

func GetBasePathTestStorage() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	//fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	return path.Dir(filename)
}
