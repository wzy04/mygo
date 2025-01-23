package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	result := Hello()
	if result != expected {
		t.Errorf("期望输出 %q，实际得到 %q", expected, result)
	}
}
