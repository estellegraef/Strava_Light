/*
 * 2848869
 * 8089098
 * 3861852
 */
package activity

import (
	"fmt"
	"testing"
)

func TestGetActivity(t *testing.T) {
	activity := GetActivity("user1", 1)
	fmt.Print(activity)
}
