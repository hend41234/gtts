package text

import (
	"regexp"
)

func ToSSML(input string) string {
	replacements := []struct {
		pattern string
		replace string
	}{
		{`\.{3}`, `<break time="600ms"/>`},                                   // elipsis
		{`—`, `<break time="300ms"/>`},                                       // em dash
		{`\?`, `<prosody pitch="+2st">?</prosody><break time="700ms"/>`},     // pertanyaan
		{`\!`, `<emphasis level="strong">!</emphasis><break time="700ms"/>`}, // seruan
		{`,`, `<break time="300ms"/>`},                                       // koma
		{`\.\s`, `<break time="500ms"/>`},                                    // titik akhir kalimat
	}

	text := input
	for _, r := range replacements {
		re := regexp.MustCompile(r.pattern)
		text = re.ReplaceAllString(text, r.replace)
	}

	// Bungkus jadi SSML
	ssml := `<speak><voice name="en-GB-Neural2-D"><prosody rate="medium" pitch="-1st">` +
		text +
		`</prosody></voice></speak>`
	return ssml
}
