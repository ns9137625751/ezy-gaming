package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

const CookieName = "admin_session"
const CustomerCookieName = "customer_session"

type SessionUser struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func Sign(user SessionUser, secret string) (string, error) {
	payload, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	b64 := base64.RawURLEncoding.EncodeToString(payload)
	sig := sign(b64, secret)
	return b64 + "." + sig, nil
}

func Verify(token, secret string) (*SessionUser, error) {
	parts := strings.SplitN(token, ".", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("malformed token")
	}
	b64, sig := parts[0], parts[1]
	if sign(b64, secret) != sig {
		return nil, fmt.Errorf("invalid signature")
	}
	payload, err := base64.RawURLEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	var u SessionUser
	if err := json.Unmarshal(payload, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func sign(data, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
