package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
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

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("cookie")
	if err != nil {
		return nil, err
	}

	value := make(map[string]string)
	if err = s.Decode("cookie", cookie.Value, &value); err != nil {
		return nil, err
	}

	return value, nil
}

func Delete(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:    "cookie",
		Value:   "",
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(w, cookie)
}
