package main

import (
	"flag"
	"fmt"
	"os"
)

func readBytes(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error while reading %s, cause: %s", path, err.Error())
	}
	fmt.Printf("%d %s", len(data), path)
	return nil
}

func main() {
	byteCountCmd := flag.Bool("c", false, "add flag -c to count the number of bytes in the file")
	flag.Parse()
	if *byteCountCmd {
		err := readBytes(flag.Arg(0))
		if err != nil {
			panic(err.Error())
		}
	}
}