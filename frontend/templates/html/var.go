package html

import (
	"path"
	"path/filepath"
	"runtime"
)

func GetBasePath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	//fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	return path.Dir(filename)
}

func GetLayoutPath() string {
	return filepath.Join(GetBasePath(), "layout.html")
}

func GetDetailPath() string {
	return filepath.Join(GetBasePath(), "detail.html")
}

func GetEditPath() string {
	return filepath.Join(GetBasePath(), "edit.html")
}

func GetIndexPath() string {
	return filepath.Join(GetBasePath(), "index.html")
}

func GetItemsPath() string {
	return filepath.Join(GetBasePath(), "items.html")
}

func GetSearchPath() string {
	return filepath.Join(GetBasePath(), "search.html")
}

func GetUploadPath() string {
	return filepath.Join(GetBasePath(), "upload.html")
}
