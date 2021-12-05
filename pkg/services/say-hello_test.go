package services_test

import (
	"testing"

	"github.com/lucasmmo/setup-go/pkg/services"
)

func TestSayHello(t *testing.T) {
	t.Run("Should say hello world", func(t *testing.T) {
		s := services.SayHello("")
		if s != "Hello, World!" {
			t.Errorf("expected: Hello, World!, received: %s", s)
		}
	})
}
