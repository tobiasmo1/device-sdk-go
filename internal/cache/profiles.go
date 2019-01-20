// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"fmt"
	"strings"

	"github.com/edgexfoundry/edgex-go/pkg/models"
	e_models "github.com/edgexfoundry/edgex-go/pkg/models"
)

const (
	getOpsStr string = "get"
	setOpsStr string = "set"
)

var (
	pc *profileCache
)

type ProfileCache interface {
	ForName(name string) (e_models.DeviceProfile, bool)
	ForId(id string) (e_models.DeviceProfile, bool)
	All() []e_models.DeviceProfile
	Add(profile e_models.DeviceProfile) error
	Update(profile e_models.DeviceProfile) error
	Remove(id string) error
	RemoveByName(name string) error
	DeviceObject(profileName string, objectName string) (e_models.DeviceObject, bool)
	CommandExists(profileName string, cmd string) (bool, error)
	ResourceOperations(profileName string, cmd string, method string) ([]e_models.ResourceOperation, error)
	ResourceOperation(profileName string, object string, method string) (e_models.ResourceOperation, error)
}

type profileCache struct {
	dpMap    map[string]e_models.DeviceProfile // key is DeviceProfile name
	nameMap  map[string]string                 // key is id, and value is DeviceProfile name
	doMap    map[string]map[string]e_models.DeviceObject
	getOpMap map[string]map[string][]e_models.ResourceOperation
	setOpMap map[string]map[string][]e_models.ResourceOperation
	cmdMap   map[string]map[string]e_models.Command
}

func (p *profileCache) ForName(name string) (e_models.DeviceProfile, bool) {
	dp, ok := p.dpMap[name]
	return dp, ok
}

func (p *profileCache) ForId(id string) (e_models.DeviceProfile, bool) {
	name, ok := p.nameMap[id]
	if !ok {
		return e_models.DeviceProfile{}, ok
	}

	dp, ok := p.dpMap[name]
	return dp, ok
}

func (p *profileCache) All() []e_models.DeviceProfile {
	ps := make([]e_models.DeviceProfile, len(p.dpMap))
	i := 0
	for _, profile := range p.dpMap {
		ps[i] = profile
		i++
	}
	return ps
}

func (p *profileCache) Add(profile e_models.DeviceProfile) error {
	if _, ok := p.dpMap[profile.Name]; ok {
		return fmt.Errorf("device profile %s has already existed in cache", profile.Name)
	}
	p.dpMap[profile.Name] = profile
	p.nameMap[profile.Id] = profile.Name
	p.doMap[profile.Name] = deviceObjectSliceToMap(profile.DeviceResources)
	p.getOpMap[profile.Name], p.setOpMap[profile.Name] = profileResourceSliceToMaps(profile.Resources)
	p.cmdMap[profile.Name] = commandSliceToMap(profile.Commands)
	return nil
}

func deviceObjectSliceToMap(deviceObjects []e_models.DeviceObject) map[string]e_models.DeviceObject {
	result := make(map[string]e_models.DeviceObject, len(deviceObjects))
	for _, do := range deviceObjects {
		result[do.Name] = do
	}
	return result
}

func profileResourceSliceToMaps(profileResources []e_models.ProfileResource) (map[string][]e_models.ResourceOperation, map[string][]models.ResourceOperation) {
	getResult := make(map[string][]e_models.ResourceOperation, len(profileResources))
	setResult := make(map[string][]e_models.ResourceOperation, len(profileResources))
	for _, pr := range profileResources {
		if len(pr.Get) > 0 {
			getResult[pr.Name] = pr.Get
		}
		if len(pr.Set) > 0 {
			setResult[pr.Name] = pr.Set
		}
	}
	return getResult, setResult
}

func commandSliceToMap(commands []e_models.Command) map[string]e_models.Command {
	result := make(map[string]e_models.Command, len(commands))
	for _, cmd := range commands {
		result[cmd.Name] = cmd
	}
	return result
}

func (p *profileCache) Update(profile e_models.DeviceProfile) error {
	if err := p.Remove(profile.Id); err != nil {
		return err
	}
	return p.Add(profile)
}

func (p *profileCache) Remove(id string) error {
	name, ok := p.nameMap[id]
	if !ok {
		return fmt.Errorf("device profile %s does not exist in cache", id)
	}

	return p.RemoveByName(name)
}

