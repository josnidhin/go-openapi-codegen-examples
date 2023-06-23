/*
Digital Value Services API

Testing BalancesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dvsapi

import (
	"context"
	openapiclient "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_dvsapi_BalancesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BalancesApiService GetBalances", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BalancesApi.GetBalances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
