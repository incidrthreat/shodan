package restapi

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/incidrthreat/shodan/config"
	"github.com/incidrthreat/shodan/net/http"
	"github.com/incidrthreat/shodan/net/httputil"
)

type Scan interface {
	Protocols(ctx context.Context) (string, error)
	Scan(ctx context.Context, ips []string) (string, error)
	ScanInternet(ctx context.Context, port int, protocol string) (string, error)
	ScanStatus(ctx context.Context, id string) (string, error)
}

type scan struct {
	key    string
	config config.Scan
}

func (scan *scan) Protocols(ctx context.Context) (string, error) {
	url := scan.config.ProtocolsURL

	options := make(map[string]string)
	options[config.KEY] = scan.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (scan *scan) Scan(ctx context.Context, ips []string) (string, error) {
	url := scan.config.ScanURL

	data := url.Values{}
	data.set("key", scan.key)
	data.Add("ips", ips)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s", url), bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	response, e := http.Do(ctx, MethodPost, req, nil)
	return httputil.Response(response, e)
}

func (scan *scan) ScanInternet(ctx context.Context, port int, protocol string) (string, error) {
	panic("implement me")
}

func (scan *scan) ScanStatus(ctx context.Context, id string) (string, error) {
	url := scan.config.ScanStatusURL
	furl := strings.Replace(url, "{id}", id, -1)

	options := make(map[string]string)
	options[config.KEY] = scan.key

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}
