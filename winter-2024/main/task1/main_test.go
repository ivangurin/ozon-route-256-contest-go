package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {

	for i := 1; ; i++ {

		file, err := os.Open(fmt.Sprintf("tests/%d", i))
		if err != nil {
			break
		}
		defer file.Close()

		t.Run(fmt.Sprintf("Test:%d", i), func(t *testing.T) {

			in := bufio.NewReader(file)

			expected, err := os.ReadFile(fmt.Sprintf("tests/%d.a", i))
			require.Nil(t, err)

			var buffer bytes.Buffer
			out := bufio.NewWriter(&buffer)

			Run(in, out)

			out.Flush()

			result, err := io.ReadAll(bufio.NewReader(&buffer))
			require.Nil(t, err)

			require.Equal(t, string(expected), string(result))

		})

	}

}
