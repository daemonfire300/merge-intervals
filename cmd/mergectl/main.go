package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/daemonfire300/merge-intervals/pkg/merge"
)

func main() {
	var data io.Reader
	if len(os.Args) > 1 && os.Args[1] != "" {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		data = f
	} else {
		data = bufio.NewReader(os.Stdin)
	}

	var parsedData ExampleIntervalDataFormat
	if err := json.NewDecoder(data).Decode(&parsedData); err != nil {
		panic(err)
	}
	res := merge.Merge(parsedData)
	fmt.Println(res)
}

type ExampleIntervalDataFormat [][]int
