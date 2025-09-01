package language_wizard

import (
	"sync/atomic"
	"testing"
)

// // // // // // // // // // // //

func TestCurrentLanguage(t *testing.T) {
	obj := mustNew(t)
	if obj.CurrentLanguage() != "en" {
		t.Fatalf("CurrentLanguage = %q, want %q", obj.CurrentLanguage(), "en")
	}
}

func TestWords_ReturnsCopy(t *testing.T) {
	obj := mustNew(t)
	w := obj.Words()
	w["hi"] = "HACKED" // змінюємо копію
	if obj.Get("hi", "") != "Hello" {
		t.Fatalf("internal words must not change when external map is modified")
	}
}

func TestGet_EmptyID_ReturnsDefault(t *testing.T) {
	obj := mustNew(t)
	if got := obj.Get("", "DEF"); got != "DEF" {
		t.Fatalf("Get(empty) = %q, want %q", got, "DEF")
	}
}

func TestGet_FoundAndMissingWithLog(t *testing.T) {
	obj := mustNew(t)

	var logCount int32
	obj.SetLog(func(s string) {
		atomic.AddInt32(&logCount, 1)
	})

	if got := obj.Get("hi", "DEF"); got != "Hello" {
		t.Fatalf("Get(hi) = %q, want %q", got, "Hello")
	}
	if got := obj.Get("unknown", "DEF"); got != "DEF" {
		t.Fatalf("Get(unknown) = %q, want %q", got, "DEF")
	}
	if c := atomic.LoadInt32(&logCount); c != 1 {
		t.Fatalf("log called %d times, want 1", c)
	}
}
