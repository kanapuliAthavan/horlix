package logger

import (
	"testing"
)

func Test_InitAppLogger(t *testing.T) {
	paths := []string{
		"",
		"/tmp",
		"/var/log",
	}
	for _, path := range paths {
		err := InitAppLogger(path)
		if err != nil {
			t.Errorf("expected log horlix.log in the path %s, got %v", path, err)
		}
	}
	//Now InitLogger must have initiated aLogger
	if aLogger == nil {
		t.Error("expected to instatiate aLogger but got nil")
	}
}

func Test_InitTransLogger(t *testing.T) {
	paths := []string{
		"",
		"/tmp",
		"/var/log/",
	}
	for _, path := range paths {
		err := InitTransLogger(path)
		if err != nil {
			t.Errorf("expected Trans log transaction.log in the path %s, got %v", path, err)
		}
	}
}
