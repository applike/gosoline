package http

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type restyClient interface {
	AddRetryCondition(conditionFunc resty.RetryConditionFunc) *resty.Client
	NewRequest() *resty.Request
	SetCookie(cookie *http.Cookie) *resty.Client
	SetCookies(cookies []*http.Cookie) *resty.Client
	SetProxy(proxy string) *resty.Client
	SetRedirectPolicy(policies ...interface{}) *resty.Client
	SetTimeout(timeout time.Duration) *resty.Client
}
