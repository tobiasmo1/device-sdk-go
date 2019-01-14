// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package mock

import (
	"net/http"

	"github.com/edgexfoundry/edgex-go/pkg/clients/types"
<<<<<<< 26d17496433254eb5129c1da2e49ed4fa355eaa9
	"github.com/edgexfoundry/edgex-go/pkg/models"
=======
	e_models "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/globalsign/mgo/bson"
>>>>>>> - Adds consistent of "e_models" to reference edgex-go/pkg/models; and "logger"
)

type AddressableClientMock struct {
}

func (AddressableClientMock) Addressable(id string) (e_models.Addressable, error) {
	panic("implement me")
}

func (AddressableClientMock) Add(addr *e_models.Addressable) (string, error) {
	return "5b977c62f37ba10e36673802", nil
}

<<<<<<< 26d17496433254eb5129c1da2e49ed4fa355eaa9
func (AddressableClientMock) AddressableForName(name string) (models.Addressable, error) {
	var addressable = models.Addressable{Id: "5b977c62f37ba10e36673802", Name: name}
=======
func (AddressableClientMock) AddressableForName(name string) (e_models.Addressable, error) {
	var addressable = e_models.Addressable{Id: bson.ObjectIdHex("5b977c62f37ba10e36673802").Hex(), Name: name}
>>>>>>> - Adds consistent of "e_models" to reference edgex-go/pkg/models; and "logger"
	var err error = nil
	if name == "" {
		err = types.NewErrServiceClient(http.StatusNotFound, nil)
	}

	return addressable, err
}

func (AddressableClientMock) Update(addr e_models.Addressable) error {
	return nil
}

func (AddressableClientMock) Delete(id string) error {
	return nil
}
