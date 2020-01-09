package resources

import (
	"fmt"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	fmt.Printf(GetBasePath())
}

func TestGetTestGpxPath(t *testing.T) {
	fmt.Printf(GetTestGpxPath())
}

func TestGetTestZipPath(t *testing.T) {
	fmt.Printf(GetTestZipPath())
}

func TestGetTestInvalidPath(t *testing.T) {
	fmt.Printf(GetTestInvalidPath())
}
