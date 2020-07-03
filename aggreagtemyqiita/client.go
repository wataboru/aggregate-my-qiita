// Package aggreagtemyqiita is logic package
package aggreagtemyqiita

import (
	"fmt"
	"net/http"
	"strconv"
)

type client struct {
	http.Client
	token string
}

func (client *client) do(r *http.Request) (*http.Response, error) {

	res, err := client.Do(r)

	if err != nil {
		return &http.Response{}, fmt.Errorf("[ERR]: %s", err)
	}

	switch res.StatusCode {
	case 200, 201, 204:
	case 400:
		message := "StatusCode: 400 Bad Request"
		return &http.Response{}, fmt.Errorf("[ERR]: %s", message)
	case 401:
		message := "StatusCode: 401 Unauthorized"
		return &http.Response{}, fmt.Errorf("[ERR]: %s", message)
	case 403:
		message := "StatusCode: 403 Forbidden"
		return &http.Response{}, fmt.Errorf("[ERR]: %s", message)
	case 404:
		message := "StatusCode: 404 Not Found"
		return &http.Response{}, fmt.Errorf("[ERR]: %s", message)
	case 500:
		message := "StatusCode: 500 Internal Server Error"
		return &http.Response{}, fmt.Errorf("[ERR]: %s", message)
	default:
		message := "StatusCode: " + strconv.Itoa(res.StatusCode) + " Unknown Error occurred"
		return &http.Response{}, fmt.Errorf("[ERR] :%s", message)
	}

	return res, nil
}

func (client *client) request(url string) (*http.Response, error) {
	//fmt.Printf("[INFO]: %s\n", "Request to "+url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, fmt.Errorf("[ERR] :%s", err)
	}

	request.Header.Set("Authorization", "Bearer "+client.token)

	res, err := client.do(request)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (client *client) parallelRequest(pageDetailItemCh chan pageDetailItem, url string) error {
	//fmt.Printf("[INFO]: %s\n", "Request to "+url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("[ERR] :%s", err)
	}

	request.Header.Set("Authorization", "Bearer "+client.token)

	res, err := client.do(request)
	if err != nil {
		return err
	}

	var pageDetailItem pageDetailItem
	decodeBody(res, &pageDetailItem)

	pageDetailItemCh <- pageDetailItem
	return nil
}
