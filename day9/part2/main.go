package main

import (
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/tboerger/aoc2024/pkg"
)

type Space struct {
	Pos, Length int
}
type File struct {
	Val, Pos, Length int
	Moved            bool
}

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
		diskmap []int
		blocks  []int
		files   = make(map[int]*File)
		spaces  = make([]*Space, 0)
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			for _, char := range strings.Split(row, "") {
				num, _ := strconv.Atoi(char)
				diskmap = append(diskmap, num)
			}
		}
	})

	blocks, files, spaces = parseDiskmap(diskmap, files, spaces)
	slog.Debug("parsed", "count", len(blocks), "val", blocks)

	blocks = compactDisk(blocks, files, spaces)
	slog.Debug("compact", "count", len(blocks), "val", blocks)

	return strconv.Itoa(calculateChecksum(blocks))
}

func parseDiskmap(diskmap []int, files map[int]*File, spaces []*Space) ([]int, map[int]*File, []*Space) {
	result := []int{}

	for i, val := range diskmap {
		if i%2 == 0 {
			result = append(result, slices.Repeat([]int{i / 2}, val)...)
			files[i/2] = &File{i / 2, len(result) - val, val, false}
		} else {
			result = append(result, slices.Repeat([]int{-1}, val)...)
			spaces = append(spaces, &Space{len(result) - val, val})
		}
	}

	return result, files, spaces
}

func compactDisk(blocks []int, files map[int]*File, spaces []*Space) []int {
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i] <= 0 {
			continue
		}

		if f, ok := files[blocks[i]]; ok {
			if f.Moved {
				continue
			}

			for sid, space := range spaces {
				slog.Debug("check", "blocks", blocks, "pos", i, "file", f, "space", space)

				if space.Length >= f.Length {
					for x := f.Pos; x < f.Pos+f.Length; x++ {
						blocks[x] = -1
					}

					for x := space.Pos; x < space.Pos+f.Length; x++ {
						blocks[x] = f.Val
					}

					if space.Length-f.Length < 1 {
						spaces = removeSpace(spaces, sid)
					} else {
						space.Pos = space.Pos + f.Length
						space.Length = space.Length - f.Length
					}

					slog.Debug("moved", "blocks", blocks, "pos", i, "file", f, "space", space)
					break
				}
			}

			f.Moved = true
		}
	}

	return blocks
}

func calculateChecksum(blocks []int) int {
	result := 0

	for pos, val := range blocks {
		if val != -1 {
			result += pos * val
		}
	}

	return result
}

func removeSpace(slice []*Space, s int) []*Space {
	return append(slice[:s], slice[s+1:]...)
}
