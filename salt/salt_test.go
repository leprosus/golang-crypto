package salt_test

import (
	"testing"

	"github.com/leprosus/golang-crypto/salt"
)

func TestGenerateSalt(t *testing.T) {
	t.Parallel()

	const length = 8

	s, err := salt.GenerateSalt(length)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(s) != length {
		t.Fatal("GenerateSalt returns a unexpected length salt")
	}
}
