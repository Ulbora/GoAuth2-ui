//Package handlers ...
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
)

func TestOauthRestHandlerLogger_SetDebugLogLevel(t *testing.T) {
	var oh OauthHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"debug"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "DEBUG" {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetDebugLogLevelBadReq(t *testing.T) {
	var oh OauthHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"debug"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetInfoLogLevel(t *testing.T) {
	var oh OauthHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"info"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "INFO" {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetAllLogLevel(t *testing.T) {
	var oh OauthHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"all"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "ALL" {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetOffLogLevel(t *testing.T) {
	var oh OauthHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"off"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "OFF" {
		t.Fail()
	}
}

func TestOauthRestHandlerLogger_SetOffLogLevelLogKey(t *testing.T) {
	os.Setenv("LOGGING_KEY", "45sdbb2345")
	var oh OauthHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"off"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !lres.Success || lres.LogLevel != "OFF" {
		t.Fail()
	}
	os.Unsetenv("LOGGING_KEY")
}

func TestOauthRestHandlerLogger_SetOffLogLevelLogKeyWrongKey(t *testing.T) {
	os.Setenv("LOGGING_KEY", "45sdbb23455")
	var oh OauthHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"off"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
	os.Unsetenv("LOGGING_KEY")
}

func TestOauthRestHandlerLogger_SetBadMediaLogLevel(t *testing.T) {
	var oh OauthHandler
	var logger lg.Logger
	oh.Log = &logger

	h := oh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"logLevel":"off"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Logging_KEY", "45sdbb2345")
	w := httptest.NewRecorder()
	h.SetLogLevel(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var lres LogResponse
	json.Unmarshal(body, &lres)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 415 {
		t.Fail()
	}
}

type testObj struct {
	Valid bool   `json:"valid"`
	Code  string `json:"code"`
}

func TestOauthRestHandlerLogger_ProcessBodyBadObj(t *testing.T) {
	var oh OauthHandler
	var l lg.Logger
	oh.Log = &l
	var robj testObj
	robj.Valid = true
	robj.Code = "3"
	// var res http.Response
	// res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	var sURL = "http://localhost/test"
	aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
	var obj testObj
	suc, _ := oh.ProcessBody(r, nil)
	if suc || obj.Valid != false || obj.Code != "" {
		t.Fail()
	}
}
