package test

import (
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/encoding/json"
	"io/ioutil"
	"net/http"
)

type healthcheck struct {
	Port int `cfg:"port" default:"0"`
}

type healthcheckMockSettings struct {
	*mockSettings
	Healthcheck healthcheck `cfg:"healthcheck"`
}

type localstackHealthcheck struct {
	Services localstackHealthcheckServices
}

type localstackHealthcheckServices struct {
	Cloudwatch string
	Kinesis    string
	SNS        string
	SQS        string
}

func localstackHealthCheck(settings *healthcheckMockSettings, services ...string) func() error {
	return func() error {
		url := fmt.Sprintf("http://%s:%d/health", settings.Host, settings.Healthcheck.Port)
		resp, err := http.Get(url)

		if err != nil {
			return err
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		localstackHealthcheck := &localstackHealthcheck{}
		err = json.Unmarshal(body, localstackHealthcheck)

		if err != nil {
			return err
		}

		for _, service := range services {
			switch service {
			case componentCloudwatch:
				if localstackHealthcheck.Services.Cloudwatch != "running" {
					return fmt.Errorf("service cloudwatch is in state %s", localstackHealthcheck.Services.Cloudwatch)
				}
			case componentKinesis:
				if localstackHealthcheck.Services.Kinesis != "running" {
					return fmt.Errorf("service kinesis is in state %s", localstackHealthcheck.Services.Kinesis)
				}
			case componentSns:
				if localstackHealthcheck.Services.SNS != "running" {
					return fmt.Errorf("service sns is in state %s", localstackHealthcheck.Services.SNS)
				}
			case componentSqs:
				if localstackHealthcheck.Services.SQS != "running" {
					return fmt.Errorf("service sqs is in state %s", localstackHealthcheck.Services.SQS)
				}
			}
		}

		return nil
	}
}

func healthcheckSettings(config cfg.Config, name string) healthcheck {
	healthcheck := healthcheck{}
	key := fmt.Sprintf("mocks.%s.healthcheck", name)
	config.UnmarshalKey(key, &healthcheck)

	return healthcheck
}
