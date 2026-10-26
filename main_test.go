package main

import "testing"

func TestTitle(t *testing.T) {
	if got := title("sandbox app"); got != "Sandbox App" {
		t.Errorf("title = %q", got)
	}
}
