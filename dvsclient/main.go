/**
 * @author Jose Nidhin
 */
package main

import (
	"context"
	"log"
	"os"
	"time"

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

	logger.Printf("Servers: %+v\n\n", cfg.Servers)

	apiClient := dvsapi.NewAPIClient(cfg)
	balance, res, err := apiClient.BalancesApi.GetBalances(ctx).UnitType(dvsapi.UNITTYPES_CURRENCY).Unit("USD").Execute()
	logger.Printf("balance:%+v\n res: %+v\n err: %+v\n\n", balance, res, err)

	postLookupMobileNumberRequest := dvsapi.NewPostLookupMobileNumberRequest("+6598765432")
	lookupMobileNumber, res, err := apiClient.MobileNumberApi.PostLookupMobileNumber(ctx).PostLookupMobileNumberRequest(*postLookupMobileNumberRequest).Execute()
	logger.Printf("lookupMobileNumber:%+v\n res: %+v\n err: %+v\n\n", lookupMobileNumber, res, err)

	postLookupMobileNumberRequest = dvsapi.NewPostLookupMobileNumberRequest("+659876543212312")
	lookupMobileNumber, res, err = apiClient.MobileNumberApi.PostLookupMobileNumber(ctx).PostLookupMobileNumberRequest(*postLookupMobileNumberRequest).Execute()
	logger.Printf("lookupMobileNumber:%+v\n res: %+v\n err: %+v\n\n", lookupMobileNumber, res, err)
	getApiError(err)

	postLookupMobileNumberRequest = dvsapi.NewPostLookupMobileNumberRequest("+6598765432")
	sCtx, sCtxCancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer sCtxCancel()
	lookupMobileNumber, res, err = apiClient.MobileNumberApi.PostLookupMobileNumber(sCtx).PostLookupMobileNumberRequest(*postLookupMobileNumberRequest).Execute()
	logger.Printf("lookupMobileNumber:%+v\n res: %+v\n err: %+v\n\n", lookupMobileNumber, res, err)
}

func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logger.Panicf("Expected env var %s to be set", key)
	}
	return val
}

func getApiError(err error) {
	apiErr, ok := err.(*dvsapi.GenericOpenAPIError)
	if !ok {
		return
	}

	m := apiErr.Model()
	errs, ok := m.(dvsapi.Errors)
	if !ok {
		return
	}
	logger.Printf("errors: %+v\n code: %d\n msg: %s\n\n", errs, errs.Errors[0].Code, errs.Errors[0].Message)
}
