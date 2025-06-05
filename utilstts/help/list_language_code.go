package utilhelp

import (
	"fmt"

	"github.com/hend41234/gctts/voices"
)

func HelpListLanguageCode() {
	fmt.Println("list language code :")
	unique := []string{}
	voices := voices.ListVoices
	for _, l := range voices.Voice {
		skip := false
		for _, u := range unique {
			if l.LanguageCodes[0] == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, l.LanguageCodes[0])
		}
	}

	for _, lc := range unique {
		fmt.Println("-", lc)
	}
}

func HelpListLanguageCodeName() {
	unique := []string{}
	voices := voices.ListVoices
	for _, l := range voices.Voice {
		skip := false
		for _, u := range unique {
			if l.Name == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, l.Name)
		}
	}

	for _, lc := range unique {
		fmt.Println("- ", lc)
	}
}
