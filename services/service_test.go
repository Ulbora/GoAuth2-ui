//Package services ...
package services

import (
	"testing"
)

func TestOauth2Service_SetToken(t *testing.T) {
	var s Oauth2Service
	s.SetToken("ttttt")
	if s.Token != "ttttt" {
		t.Fail()
	}
}
