package logger

import "testing"

func TestLoggerInit(t *testing.T) {
	Init()
	if Log == nil {
		t.Fatal("logger not initialized")
	}
	Log.Info("test")
}
