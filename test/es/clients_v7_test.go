// +build integration

package es_test

import (
	"github.com/applike/gosoline/pkg/es"
	"github.com/applike/gosoline/pkg/mdl"
	"github.com/applike/gosoline/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Skip("Skipping for now due to issues with multiple docker based tests")

	configFilePath := "config-v7.test.yml"

	defer test.Shutdown()

	test.Boot(mdl.String(configFilePath))

	config, logger := getMocks(configFilePath)

	clientV7 := es.NewClient(config, logger, "test_v7")

	res, err := clientV7.Info()

	assert.NoError(t, err, "can't get Info from ElasticSearch")
	assert.NotEqual(t, res.IsError(), nil, "response with error")
}

func TestGetAwsClient(t *testing.T) {
	config, logger := getMocks("config-v7.test.yml")

	endpointKey := "es_test_v7_aws_endpoint"
	if !config.IsSet(endpointKey) {
		t.Skipf("%s missed in config", endpointKey)
		return
	}

	clientAwsV7, err := es.GetAwsClientV6(logger, config.GetString(endpointKey))
	assert.NoError(t, err, "can't get Info from ElasticSearch")

	res, err := clientAwsV7.Info()
	assert.NoError(t, err, "can't get Info from ElasticSearch at AWS")
	assert.NotEqual(t, res.IsError(), nil, "response with error")
}
