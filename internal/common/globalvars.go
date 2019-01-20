// -*- mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	ds_models "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/edgex-go/pkg/clients/coredata"
	logger "github.com/edgexfoundry/edgex-go/pkg/clients/logging"
	"github.com/edgexfoundry/edgex-go/pkg/clients/metadata"
	e_models "github.com/edgexfoundry/edgex-go/pkg/models"
)

var (
	ServiceName           string
	ServiceVersion        string
	CurrentConfig         *Config
	CurrentDeviceService  e_models.DeviceService
	UseRegistry           bool
	ServiceLocked         bool
	Driver                ds_models.ProtocolDriver
	EventClient           coredata.EventClient
	AddressableClient     metadata.AddressableClient
	DeviceClient          metadata.DeviceClient
	DeviceServiceClient   metadata.DeviceServiceClient
	DeviceProfileClient   metadata.DeviceProfileClient
	LoggingClient         logger.LoggingClient
	ValueDescriptorClient coredata.ValueDescriptorClient
	ScheduleClient        metadata.ScheduleClient
	ScheduleEventClient   metadata.ScheduleEventClient
)
