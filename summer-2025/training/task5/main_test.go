package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	files, err := os.ReadDir("./tests")
	require.NoError(t, err)

	re := regexp.MustCompile(`^(\d+)(?:-(\d+))?`)

	sort.Slice(files, func(i, j int) bool {
		// Извлекаем числа из имен файлов для правильной числовой сортировки
		nameI := files[i].Name()
		nameJ := files[j].Name()

		// Функция для извлечения чисел из имени файла в формате "число-число.расширение"
		extractNumbers := func(filename string) (int, int, bool) {
			// Используем регулярное выражение для поиска паттерна "число-число"

			matches := re.FindStringSubmatch(filename)
			if len(matches) > 1 {
				// Первое число всегда должно быть
				firstNum, err1 := strconv.Atoi(matches[1])
				if err1 != nil {
					return 0, 0, false
				}

				// Второе число может отсутствовать
				secondNum := 0
				if len(matches) > 2 && matches[2] != "" {
					var err2 error
					secondNum, err2 = strconv.Atoi(matches[2])
					if err2 != nil {
						return firstNum, 0, true // Если второе число не распозналось, используем только первое
					}
				}

				return firstNum, secondNum, true
			}

			return 0, 0, false
		}

		// Извлекаем числа из имен файлов
		firstI, secondI, okI := extractNumbers(nameI)
		firstJ, secondJ, okJ := extractNumbers(nameJ)

		// Если оба имени содержат числа, сравниваем их
		if okI && okJ {
			// Сначала сравниваем первые числа
			if firstI != firstJ {
				return firstI < firstJ
			}
			// Если первые числа одинаковые, сравниваем вторые
			return secondI < secondJ
		}

		// Иначе сравниваем как строки
		return nameI < nameJ
	})

	for _, file := range files {
		if strings.Contains(file.Name(), ".") {
			continue
		}

		// if file.Name() != "8" {
		// 	continue
		// }

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
