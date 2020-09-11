package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fset := token.NewFileSet()

	fh, err := os.Open("../cmd/govim/config/config.go")
	if err != nil {
		log.Print(err)
	}
	f, err := parser.ParseFile(fset, "", fh, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, c := range f.Comments {
		s := c.Text()
		if strings.HasPrefix(s, "Command") {
			s = strings.Replace(s, "Command", "GOVIM", 1)

			re := regexp.MustCompile(`(?s)^(GOVIM.*? )(.*)`)

			match := re.FindStringSubmatch(s)

			if len(match) > 2 {
				cmd := strings.TrimSpace(match[1])
				cmdlink := "*:" + cmd + "*"
				cmdDesc := match[2]
				fmt.Printf("%+78s\n", cmdlink)
				fmt.Printf(":%s\n", cmd)
				fmt.Println(cmdDesc)
			}

		}

	}

}
