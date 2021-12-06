package hello_test

import (
	"testing"

	"github.com/lucasmmo/setup-go/pkg/hello"
)

func TestSayHello(t *testing.T) {
	t.Run("Should say hello world", func(t *testing.T) {
		s := hello.Say("")
		if s != "Hello, World!" {
			t.Errorf("expected: Hello, World!, received: %s", s)
		}
	})
}
