//Package handlers ...
package handlers

import (
	"fmt"
	"net/http"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
)

func TestOauthHandler_getRedirectURI(t *testing.T) {
	var h OauthHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	uri := h.getRedirectURI(r, "/dosometesting")
	if uri != "https://test.com/dosometesting" {
		t.Fail()
	}
}

func TestOauthHandler_getRedirectURI2(t *testing.T) {
	var h OauthHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	h.OauthHost = "test.com"
	h.SchemeDefault = "https://"
	r, _ := http.NewRequest("POST", "/test.com", nil)
	r.Host = "test.com"
	fmt.Println("req host: ", r.Host)
	uri := h.getRedirectURI(r, "/dosometesting")
	if uri != "https://test.com/dosometesting" {
		t.Fail()
	}
}
