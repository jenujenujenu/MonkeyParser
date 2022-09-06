package MonkeyParser

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type Document struct {
	Elements []interface{}
}

func FromFile(f fs.File) (doc Document) {

	arr := make([]string, 0)
	buffer := strings.Builder{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//line := scanner.Text()
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			arr = append(arr, buffer.String())
			buffer.Reset()
		} else {
			buffer.WriteString(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for _, el := range arr {
		doc.ParseElement(el)
	}

	return
}
