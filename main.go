package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	
)


type CommandInput struct {
	DoByteCount *bool
	DoLineCount *bool
	DoWordCount *bool
}
type Results struct {
	ByteCount int64
	LineCount int
	WordCount int
	CharCount int
}

func WcTool(input CommandInput, filePath string) (*Results, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	results := &Results{ByteCount: 0, LineCount: 0, WordCount: 0, CharCount: 0}
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	filter := func (r rune) bool {
		return unicode.IsLetter(r)  || unicode.IsNumber(r) || unicode.IsSymbol(r) || unicode.IsSpace(r) || unicode.IsGraphic(r)
	}
	results.ByteCount = info.Size()
	reader := bufio.NewReader(f)
	for  {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF{
				break
			}else{
				return nil,err
			}
		}
		results.LineCount ++
		trimLine := strings.TrimFunc(line, func(r rune) bool {return !unicode.IsLetter(r) && !unicode.IsSpace(r)})
		results.WordCount += len(strings.Split(trimLine, " "))
		for _, ch := range line {
			if filter(ch) {
				results.CharCount ++
			}
		}
		
		// results.CharCount += len(filterChars)
	}
	

	
	return results,nil
}

func main() {
	
	input := CommandInput{
		DoByteCount: flag.Bool("c", false, "add flag -c to count the number of bytes in the file"),
		DoLineCount: flag.Bool("l", false, "add flag -l to count the number of lines in the file"),
		DoWordCount:flag.Bool("w", false, "add flag -w to count the number of words in the file") ,
	}
	flag.Parse()
	r, err := WcTool(input, flag.Arg(0))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes: %d, Lines: %d, Words: %d, Chars: %d\n", r.ByteCount, r.LineCount, r.WordCount, r.CharCount)
}