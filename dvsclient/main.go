/**
 * @author Jose Nidhin
 */
package main

import (
	"context"
	"log"
	"os"

	"github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
)

const (
	EnvApiKey    = "DVS_API_KEY"
	EnvApiSecret = "DVS_API_KEY_SECRET"
)

var logger = log.Default()

func main() {
	username := mustGetEnv(EnvApiKey)
	password := mustGetEnv(EnvApiSecret)
	basicAuth := dvsapi.BasicAuth{
		UserName: username,
		Password: password,
	}

	ctx := context.WithValue(context.Background(), dvsapi.ContextBasicAuth, basicAuth)
	cfg := dvsapi.NewConfiguration()

	logger.Printf("%+v\n", cfg.Servers)

	apiClient := dvsapi.NewAPIClient(cfg)
	balance, res, err := apiClient.BalancesApi.GetBalances(ctx).UnitType(dvsapi.UNITTYPES_CURRENCY).Unit("USD").Execute()
	if err != nil {
		logger.Printf("%+v\n", err)
	}

	logger.Printf("Balances\n%+v\n, %+v\n\n", balance, res)

	postLookupMobileNumberRequest := dvsapi.NewPostLookupMobileNumberRequest("+6598765432")
	lookupMobileNumber, res, err := apiClient.MobileNumberApi.PostLookupMobileNumber(ctx).PostLookupMobileNumberRequest(*postLookupMobileNumberRequest).Execute()
	if err != nil {
		logger.Printf("%+v\n", err)
	}

	logger.Printf("Mobile Number Lookup\n%+v\n, %+v\n\n", lookupMobileNumber, res)
}

func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logger.Panicf("Expected env var %s to be set", key)
	}
	return val
}
