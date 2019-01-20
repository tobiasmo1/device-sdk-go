// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package mock

import (
	"errors"
	"fmt"

	e_models "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/globalsign/mgo/bson"
)

const (
	invalidDeviceId = "5b9a4f9a64562a2f966fdb0b"
)

type DeviceClientMock struct{}

func (dc *DeviceClientMock) Add(dev *e_models.Device) (string, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) Delete(id string) error {
	panic("implement me")
}

func (dc *DeviceClientMock) DeleteByName(name string) error {
	panic("implement me")
}

func (dc *DeviceClientMock) CheckForDevice(token string) (e_models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) Device(id string) (e_models.Device, error) {
	if id == invalidDeviceId {
		return e_models.Device{}, fmt.Errorf("invalid id")
	}
	return e_models.Device{}, nil
}

func (dc *DeviceClientMock) DeviceForName(name string) (e_models.Device, error) {
	var device = e_models.Device{Id: bson.ObjectIdHex("5b977c62f37ba10e36673802").Hex(), Name: name}
	var err error = nil
	if name == "" {
		err = errors.New("Item not found")
	}

	return device, err
}

func (dc *DeviceClientMock) Devices() ([]e_models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesByLabel(label string) ([]e_models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForAddressable(addressableid string) ([]e_models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForAddressableByName(addressableName string) ([]e_models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForProfile(profileid string) ([]e_models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForProfileByName(profileName string) ([]e_models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForService(serviceid string) ([]e_models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForServiceByName(serviceName string) ([]e_models.Device, error) {
	return []e_models.Device{}, nil
}

func (dc *DeviceClientMock) Update(dev e_models.Device) error {
	return nil
}

func (dc *DeviceClientMock) UpdateAdminState(id string, adminState string) error {
	return nil
}

func (dc *DeviceClientMock) UpdateAdminStateByName(name string, adminState string) error {
	return nil
}

func (dc *DeviceClientMock) UpdateLastConnected(id string, time int64) error {
	return nil
}

func (dc *DeviceClientMock) UpdateLastConnectedByName(name string, time int64) error {
	return nil
}

func (dc *DeviceClientMock) UpdateLastReported(id string, time int64) error {
	return nil
}

func (dc *DeviceClientMock) UpdateLastReportedByName(name string, time int64) error {
	return nil
}

func (dc *DeviceClientMock) UpdateOpState(id string, opState string) error {
	return nil
}

func (dc *DeviceClientMock) UpdateOpStateByName(name string, opState string) error {
	return nil
}
