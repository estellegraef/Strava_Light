package activity

import (
	"fmt"
	"testing"
)

func TestGetActivity(t *testing.T) {
	activity := GetActivity("user1", 1)
	fmt.Print(activity)
}
