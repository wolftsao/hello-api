// Package translation houses logic to take a given word and find it in a different language
package translation

import "strings"

// StaticService has data that does not change.
type StaticService struct{}

// NewStaticService creates anew instance of a service that uses static data.
func NewStaticService() *StaticService {
	return &StaticService{}
}

// Translate a given word to the passed language in language.
func (s *StaticService) Translate(word string, language string) string {
	word = sanitizeInput(word)
	language = sanitizeInput(language)

	if word != "hello" {
		return ""
	}

	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	case "french":
		return "bonjour"
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
