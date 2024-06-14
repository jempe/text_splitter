## Text Splitter - Text Chunking Library

This Go library provides a function `SplitTextInChunks` to split a string into smaller chunks based on a desired length and optional delimiters.

### Installation

There are two ways to use this library:

1. **As a standalone package:**

   Clone this repository and run `go install ./...` in the project directory.

2. **As a dependency:**

   1. Add the following line to your `go.mod` file:

      ```
      replace github.com/your-username/splitter v0.0.1 => ./splitter
      ```

      Replace `github.com/your-username/splitter` with the actual import path for your library and `v0.0.1` with the desired version number (if applicable).

   2. Import the package in your Go code:

      ```go
      import "github.com/jempe/text_splitter/splitter"
      ```

### Usage

The `SplitTextInChunks` function takes four arguments:

* `text` (string): The text to be split.
* `chunkLength` (int): The desired length of each chunk.
* `rangeLength` (int): An optional parameter that specifies a tolerance range around the `chunkLength`. Chunks will be created even if they are slightly longer than `chunkLength` as long as they fall within the `chunkLength` plus `rangeLength`. Set this to 0 for strict adherence to `chunkLength`.
* `delimiters` ([]rune): An optional slice of runes (characters) that will be treated as chunk separators. When a delimiter is encountered, a new chunk will be created even if the current chunk hasn't reached the `chunkLength`.

The function returns a slice of strings, where each element is a chunk of the original text.

**Example:**

```go
package main

import (
	"fmt"
	"github.com/jempe/text_splitter/splitter"
)

func main() {
	text := "This is a long string that needs to be split into smaller chunks. We can define a desired chunk length and optionally specify delimiters."
	chunkLength := 20
	rangeLength := 5  // Allow chunks to be 5 characters shorter or longer than chunkLength
	delimiters := []rune{'.', ','}

	chunks := splitter.SplitTextInChunks(text, chunkLength, rangeLength, delimiters)

	for _, chunk := range chunks {
		fmt.Println(chunk)
	}
}
```

This example will split the text into chunks with a maximum length of 25 characters (20 + 5) and treat "." and "," as chunk separators.

### Contributing

We welcome contributions to this library! Please feel free to submit pull requests with improvements or bug fixes.