func (p *profileCache) RemoveByName(name string) error {
	profile, ok := p.dpMap[name]
	if !ok {
		return fmt.Errorf("device profile %s does not exist in cache", name)
	}

	delete(p.dpMap, name)
	delete(p.nameMap, profile.Id)
	delete(p.doMap, name)
	delete(p.getOpMap, name)
	delete(p.setOpMap, name)
	delete(p.cmdMap, name)
	return nil
}

func (p *profileCache) DeviceObject(profileName string, objectName string) (e_models.DeviceObject, bool) {
	objs, ok := p.doMap[profileName]
	if !ok {
		return e_models.DeviceObject{}, ok
	}

	obj, ok := objs[objectName]
	return obj, ok
}

// CommandExists returns a bool indicating whether the specified command exists for the
// specified (by name) device. If the specified device doesn't exist, an error is returned.
func (p *profileCache) CommandExists(profileName string, cmd string) (bool, error) {
	commands, ok := p.cmdMap[profileName]
	if !ok {
		err := fmt.Errorf("profiles: CommandExists: specified profile: %s not found", profileName)
		return false, err
	}

	if _, ok := commands[cmd]; !ok {
		return false, nil
	}

	return true, nil
}

// Get ResourceOperations
func (p *profileCache) ResourceOperations(profileName string, cmd string, method string) ([]e_models.ResourceOperation, error) {
	var resOps []e_models.ResourceOperation
	var rosMap map[string][]e_models.ResourceOperation
	var ok bool
	if strings.ToLower(method) == getOpsStr {
		if rosMap, ok = p.getOpMap[profileName]; !ok {
			return nil, fmt.Errorf("profiles: ResourceOperations: specified profile: %s not found", profileName)
		}
	} else if strings.ToLower(method) == setOpsStr {
		if rosMap, ok = p.setOpMap[profileName]; !ok {
			return nil, fmt.Errorf("profiles: ResourceOperations: specified profile: %s not found", profileName)
		}
	}

	if resOps, ok = rosMap[cmd]; !ok {
		return nil, fmt.Errorf("profiles: ResourceOperations: specified cmd: %s not found", cmd)
	}
	return resOps, nil
}

// Return the first matched ResourceOperation
func (p *profileCache) ResourceOperation(profileName string, object string, method string) (e_models.ResourceOperation, error) {
	var ro e_models.ResourceOperation
	var rosMap map[string][]e_models.ResourceOperation
	var ok bool
	if strings.ToLower(method) == getOpsStr {
		if rosMap, ok = p.getOpMap[profileName]; !ok {
			return ro, fmt.Errorf("profiles: ResourceOperation: specified profile: %s not found", profileName)
		}
	} else if strings.ToLower(method) == setOpsStr {
		if rosMap, ok = p.setOpMap[profileName]; !ok {
			return ro, fmt.Errorf("profiles: ResourceOperations: specified profile: %s not found", profileName)
		}
	}

	if ro, ok = retrieveFirstRObyObject(rosMap, object); !ok {
		return ro, fmt.Errorf("profiles: specified ResourceOperation by object %s not found", object)
	}
	return ro, nil
}

func retrieveFirstRObyObject(rosMap map[string][]e_models.ResourceOperation, object string) (e_models.ResourceOperation, bool) {
	for _, ros := range rosMap {
		for _, ro := range ros {
			if ro.Object == object {
				return ro, true
			}
		}
	}
	return e_models.ResourceOperation{}, false
}

func newProfileCache(profiles []e_models.DeviceProfile) ProfileCache {
	defaultSize := len(profiles) * 2
	dpMap := make(map[string]e_models.DeviceProfile, defaultSize)
	nameMap := make(map[string]string, defaultSize)
	doMap := make(map[string]map[string]e_models.DeviceObject, defaultSize)
	getOpMap := make(map[string]map[string][]e_models.ResourceOperation, defaultSize)
	setOpMap := make(map[string]map[string][]e_models.ResourceOperation, defaultSize)
	cmdMap := make(map[string]map[string]e_models.Command, defaultSize)
	for _, dp := range profiles {
		dpMap[dp.Name] = dp
		nameMap[dp.Id] = dp.Name
		doMap[dp.Name] = deviceObjectSliceToMap(dp.DeviceResources)
		getOpMap[dp.Name], setOpMap[dp.Name] = profileResourceSliceToMaps(dp.Resources)
		cmdMap[dp.Name] = commandSliceToMap(dp.Commands)
	}
	pc = &profileCache{dpMap: dpMap, nameMap: nameMap, doMap: doMap, getOpMap: getOpMap, setOpMap: setOpMap, cmdMap: cmdMap}
	return pc
}

func Profiles() ProfileCache {
	if pc == nil {
		InitCache()
	}
	return pc
}
