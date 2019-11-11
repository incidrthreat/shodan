package restapi

import (
	"context"
	. "net/http"
	"github.com/incidrthreat/shodan/config"
	"github.com/incidrthreat/shodan/net/http"
	"github.com/incidrthreat/shodan/net/httputil"

	"strings"
)

type Experimental interface {
	Honeyscore(ctx context.Context, ip string) (string, error)
}

type experimental struct {
	key    string
	config config.Experimental
}

func (experimental *experimental) Honeyscore(ctx context.Context, ip string) (string, error) {
	url := experimental.config.HoneyscoreURL
	furl := strings.Replace(url, "{ip}", ip, -1)

	options := make(map[string]string)
	options[config.KEY] = experimental.key

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}
