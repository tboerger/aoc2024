package pkg

import (
	"bufio"
	"log/slog"
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

	if err != nil {
		slog.Error("failed to read file", "err", err)
		os.Exit(1)
	}

	reader := bufio.NewReaderSize(handle, 10*4096)

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
