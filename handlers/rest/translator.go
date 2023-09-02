package rest

import (
	"encoding/json"
	"net/http"
	"shipping-go/translation"
	"strings"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

const defaultLanguage = "english"

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	if word == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if word != "hello" {
		w.WriteHeader(http.StatusNotFound)
	}

	translation := translation.Translate(word, language)

	if translation == "unsupported language" {
		w.WriteHeader(http.StatusNotFound)
	}
	resp := Resp{
		Language:    language,
		Translation: translation,
	}
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
