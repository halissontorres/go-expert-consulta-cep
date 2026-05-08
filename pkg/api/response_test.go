package api

import "testing"

func TestNewResposta(t *testing.T) {
	t.Parallel()

	s := NewResposta()
	if s == nil {
		t.Error("expected not nil")
	}
}
