package main

import "testing"

func TestState_noGodMode(t *testing.T) {
	if enableGodMode {
		t.Fatal("Don't leave god mode enabled :P")
	}
}
