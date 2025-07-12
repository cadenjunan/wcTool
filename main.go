package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func readBytes(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error while reading %s, cause: %s", path, err.Error())
	}
	fmt.Printf("%d %s\n", len(data), path)
	return nil
}
func lineCount(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error while opening %s, cause: %s", path, err.Error())
	}
	reader := bufio.NewReader(f)
	lines := 0
	cond: for {
		_, err := reader.ReadString('\n')
		if err == io.EOF {
			break cond
		}
		lines ++
	}
	fmt.Printf("%d %s\n",lines, path)
	return nil
}

func main() {
	byteCountCmd := flag.Bool("c", false, "add flag -c to count the number of bytes in the file")
	lineCountCmd := flag.Bool("l", false, "add flag -l to count the number of bytes in the file")
	flag.Parse()
	if *byteCountCmd {
		err := readBytes(flag.Arg(0))
		if err != nil {
			panic(err.Error())
		}
	}else if *lineCountCmd {
		lineCount(flag.Arg(0))
	}
}