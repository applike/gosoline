package suite

import (
	"context"
	"fmt"
	"github.com/applike/gosoline/pkg/apiserver"
	"github.com/applike/gosoline/pkg/application"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/kernel"
	"github.com/applike/gosoline/pkg/mon"
	"github.com/applike/gosoline/pkg/test/env"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func init() {
	testCaseDefinitions["apiServer"] = testCaseDefinition{
		matcher: isTestCaseApiServer,
		builder: buildTestCaseApiServer,
	}
}

type TestingSuiteApiDefinitionsAware interface {
	SetupApiDefinitions() apiserver.Definer
}

type ApiServerTestCase struct {
	Method             string
	Url                string
	Headers            map[string]string
	Body               interface{}
	ExpectedStatusCode int
	ExpectedResult     interface{}
	ExpectedErr        error
	Assert             func() error
}

func (c ApiServerTestCase) request(client *resty.Client) (*resty.Response, error) {
	req := client.R()

	if c.Headers != nil {
		req.SetHeaders(c.Headers)
	}

	if c.Body != nil {
		req.SetBody(c.Body)
	}

	if c.ExpectedResult != nil {
		req.SetResult(c.ExpectedResult)
	}

	return req.Execute(c.Method, c.Url)
}

func isTestCaseApiServer(method reflect.Method) bool {
	if method.Func.Type().NumIn() != 1 {
		return false
	}

	if method.Func.Type().NumOut() != 1 {
		return false
	}

	actualType0 := method.Func.Type().Out(0)
	expectedType := reflect.TypeOf((*ApiServerTestCase)(nil))

	return actualType0 == expectedType
}

func buildTestCaseApiServer(suite TestingSuite, method reflect.Method) (testCaseRunner, error) {
	var ok bool
	var apiDefinitionAware TestingSuiteApiDefinitionsAware
	var server *apiserver.ApiServer

	out := method.Func.Call([]reflect.Value{reflect.ValueOf(suite)})
	tc := out[0].Interface().(*ApiServerTestCase)

	if apiDefinitionAware, ok = suite.(TestingSuiteApiDefinitionsAware); !ok {
		return nil, fmt.Errorf("the suite has to implement the TestingSuiteApiDefinitionsAware interface to be able to run apiserver test cases")
	}

	apiDefinitions := apiDefinitionAware.SetupApiDefinitions()

	return func(t *testing.T, suite TestingSuite, suiteOptions *suiteOptions, environment *env.Environment) {
		suite.SetT(t)

		suiteOptions.appModules["api"] = func(ctx context.Context, config cfg.Config, logger mon.Logger) (kernel.Module, error) {
			module, err := apiserver.New(apiDefinitions)(ctx, config, logger)

			if err != nil {
				return nil, err
			}

			server = module.(*apiserver.ApiServer)

			return server, err
		}

		suiteOptions.addAppOption(application.WithConfigMap(map[string]interface{}{
			"api_port": 0,
		}))

		runTestCaseApplication(t, suite, suiteOptions, environment, func(app *appUnderTest) {
			port, err := server.GetPort()

			if err != nil {
				assert.FailNow(t, err.Error(), "can not get port of server")
				return
			}

			url := fmt.Sprintf("http://127.0.0.1:%d", *port)
			client := resty.New().SetHostURL(url)
			response, err := tc.request(client)

			assert.Equal(t, tc.ExpectedStatusCode, response.StatusCode(), "response status code should match")

			if tc.ExpectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.ExpectedErr.Error())
			}

			app.Stop()
			app.WaitDone()

			if tc.Assert != nil {
				if err := tc.Assert(); err != nil {
					assert.FailNowf(t, err.Error(), "there should be no error on assert")
				}
			}
		})
	}, nil
}
