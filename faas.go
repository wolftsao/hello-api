// Package faas is used for function definitions
package faas

import (
	"net/http"

	"github.com/wolftsao/hello-api/handlers/rest"
	"github.com/wolftsao/hello-api/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.NewTranslateHandler(translation.NewStaticService()).TranslateHandler(w, r)
	// rest.TranslateHandler(w, r)
}
