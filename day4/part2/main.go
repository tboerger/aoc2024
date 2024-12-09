package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/tboerger/aoc2024/pkg"
)

func main() {
	testing := false

	for _, arg := range os.Args[1:] {
		switch arg {
		case "--debug":
			slog.SetLogLoggerLevel(slog.LevelDebug)
		case "--testing":
			testing = true
		}
	}

	slog.Info("finished", "result", run(testing))
}

func run(testing bool) string {
	var (
		result = 0
		grid   []string
	)

	pkg.Parse(testing, func(rows []string) {
		grid = rows
	})

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if match(grid, r, c) {
				result += 1
			}
		}
	}

	return strconv.Itoa(result)
}

func match(grid []string, r, c int) bool {
	if r-1 < 0 || r+1 >= len(grid) || c-1 < 0 || c+1 >= len(grid[0]) {
		return false
	}

	tc := string(grid[r][c])

	tl := string(grid[r-1][c-1])
	tr := string(grid[r-1][c+1])
	bl := string(grid[r+1][c-1])
	br := string(grid[r+1][c+1])

	// tl to br && bl to tr
	// tr to bl && tl to br
	// br to tl && tr to bl
	// bl to tr && br to tl

	return (tl == "M" && tc == "A" && br == "S" && bl == "M" && tr == "S") ||
		(tr == "M" && tc == "A" && bl == "S" && tl == "M" && br == "S") ||
		(br == "M" && tc == "A" && tl == "S" && tr == "M" && bl == "S") ||
		(bl == "M" && tc == "A" && tr == "S" && br == "M" && tl == "S")
}
