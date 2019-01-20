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

type ScheduleClientMock struct {
}

func (s *ScheduleClientMock) Add(dev *e_models.Schedule) (string, error) {
	return "", nil
}

func (s *ScheduleClientMock) Delete(id string) error {
	return nil
}

func (s *ScheduleClientMock) DeleteByName(name string) error {
	return nil
}

func (s *ScheduleClientMock) Schedule(id string) (e_models.Schedule, error) {
	return e_models.Schedule{}, nil
}

func (s *ScheduleClientMock) Schedules() ([]e_models.Schedule, error) {
	return []e_models.Schedule{}, nil
}

func (s *ScheduleClientMock) Update(dev e_models.Schedule) error {
	return nil
}

func (s *ScheduleClientMock) ScheduleForName(name string) (e_models.Schedule, error) {
	var schedule = e_models.Schedule{Name: name}
	var err error = nil
	if name == "" {
		err = errors.New("schedule not exist")
	}
	return schedule, err
}
