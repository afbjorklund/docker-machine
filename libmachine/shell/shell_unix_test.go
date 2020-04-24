// +build !windows

package shell

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownShell(t *testing.T) {
	defer func(shell string) { os.Setenv("SHELL", shell) }(os.Getenv("SHELL"))
	os.Setenv("SHELL", "")

	shell, err := Detect()

	assert.Equal(t, err, ErrUnknownShell)
	assert.Empty(t, shell)
}

func TestDetectNoEnv(t *testing.T) {
	defer func(shell string) { os.Setenv("SHELL", shell) }(os.Getenv("SHELL"))
	os.Unsetenv("SHELL")

	shell, err := Detect()

	assert.Equal(t, filepath.Base(getCurrentUserShell()), shell)
	assert.NoError(t, err)
}
