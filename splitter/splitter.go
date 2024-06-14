package splitter

import (
	"bytes"
)

func SplitTextInChunks(text string, chunkLength int, rangeLength int, delimiters []rune) []string {
	var chunks []string
	current := new(bytes.Buffer)

	for _, char := range text {
		isDelimiter := false

		for _, delimiter := range delimiters {
			if char == delimiter {
				isDelimiter = true
				break
			}
		}

		if (isDelimiter && current.Len() >= chunkLength) || current.Len() >= chunkLength+rangeLength {
			chunks = append(chunks, current.String())
			current.Reset()
		}

		current.WriteRune(char)
	}

	// Add the remaining text to the last chunk
	chunks = append(chunks, current.String())

	return chunks
}
