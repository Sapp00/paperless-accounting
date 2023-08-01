package paperless

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sapp/paperless-accounting/config"
	"strconv"
)

type Paperless struct {
	conf *config.Config
}

func (p *Paperless) paperlessDocumentQueryExecute(query string, allPages bool) ([]PaperlessDocument, error) {
	uri, err := url.JoinPath(p.conf.PAPERLESS_URL, "/api/documents/")

	if err != nil {
		return nil, errors.New("Invalid URL: " + err.Error())
	}

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.New("Cannot create request: " + err.Error())
	}

	client := &http.Client{}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json;charset=UTF-8")
	request.Header.Add("Authorization", "Token "+p.conf.PAPERLESS_AUTH_TOKEN)

	pageNumber := 1
	// will be the output
	var out []PaperlessDocument
	outPositionStart := 0
	hasNext := true
	for hasNext {

		reqQuery := request.URL.Query()
		reqQuery.Set("page", strconv.Itoa(pageNumber))
		reqQuery.Set("query", query)
		request.URL.RawQuery = reqQuery.Encode()

		res, err := client.Do(request)
		if err != nil {
			return nil, errors.New("Error sending request: " + err.Error())
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			return nil, errors.New("Got unexpected Status Code: " + strconv.Itoa(res.StatusCode))
		}

		// parse json
		var res_json paperlessDocumentResponse
		err = json.NewDecoder(res.Body).Decode(&res_json)
		if err != nil {
			return nil, errors.New("Could not read response body: " + err.Error() + "\n")
		}

		// init array?
		if out == nil {
			out = make([]PaperlessDocument, res_json.Count)
		}

		// copy results
		outPositionEnd := outPositionStart + len(res_json.Results)
		copy(out[outPositionStart:outPositionEnd], res_json.Results)
		outPositionStart = outPositionEnd

		if !allPages || res_json.Next == "" {
			hasNext = false
		} else {
			pageNumber++
		}
	}

	return out, nil
}

func (p *Paperless) PaperlessDocumentQuery(query string) ([]PaperlessDocument, error) {
	//var out []PaperlessDocument

	out, err := p.paperlessDocumentQueryExecute(query, true)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func Init(conf *config.Config) (*Paperless, error) {
	p := Paperless{
		conf: conf,
	}

	return &p, nil
}
