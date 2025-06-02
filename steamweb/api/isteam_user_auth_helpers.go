package api

import (
	"crypto/aes"
	"crypto/rand"
	"net/url"

	"github.com/13k/valve.go/steamcrypto"
	"github.com/13k/valve.go/steamid"
	"github.com/13k/valve.go/steamres/steamlang"
)

// ISteamUserAuthAuthenticateUserFormData generates request form data for method
// ISteamUserAuth/AuthenticateUser.
func ISteamUserAuthAuthenticateUserFormData(
	steamID steamid.SteamID,
	loginKey string,
) (url.Values, error) {
	sessionKey := make([]byte, 32)

	if _, err := rand.Read(sessionKey); err != nil {
		return nil, err
	}

	encryptedSessionKey, err := steamcrypto.RSAEncrypt(steamcrypto.GetPublicKey(steamlang.EUniverse_Public), sessionKey)

	if err != nil {
		return nil, err
	}

	ciph, err := aes.NewCipher(sessionKey)

	if err != nil {
		return nil, err
	}

	encryptedLoginKey, err := steamcrypto.SymmetricEncrypt(ciph, []byte(loginKey))

	if err != nil {
		return nil, err
	}

	values := url.Values{}

	values.Set("steamid", steamID.FormatString())
	values.Set("sessionkey", string(encryptedSessionKey))
	values.Set("encrypted_loginkey", string(encryptedLoginKey))

	return values, nil
}
