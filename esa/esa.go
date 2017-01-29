// -*- coding: utf-8 -*-
//---------------------------------#
// File Name     : esa/esa.go
// Author        : todoroki
// Date Created  : 2017-01-29
//---------------------------------#

package esaClient

import (
	"os"

	"../env"

	"github.com/yuichiro-h/go-esa"
)

func NewEsaClient() *esa.Client {
	accessToken := os.Getenv(env.EsaAccessToken)
	return esa.New(&esa.Config{AccessToken: accessToken})
}
