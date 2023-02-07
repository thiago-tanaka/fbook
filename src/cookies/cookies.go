package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"webapp/src/config"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New(
		config.HashKey,
		config.BlockKey,
	)
}

func Save(w http.ResponseWriter, ID, token string) error {
	value := map[string]string{
		"ID":    ID,
		"token": token,
	}

	if encoded, err := s.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:     "cookie",
			Value:    encoded,
			Path:     "/",
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)

	}
	return nil
}
