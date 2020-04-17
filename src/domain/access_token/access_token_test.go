package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	// if expirationTime != 24 {
	// 	t.Error("expiration time should be 24 hours")
	// }
	//we can use assert of testify package to all this in one line
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	// if at.IsExpired() {
	// 	t.Error("Brand New Access Token Should Not Be Expired")
	// }
	//this can be replaced by

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	// if at.AccessToken != "" {
	// 	t.Error("New Access Token Should Not Have A Defined Access Token")
	// }
	//this can be replaced by

	assert.EqualValues(t, "", at.AccessToken, "New Access Token Should Have A Defined Access Token")

	// if at.UserID != 0 {
	// 	t.Error("New Access Token Should Not Have An Associated UserID")
	// }
	//this can be replaced by

	assert.EqualValues(t, 0, at.UserID, "New Access Token Should Not Have Ab Associated UserID")
	//ikt can alse be replaced by the below code
	//assert.True(t, at.UserID == 0, "New Access Token Should Not Have An Associated UserID")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	// if !at.IsExpired() {
	// 	t.Error("Empty Access Token Should Be Expired By Default")
	// }
	assert.True(t, at.IsExpired(), "Empty Access Token Should Be Expired By Default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	// if at.IsExpired() {
	// 	t.Error("Access Token Expiring three Hours From Now Should Not Be Expired")
	// }
	assert.False(t, at.IsExpired(), "Access Token Expiring three Hours From Now Should Not Be Expired")
}
