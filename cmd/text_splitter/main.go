package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jempe/text_splitter/splitter"
)

var chunkLength = flag.Int("chunk-length", 1000, "The length of each chunk")
var rangeLength = flag.Int("range-length", 50, "The range of characters to check for delimiters")
var delimiters = flag.String("delimiters", " ,\n,\t", "The delimiters to split the text by")

func main() {

	// Parse the flags
	flag.Parse()

	delimRunes := strings.Split(*delimiters, ",")
	var delimiters []rune
	for _, delimRune := range delimRunes {
		delimiters = append(delimiters, []rune(delimRune)[0])
	}

	textFromStdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	chunks := splitter.SplitTextInChunks(string(textFromStdin), *chunkLength, *rangeLength, delimiters)

	chunksJson, err := json.Marshal(chunks)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(chunksJson))
}
