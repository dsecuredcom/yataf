package analyze

import (
	"github.com/Damian89/FileAnalyzer/pkg/config"
	_struct "github.com/Damian89/FileAnalyzer/pkg/struct"
	"regexp"
	"strings"
)

func AnalyzeContent(Content string, Types []string) _struct.Results {

	RegularExpressionConfig := config.GetRegExConfig(Types)
	NegativeExtractions := config.GetNegativeExtractions()

	var Results _struct.Results
	FoundExtractions := make(map[string]bool)

	for _, RegEx := range RegularExpressionConfig.Elements {

		RegExMatcher := regexp.MustCompile(RegEx.Matcher)
		RegExResults := RegExMatcher.FindAllString(Content, -1)

		if len(RegExResults) == 0 {
			continue
		}

		CleanedExtractedResults := make([]string, 0)

		for _, Extraction := range RegExResults {
			if Contains(NegativeExtractions, Extraction) {
				continue
			}

			if FoundExtractions[Extraction] {
				continue
			}

			FoundExtractions[Extraction] = true

			CleanedExtractedResults = append(CleanedExtractedResults, Extraction)
		}

		if len(CleanedExtractedResults) == 0 {
			continue
		}

		Results.Result = append(Results.Result, _struct.Result{
			RegEx:         RegEx,
			DataExtracted: CleanedExtractedResults,
		})
	}

	return Results
}

func Contains(Extractions []string, Extraction string) bool {
	for _, NegativeExtraction := range Extractions {

		if strings.TrimSpace(NegativeExtraction) == strings.TrimSpace(Extraction) {
			return true
		}
	}

	return false
}
