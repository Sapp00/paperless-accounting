package paperless

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func (p *Paperless) GetDocumentThumb(id int) ([]byte, error) {
	uri, err := url.JoinPath(p.conf.PAPERLESS_URL, "/api/documents/"+strconv.Itoa(id)+"/thumb/")

	if err != nil {
		return nil, errors.New("Invalid URL: " + err.Error())
	}

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.New("Cannot create request: " + err.Error())
	}

	client := &http.Client{}

	request.Header.Add("Authorization", "Token "+p.conf.PAPERLESS_AUTH_TOKEN)

	res, err := client.Do(request)
	if err != nil {
		return nil, errors.New("Error sending request: " + err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Got unexpected Status Code: " + strconv.Itoa(res.StatusCode))
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
