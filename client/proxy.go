package client

import (
	"fbreqs/console"
	"fbreqs/random"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

type Proxy struct {
	ProxyList []*url.URL
	Random    bool
}

func (proxy *Proxy) GetProxy() (*url.URL, error) {
	if len(proxy.ProxyList) == 0 {
		return nil, EmptyProxyList
	}
	if proxy.Random {
		return random.SliceValue(proxy.ProxyList).(*url.URL), nil
	}
	return proxy.ProxyList[0], nil
}

func GetProxyConfig() *Proxy {
	proxy := &Proxy{
		ProxyList: parseProxyList(viper.GetStringSlice("request.proxy.list")),
		Random:    viper.GetBool("request.proxy.random"),
	}

	if len(proxy.ProxyList) == 0 {
		console.Warn("Running without any proxy")
	}

	return proxy
}

func parseProxyList(list []string) (parsed []*url.URL) {
	for _, rawURL := range list {
		rawURL = strings.TrimSpace(rawURL)

		// Check if URL is not empty
		if rawURL == "" {
			continue
		}

		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			console.Warn("Could not parse proxy '%s': %s", rawURL, err)
			continue
		}
		parsed = append(parsed, parsedURL)
	}
	return parsed
}
