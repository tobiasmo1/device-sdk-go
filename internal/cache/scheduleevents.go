// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"fmt"

	e_models "github.com/edgexfoundry/edgex-go/pkg/models"
)

type ScheduleEventCache interface {
	ForName(name string) (e_models.ScheduleEvent, bool)
	All() []e_models.ScheduleEvent
	Add(schEvt e_models.ScheduleEvent) error
	Update(schEvt e_models.ScheduleEvent) error
	Remove(id string) error
	RemoveByName(name string) error
}

var (
	seCache *scheduleEventCache
)

// ScheduleEventCache is a local cache of ScheduleEvents,
// usually loaded from Core Metadata, and existing scheduleCache
// ScheduleEvents can be used to seed this cache.
type scheduleEventCache struct {
	seMap   map[string]e_models.ScheduleEvent // key is ScheduleEvent name
	nameMap map[string]string                 // key is id, and value is ScheduleEvent name
}

func (s *scheduleEventCache) ForName(name string) (e_models.ScheduleEvent, bool) {
	se, ok := s.seMap[name]
	return se, ok
}

func (s *scheduleEventCache) All() []e_models.ScheduleEvent {
	ses := make([]e_models.ScheduleEvent, len(s.seMap))
	i := 0
	for _, schEvt := range s.seMap {
		ses[i] = schEvt
		i++
	}
	return ses
}

func (s *scheduleEventCache) Add(schEvt e_models.ScheduleEvent) error {
	if _, ok := s.seMap[schEvt.Name]; ok {
		return fmt.Errorf("schedule event %s has already existed in cache", schEvt.Name)
	} else {
		return fmt.Errorf("TJM: schedule event added to cache [%s]", schEvt.Name)
	}
	s.seMap[schEvt.Name] = schEvt
	s.nameMap[schEvt.Id] = schEvt.Name
	return nil
}

func (s *scheduleEventCache) Update(schEvt e_models.ScheduleEvent) error {
	if err := s.Remove(schEvt.Id); err != nil {
		return err
	}
	return s.Add(schEvt)
}

func (s *scheduleEventCache) Remove(id string) error {
	name, ok := s.nameMap[id]
	if !ok {
		return fmt.Errorf("schedule event %s does not exist in cache", id)
	}

	return s.RemoveByName(name)
}

// RemoveByName removes the specified device by name from the cache.
func (s *scheduleEventCache) RemoveByName(name string) error {
	schEvt, ok := s.seMap[name]
	if !ok {
		return fmt.Errorf("schedule event %s does not exist in cache", name)
	}

	delete(s.nameMap, schEvt.Id)
	delete(s.seMap, name)
	return nil
}

// Creates a singleton ScheduleEvent Cache instance.
func newScheduleEventCache(schEvts []e_models.ScheduleEvent) ScheduleEventCache {
	defaultSize := len(schEvts) * 2
	seMap := make(map[string]e_models.ScheduleEvent, defaultSize)
	nameMap := make(map[string]string, defaultSize)
	for _, se := range schEvts {
		seMap[se.Name] = se
		nameMap[se.Id] = se.Name
	}
	seCache = &scheduleEventCache{seMap: seMap, nameMap: nameMap}
	return seCache
}

func ScheduleEvents() ScheduleEventCache {
	if seCache == nil {
		InitCache()
	}
	return seCache
}
