package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		var l int
		fmt.Fscanln(in, &l)

		sb := strings.Builder{}

		for j := 0; j < l; j++ {
			s, _ := in.ReadString('\n')
			sb.WriteString(strings.TrimRight(s, "\n"))
		}

		folder := Folder{}
		err := json.Unmarshal([]byte(sb.String()), &folder)
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(out, countViruses(folder, false))
	}
}

func countViruses(folder Folder, hacked bool) int {
	var res int

	if !hacked {
		for _, file := range folder.Files {
			if filepath.Ext(file) == ".hack" {
				hacked = true
				break
			}
		}
	}

	if hacked {
		res = len(folder.Files)
	}

	for _, subFolder := range folder.Folders {
		res += countViruses(subFolder, hacked)
	}

	return res
}
