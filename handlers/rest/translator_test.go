package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shipping-go/handlers/rest"
	"testing"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{"/hello", http.StatusOK, "english", "hello"},
		{"/hello/?language=german", http.StatusOK, "german", "hallo"},
		{"/hello/?language=french", http.StatusNotFound, "french", "unsupported language"},
		{"/book/?language=german", http.StatusNotFound, "german", "unsupported word"},
	}
	handler := http.HandlerFunc(rest.TranslateHandler)
	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", test.Endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(rr, req)
		if rr.Code != test.StatusCode {
			t.Errorf("expected status code %d, actual %d", test.StatusCode, rr.Code)
		}
		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)
		if resp.Language != test.ExpectedLanguage {
			t.Errorf("expected language %s, actual %s", test.ExpectedLanguage, resp.Language)
		}
		if resp.Translation != test.ExpectedTranslation {
			t.Errorf("expected translation %s, actual %s", test.ExpectedTranslation, resp.Translation)
		}
	}

}
