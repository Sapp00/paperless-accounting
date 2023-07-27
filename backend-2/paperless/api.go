package paperless

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sapp/paperless-accounting/config"
)

type Paperless struct {
	conf config.Config
}

func (p *Paperless) paperlessDocumentQueryExecute(query string, pageNumber int) ([]PaperlessDocument, error) {
	uri, err := url.JoinPath(p.conf.PAPERLESS_URL, "/api/documents/?page="+string(pageNumber)+"&query="+query)
	if err != nil {
		return nil, errors.New("Invalid URL: " + err.Error())
	}

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.New("Cannot create request: " + err.Error())
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json;charset=UTF-8")
	request.Header.Add("Authorization", "Token "+p.conf.PAPERLESS_AUTH_TOKEN)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, errors.New("Error sending request: " + err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Got unexpected Status Code: " + string(res.StatusCode))
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Could not read response body: " + err.Error() + "\n")
	}

	fmt.Printf("client: response body: %s\n", resBody)

	var out [5]PaperlessDocument

	return out, nil
}

func (p *Paperless) PaperlessDocumentQuery(query string) []PaperlessDocument {
	var out []PaperlessDocument

	return out
}
