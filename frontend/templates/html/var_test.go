package html

import (
	"fmt"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	fmt.Printf(GetBasePath())
}

func TestGetLayoutPath(t *testing.T) {
	fmt.Printf(GetLayoutPath())
}

func TestGetDetailPath(t *testing.T) {
	fmt.Printf(GetDetailPath())
}

func TestGetEditPath(t *testing.T) {
	fmt.Printf(GetEditPath())
}

func TestGetIndexPath(t *testing.T) {
	fmt.Printf(GetIndexPath())
}

func TestGetItemsPath(t *testing.T) {
	fmt.Printf(GetItemsPath())
}

func TestGetSearchPath(t *testing.T) {
	fmt.Printf(GetSearchPath())
}

func TestGetUploadPath(t *testing.T) {
	fmt.Printf(GetUploadPath())
}
