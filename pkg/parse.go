package pkg

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strings"
)

func Parse(testing bool, f func([]string)) {
	result := []string{}

	_, source, _, _ := runtime.Caller(1)
	filename := path.Join(path.Dir(source), "..", "input.txt")

	if testing {
		filename = path.Join(path.Dir(source), "..", "test.txt")
	}

	handle, err := os.Open(filename)
	Check(err)

	reader := bufio.NewReader(handle)

	for {
		row, prefix, err := reader.ReadLine()

		if err != nil {
			break
		}

		if prefix {
			continue
		}

		result = append(result, strings.TrimSpace(string(row)))
	}

	f(result)
}
