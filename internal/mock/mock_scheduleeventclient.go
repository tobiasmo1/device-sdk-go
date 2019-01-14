// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package mock

import (
	"errors"

	e_models "github.com/edgexfoundry/edgex-go/pkg/models"
)

type ScheduleEventClientMock struct{}

func (ScheduleEventClientMock) Add(dev *e_models.ScheduleEvent) (string, error) {
	return "", nil
}

func (ScheduleEventClientMock) Delete(id string) error {
	panic("implement me")
}

func (ScheduleEventClientMock) DeleteByName(name string) error {
	panic("implement me")
}

func (ScheduleEventClientMock) ScheduleEvent(id string) (e_models.ScheduleEvent, error) {
	panic("implement me")
}

func (ScheduleEventClientMock) ScheduleEventForName(name string) (e_models.ScheduleEvent, error) {
	var scheduleEvent = e_models.ScheduleEvent{Name: name}
	var err error = nil
	if name == "" {
		err = errors.New("scheduleEvent not exist")
	}
	return scheduleEvent, err
}

func (ScheduleEventClientMock) ScheduleEvents() ([]e_models.ScheduleEvent, error) {
	panic("implement me")
}

func (ScheduleEventClientMock) ScheduleEventsForAddressable(name string) ([]e_models.ScheduleEvent, error) {
	panic("implement me")
}

func (ScheduleEventClientMock) ScheduleEventsForAddressableByName(name string) ([]e_models.ScheduleEvent, error) {
	panic("implement me")
}

func (ScheduleEventClientMock) ScheduleEventsForServiceByName(name string) ([]e_models.ScheduleEvent, error) {
	return []e_models.ScheduleEvent{}, nil
}

func (ScheduleEventClientMock) Update(dev e_models.ScheduleEvent) error {
	panic("implement me")
}
