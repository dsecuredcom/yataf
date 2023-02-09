package main

import (
	"fmt"
	"github.com/Damian89/yataf/pkg/analyze"
	"github.com/Damian89/yataf/pkg/fetch"
	"github.com/Damian89/yataf/pkg/general"
	"os"
	"strings"
)

func main() {
	Arguments := general.GetArguments()

	var Content string

	if strings.HasPrefix(Arguments.File, "http://") || strings.HasPrefix(Arguments.File, "https://") {
		Content = fetch.GetContentOfRemoteFile(Arguments.File)
	} else {
		Content = fetch.GetContentOfLocalFile(Arguments.File)
	}

	Results := analyze.AnalyzeContent(Content, Arguments.Types)

	if len(Results.Result) == 0 {
		fmt.Printf("\033[31m[i] No results found\033[0m\n")
		os.Exit(1)
	}

	fmt.Printf("\n")

	for _, Result := range Results.Result {
		fmt.Printf("\033[34m[i] Name: %s\033[0m\n", Result.RegEx.Name)
		fmt.Printf("\033[36m[i] RegEx:\033[0m %s\n", Result.RegEx.Matcher)
		for Key, DataExtracted := range Result.DataExtracted {
			AddLine := "\r\n"
			if Result.RegEx.Type == "Urls-Paths" || len(DataExtracted) < 150 {
				AddLine = " "
			}
			if len(DataExtracted) < Arguments.CharLimit {
				fmt.Printf("\033[32m[%d]\033[0m\033[36m[DataExtracted]\033[0m%s%s\n", Key, AddLine, DataExtracted)
				continue
			}

			fmt.Printf("\033[32m[%d]\033[0m\033[36m[DataExtracted (limited to %d chars)]\033[0m%s%s\n", Key, Arguments.CharLimit, AddLine, DataExtracted[:Arguments.CharLimit])

		}
		fmt.Printf("\n")
	}

}
