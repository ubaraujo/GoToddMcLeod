package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("sum was incorrect, got: %d, want: %d", total, 10)
	}
}

func TestSub(t *testing.T) {
	total := sub(20, 10)
	if total != 10 {
		t.Errorf("sub was incorrect, got: %d, want: %d", total, 10)
	}
}

func TestSubNeg(t *testing.T) {
	got := sub(10, 20)
	want := -10

	if got != want {
		t.Errorf("sub was incorrect, got: %v, want: %v: ", got, want)
	}
}

func TestParadise(t *testing.T) {
	got := Paradise("Cabo verde")
	want := "My idea of padradise is Cabo verde"

	if got != want {
		log.Fatalf("error - want %s and got %s", got, want)
	}
}
