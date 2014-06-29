package cache

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "a-test.")
	if err != nil {
		t.Fatalf("Unable to create temp file: %s", err)
	}
	defer os.Remove(file.Name())
	_, err = Open(file.Name())
	assert.NoError(t, err, "Unable to open cache")
}
