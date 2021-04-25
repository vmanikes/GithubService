// Package utils contains basic utility functions
package utils

import (
	"strings"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	funcName := GetFuncName()
	if !strings.Contains(funcName, "TestGetFuncName") {
		t.Errorf(funcName)
	}
}
