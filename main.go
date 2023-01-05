package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	files := []string{
		"redis-cli-5.0.14",
		"redis-cli-6.2.8",
		"redis-cli-7.0.7",
	}

	for _, file := range files {
		fileContent, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		content := string(fileContent)

		regex := `(?m)(?P<version>\d+\.\d+\.\d+)\W\W\(git:%s\W-dirty`
		r := regexp.MustCompile(regex)
		allMatches := r.FindStringSubmatch(content)
		fmt.Println(allMatches)
	}
}
