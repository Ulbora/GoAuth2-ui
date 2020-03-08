//Package handlers ...
package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOauthHandler_HandleIndex(t *testing.T) {
	var oh OauthHandler
	oh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	h := oh.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleIndex(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}
