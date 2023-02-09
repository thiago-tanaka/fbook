package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

func MakeRequestWithAuth(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Read(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
