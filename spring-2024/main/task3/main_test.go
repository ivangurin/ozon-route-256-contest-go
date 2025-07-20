package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	files, err := os.ReadDir("./tests")
	require.NoError(t, err)

	for _, file := range files {
		if strings.Contains(file.Name(), ".") {
			continue
		}

		fileTask, err := os.Open(fmt.Sprintf("tests/%s", file.Name()))
		require.NoError(t, err)
		defer fileTask.Close()

		t.Run(fmt.Sprintf("Test:%s", file.Name()), func(t *testing.T) {
			in := bufio.NewReader(fileTask)

			expected, err := os.ReadFile(fmt.Sprintf("tests/%s.a", file.Name()))
			require.NoError(t, err)

			var buffer bytes.Buffer
			out := bufio.NewWriter(&buffer)

			Run(in, out)

			out.Flush()

			result, err := io.ReadAll(bufio.NewReader(&buffer))
			require.NoError(t, err)

			require.Equal(t, string(expected), string(result))
		})
	}
}
