package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {

	for i := 1; i < 100; i++ {

		s := strconv.Itoa(i)

		file, err := os.Open("tests/" + s)
		if err != nil {
			break
		}
		defer file.Close()

		t.Run("Test:"+s, func(t *testing.T) {

			in := bufio.NewReader(file)

			expected, err := os.ReadFile("tests/" + s + ".a")
			require.Nil(t, err)

			var buffer bytes.Buffer
			out := bufio.NewWriter(&buffer)

			Run(in, out)

			out.Flush()

			reader := bufio.NewReader(&buffer)
			result, err := readStringAll(reader)
			require.Nil(t, err)

			require.Equal(t, string(expected), result)

		})

	}

}

func readStringAll(reader *bufio.Reader) (string, error) {

	res := ""

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				res = res + line
				break
			} else {
				return "", err
			}
		}
		res = res + line
	}

	return res, nil

}
