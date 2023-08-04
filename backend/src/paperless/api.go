package paperless

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"sapp/paperless-accounting/config"
	"strconv"
)

func paperlessQueryExecute[R paperlessResponse, RR paperlessResponseResult](conf *config.Config, query string, apiPath string, allPages bool) ([]RR, error) {
	uri, err := url.JoinPath(conf.PAPERLESS_URL, apiPath)

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
	request.Header.Add("Authorization", "Token "+conf.PAPERLESS_AUTH_TOKEN)

	pageNumber := 1
	// will be the output
	var out []RR
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
		var res_json R
		err = json.NewDecoder(res.Body).Decode(&res_json)
		if err != nil {
			return nil, errors.New("Could not read response body: " + err.Error() + "\n")
		}

		// retrieve fields
		val := reflect.ValueOf(res_json)
		countV := val.FieldByName("Count")
		if countV.Kind() == 0 {
			return nil, errors.New("invalid format, could not find field \"Count\"")
		}
		count := countV.Int()
		nextV := val.FieldByName("Next")
		if countV.Kind() == 0 {
			return nil, errors.New("invalid format, could not find field \"Next\"")
		}
		next := nextV.String()
		resultsV := val.FieldByName("Results")
		if resultsV.Kind() == 0 {
			return nil, errors.New("invalid format, could not find field \"Results\"")
		}
		var results []RR
		if resultsV.Kind() == reflect.Slice {
			results = resultsV.Interface().([]RR)
		}

		// init array?
		if out == nil {
			out = make([]RR, count)
		}

		// copy results
		outPositionEnd := outPositionStart + len(results)
		copy(out[outPositionStart:outPositionEnd], results)
		outPositionStart = outPositionEnd

		if !allPages || next == "" {
			hasNext = false
		} else {
			pageNumber++
		}
	}

	return out, nil
}

func (p *Paperless) paperlessDocumentQuery(query string) ([]PaperlessDocument, error) {
	//var out []PaperlessDocument

	out, err := paperlessQueryExecute[paperlessDocumentResponse, PaperlessDocument](p.conf, query, "/api/documents/", true)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (p *Paperless) paperlessCorrespondentList() ([]PaperlessCorrespondent, error) {
	//var out []PaperlessDocument

	out, err := paperlessQueryExecute[paperlessCorrespondentResponse, PaperlessCorrespondent](p.conf, "", "/api/documents/", true)

	if err != nil {
		return nil, err
	}

	return out, nil
}
