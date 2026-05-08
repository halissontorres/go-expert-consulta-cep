package log

import "testing"

func Test_NewLogger(t *testing.T) {
	t.Parallel()

	l := NewLogger(WithPrefix("TEST"))
	if l == nil {
		t.Fatal("expected non-nil logger")
	}
	l.Info("Hello World")
}

func Test_NewLogger_WithPrefix(t *testing.T) {
	t.Parallel()
	l := NewLogger(WithPrefix("TEST"))
	if l.prefix != "TEST" {
		t.Fatal("expected TEST prefix")
	}
}

func Test_NewLogger_WithoutPrefix(t *testing.T) {
	t.Parallel()
	l := NewLogger()
	if l.prefix != "" {
		t.Fatal("expected empty prefix")
	}
}

func Test_NewLogger_WithEmptyPrefix(t *testing.T) {
	t.Parallel()
	l := NewLogger()
	if l.prefix != "" {
		t.Fatal("expected empty prefix")
	}
}
