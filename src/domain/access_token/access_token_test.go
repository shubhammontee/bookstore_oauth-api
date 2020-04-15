package access_token

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("Brand New Access Token Should Not Be Expired")
	}
	if at.AccessToken != "" {
		t.Error("New Access Token Should Not Have A Defined Access Token")
	}

	if at.UserID != 0 {
		t.Error("New Access Token Should Not Have An Associated UserID")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("Empty Access Token Should Be Expired By Default")
	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("Access Token Expiring three Hours From Now Should Not Be Expired")
	}
}
