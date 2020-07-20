package ocm

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type OCMAuthentication interface {
	AuthenticatePullSecret(ctx context.Context,pullSecret string) (allowed bool, err error)
}

type authentication service

var _ OCMAuthentication = &authentication{}

func (a authentication) AuthenticatePullSecret(ctx context.Context, pullSecret string) (allowed bool, err error) {
	logrus.Error("FRED 1")
	con := a.client.connection
	logrus.Error("FRED 2")
	selfTokenAutorization := con.AccountsMgmt().V1().AccessToken()
	logrus.Error("FRED 3")

	postResp, err := selfTokenAutorization.Post().
		SendContext(ctx)
	logrus.Error("FRED 4")
	if err != nil {
		logrus.Error("ERR", err)
		return false, err
	}
	body, ok := postResp.GetBody()
	logrus.Error("BODY", body)

	if !ok {
		logrus.Error("FRED 5")
		return false, fmt.Errorf("Empty response from token authentication post request")
	}

	return true, nil
}