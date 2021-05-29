package http

import (
	httpHeaders "github.com/go-http-utils/headers"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cast"
	"io"
	"net/url"
	"sync"
)

const (
	ContentTypeApplicationJson           = "application/json"
	ContentTypeApplicationXml            = "application/xml"
	ContentTypeApplicationFormUrlencoded = "application/x-www-form-urlencoded"
	HeaderValueAcceptAll                 = "*/*"
	HeaderValueAcceptApplicationJson     = ContentTypeApplicationJson
	HeaderValueAcceptApplicationXml      = ContentTypeApplicationXml
	HeaderValueAcceptTextCsv             = "text/csv"
	HeaderValueAcceptTextPlain           = "text/plain"
	HeaderValueAuthorizationTypeBasic    = "Basic"
	HeaderValueAuthorizationTypeBearer   = "Bearer"
	HeaderValueAuthorizationTypeDigest   = "Digest"
)

type Request struct {
	errs         error
	outputFile   *string
	queryParams  url.Values
	restyRequest *resty.Request
	url          *url.URL
}

var r struct {
	instance *resty.Client
	lck      sync.Mutex
}

// use NewRequest(client) or client.NewRequest() to create a request, don't create the object inline!
func NewRequest(client Client) *Request {
	r.lck.Lock()
	defer r.lck.Unlock()

	if r.instance == nil {
		r.instance = resty.New()
	}

	if client == nil {
		return &Request{
			restyRequest: r.instance.NewRequest(),
			url:          &url.URL{},
			queryParams:  url.Values{},
		}
	}

	return client.NewRequest()
}

// use NewJsonRequest to create a request that already contains the application/json content-type, don't create the object inline!
func NewJsonRequest(client Client) *Request {
	return NewRequest(client).
		WithHeader(httpHeaders.Accept, ContentTypeApplicationJson)
}

// use NewXmlRequest to create a request that already contains the application/xml content-type, don't create the object inline!
func NewXmlRequest(client Client) *Request {
	return NewRequest(client).
		WithHeader(httpHeaders.Accept, ContentTypeApplicationXml)
}

func (r *Request) WithUrl(rawUrl string) *Request {
	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		r.errs = multierror.Append(r.errs, err)

		return r
	}

	r.addQueryValues(parsedUrl.Query())

	parsedUrl.RawQuery = ""

	r.url = parsedUrl

	return r
}

func (r *Request) WithQueryParam(key string, values ...interface{}) *Request {
	for _, value := range values {
		str, err := cast.ToStringE(value)

		if err != nil {
			r.errs = multierror.Append(r.errs, err)

			continue
		}

		r.queryParams.Add(key, str)
	}

	return r
}

func (r *Request) WithQueryObject(obj interface{}) *Request {
	parts, err := query.Values(obj)

	if err != nil {
		r.errs = multierror.Append(r.errs, err)

		return r
	}

	r.addQueryValues(parts)

	return r
}

func (r *Request) WithQueryMap(values interface{}) *Request {
	parts, err := cast.ToStringMapStringSliceE(values)

	if err != nil {
		r.errs = multierror.Append(r.errs, err)

		return r
	}

	r.addQueryValues(parts)

	return r
}

func (r *Request) WithBasicAuth(username string, password string) *Request {
	r.restyRequest.SetBasicAuth(username, password)

	return r
}

func (r *Request) WithAuthToken(token string) *Request {
	r.restyRequest.SetAuthToken(token)

	return r
}

func (r *Request) WithHeader(key string, value string) *Request {
	r.restyRequest.SetHeader(key, value)

	return r
}

func (r *Request) WithBody(body interface{}) *Request {
	r.restyRequest.SetBody(body)

	return r
}

func (r *Request) WithMultipartFile(param, fileName string, reader io.Reader) *Request {
	r.restyRequest.SetFileReader(param, fileName, reader)

	return r
}

func (r *Request) WithMultipartFormData(params url.Values) *Request {
	r.restyRequest.SetFormDataFromValues(params)

	return r
}

func (r *Request) WithOutputFile(path string) *Request {
	r.outputFile = &path

	return r
}

// The following methods are mainly intended for tests
// You should not need to call them yourself

type Header map[string][]string

func (r *Request) GetHeader() Header {
	header := make(Header)

	// make a copy of our headers to prevent a caller
	// from modifying them
	for key, values := range r.restyRequest.Header {
		header[key] = append(make([]string, 0, len(values)), values...)
	}

	return header
}

func (r *Request) GetBody() interface{} {
	return r.restyRequest.Body
}

func (r *Request) GetToken() string {
	return r.restyRequest.Token
}

func (r *Request) GetUrl() string {
	r.url.RawQuery = r.queryParams.Encode()

	return r.url.String()
}

func (r *Request) GetError() error {
	return r.errs
}

func (r *Request) build() (*resty.Request, string, error) {
	if r.errs != nil {
		return nil, "", r.errs
	}

	r.url.RawQuery = r.queryParams.Encode()
	urlString := r.url.String()

	return r.restyRequest, urlString, nil
}

func (r *Request) addQueryValues(parts url.Values) {
	for key, values := range parts {
		for _, value := range values {
			r.queryParams.Add(key, value)
		}
	}
}
