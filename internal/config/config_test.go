package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("invalid config file", func(t *testing.T) {
		_, err := LoadConfig("/tmp/foo.bar")
		require.Error(t, err)

		file, err := os.CreateTemp("", "log")
		if err != nil {
			t.FailNow()
			return
		}
		_, err = file.Write([]byte("invalid json"))
		if err != nil {
			t.FailNow()
			return
		}
		_, err = LoadConfig(file.Name())
		require.Error(t, err)
	})
}
