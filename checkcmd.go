package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var lines []string
var cmd []string

func main() {
	filepath.WalkDir(".", visit)
	basket := dup_count(lines)
	keys := make([]int, 0, len(basket))

	for _, k := range basket {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for cmd := range basket {
		if basket[cmd] == keys[0] {
			fmt.Println(cmd, "used", keys[0], "times")
		}
	}

}

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func visit(path string, di fs.DirEntry, err error) error {

	if strings.Contains(path, ".cfg") {

		f, _ := os.Open(path)
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {


			if strings.Contains(strings.ToLower(scanner.Text()), strings.ToLower(os.Args[1])) {
				lines = append(lines, scanner.Text())

			}

		}

	}

	return nil
}
