package language_wizard

import (
	"errors"
	"testing"
	"time"
)

// // // // // // // // // // // //

func mustNew(t *testing.T) *LanguageWizardObj {
	t.Helper()
	obj, err := New("en", map[string]string{"hi": "Hello"})
	if err != nil {
		t.Fatalf("New failed: %v", err)
	}
	return obj
}

// //

func TestNew_Success(t *testing.T) {
	obj := mustNew(t)
	if got := obj.CurrentLanguage(); got != "en" {
		t.Fatalf("CurrentLanguage = %q, want %q", got, "en")
	}
	if got := obj.Get("hi", ""); got != "Hello" {
		t.Fatalf("Get(hi) = %q, want %q", got, "Hello")
	}
}

func TestNew_Errors(t *testing.T) {
	if _, err := New("", map[string]string{"k": "v"}); !errors.Is(err, ErrNilIsoLang) {
		t.Fatalf("want ErrNilIsoLang, got %v", err)
	}
	if _, err := New("en", nil); !errors.Is(err, ErrNilWords) {
		t.Fatalf("want ErrNilWords (nil), got %v", err)
	}
	if _, err := New("en", map[string]string{}); !errors.Is(err, ErrNilWords) {
		t.Fatalf("want ErrNilWords (empty), got %v", err)
	}
}

func TestClose_Effects(t *testing.T) {
	obj := mustNew(t)

	done := make(chan EventType, 1)
	go func() { done <- obj.Wait() }()

	obj.Close()

	select {
	case ev := <-done:
		if ev != EventClose {
			t.Fatalf("Wait after Close = %v, want EventClose", ev)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Wait did not unblock on Close")
	}

	if len(obj.Words()) != 0 {
		t.Fatalf("Words should be cleared after Close")
	}

	if err := obj.SetLanguage("de", map[string]string{"hi": "Hallo"}); !errors.Is(err, ErrClosed) {
		t.Fatalf("SetLanguage after Close = %v, want ErrClosed", err)
	}
}
