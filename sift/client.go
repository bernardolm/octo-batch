package sift

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/bernardolm/octo-batch/debug"
	"github.com/bernardolm/octo-batch/oauth"
)

var client *http.Client

func newClient() (*http.Client, error) {
	oauth2Transport, err := oauth.GetOAuth2Transport(tokenData)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Transport: oauth2Transport,
	}

	return c, nil
}

func getClient() *http.Client {
	if client != nil {
		return client
	}

	c, err := newClient()
	if err != nil {
		return nil
	}

	return c
}

type page struct {
	Page     int `url:"page"`
	PageSize int `url:"pageSize"`
}

type Meta struct {
	Pages struct {
		First int
		Last  int
		Next  int
		Prev  int
		Size  int
	}
	TotalLength int
}

type Response struct {
	Data  []byte `json:"-"`
	Links map[string]string
	Meta  Meta
}

func get(ctx context.Context, uri string) (*Response, error) {
	url := fmt.Sprintf("%s/%s", apiURL, uri)
	// debug.Print("get.url", url)

	now := time.Now()

	resp, err := getClient().Get(url)
	if err != nil {
		return nil, err
	}

	debug.Print("request elapsed", time.Since(now))

	// debug.Print("get.url.resp.Request.Header", resp.Request.Header)
	// debug.Print("get.url.resp.Status", resp.Status)

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, err
	}
	resp.Body.Close()

	res := Response{
		Data: body,
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	// debug.Print("get.res.Data", len(res.Data))
	// debug.Print("get.res.Links", res.Links)
	// debug.Print("get.res.Meta", res.Meta)

	return &res, nil
}

func paginate(ctx context.Context, uri string, p *page) (*Response, error) {
	uv, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	// debug.Print("paginate.uv", uv)

	qs := uv.Encode()
	// debug.Print("paginate.qs", qs)

	url := fmt.Sprintf("%s?%s", uri, qs)
	// debug.Print("paginate.url", url)

	res, err := get(ctx, url)
	if err != nil {
		return nil, err
	}
	// debug.Print("paginate.res", res)

	return res, nil
}
