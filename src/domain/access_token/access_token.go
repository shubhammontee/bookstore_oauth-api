package access_token

import "time"

const (
	expirationTime = 24
)

//AccessToken Struct
type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	ClientID    int64  `json:"client_id,omitempty"`
	Expires     int64  `json:"expires,omitempty"`
}

//GetNewAccessToken ...
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

//IsExpired ...
func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)

	return false
}
