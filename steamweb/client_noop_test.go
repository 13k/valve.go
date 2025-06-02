package steamweb_test

import (
	"net/http"
	"net/url"
	"time"

	"github.com/13k/valve.go/steamweb"
)

type NoopClient struct {
	req *steamweb.Request
	res *steamweb.Response
	err error

	cookies      []*http.Cookie
	jar          http.CookieJar
	debug        bool
	header       http.Header
	hostURL      string
	logger       steamweb.Logger
	proxy        string
	query        url.Values
	retryCount   int
	retryMaxWait time.Duration
	retryWait    time.Duration
	timeout      time.Duration
	transport    http.RoundTripper
}

var _ steamweb.Client = (*NoopClient)(nil)

func (c *NoopClient) Execute(req *steamweb.Request) (*steamweb.Response, error) {
	c.req = req
	return c.res, c.err
}

func (c *NoopClient) HostURL() string {
	return c.hostURL
}

func (c *NoopClient) SetCookies(cookies []*http.Cookie) steamweb.Client {
	c.cookies = cookies
	return c
}

func (c *NoopClient) SetCookieJar(jar http.CookieJar) steamweb.Client {
	c.jar = jar
	return c
}

func (c *NoopClient) SetDebug(debug bool) steamweb.Client {
	c.debug = debug
	return c
}

func (c *NoopClient) SetHeader(key string, value string) steamweb.Client {
	if c.header == nil {
		c.header = make(http.Header)
	}

	c.header.Set(key, value)

	return c
}

func (c *NoopClient) SetHeaders(header http.Header) steamweb.Client {
	if c.header == nil {
		c.header = make(http.Header)
	}

	for k, v := range header {
		c.header[k] = v
	}

	return c
}

func (c *NoopClient) SetHostURL(host string) steamweb.Client {
	c.hostURL = host
	return c
}

func (c *NoopClient) SetLogger(logger steamweb.Logger) steamweb.Client {
	c.logger = logger
	return c
}

func (c *NoopClient) SetProxy(proxy string) steamweb.Client {
	c.proxy = proxy
	return c
}

func (c *NoopClient) SetQueryParam(key string, value string) steamweb.Client {
	if c.query == nil {
		c.query = make(url.Values)
	}

	c.query.Set(key, value)

	return c
}

func (c *NoopClient) SetQueryParams(query url.Values) steamweb.Client {
	if c.query == nil {
		c.query = make(url.Values)
	}

	for k, v := range query {
		c.query[k] = v
	}

	return c
}

func (c *NoopClient) SetRetryCount(count int) steamweb.Client {
	c.retryCount = count
	return c
}

func (c *NoopClient) SetRetryMaxWaitTime(maxWait time.Duration) steamweb.Client {
	c.retryMaxWait = maxWait
	return c
}

func (c *NoopClient) SetRetryWaitTime(wait time.Duration) steamweb.Client {
	c.retryWait = wait
	return c
}

func (c *NoopClient) SetTimeout(timeout time.Duration) steamweb.Client {
	c.timeout = timeout
	return c
}

func (c *NoopClient) SetTransport(transport http.RoundTripper) steamweb.Client {
	c.transport = transport
	return c
}
