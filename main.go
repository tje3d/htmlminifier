package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	p   = flag.String("path", "./storage/framework/views", "Path to optimize")
	ext = flag.String("ext", ".php", "File extension")
)

func main() {
	flag.Parse()

	err := filepath.Walk(
		*p,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}

			if strings.HasSuffix(path, *ext) {
				data, _ := ioutil.ReadFile(path)
				var re = regexp.MustCompile(`(?m)(^\s+)`)
				replaced := re.ReplaceAllString(string(data), ``)

				ioutil.WriteFile(path, []byte(replaced), 0644)
			}

			return nil
		})

	if err != nil {
		panic(err)
	}
}
