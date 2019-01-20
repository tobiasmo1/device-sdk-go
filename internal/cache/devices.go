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

var (
	dc *deviceCache
)

// DeviceCache holds a set of devices
type DeviceCache interface {
	ForName(name string) (e_models.Device, bool)
	ForId(id string) (e_models.Device, bool)
	All() []e_models.Device
	Add(device e_models.Device) error
	Update(device e_models.Device) error
	UpdateAddressable(addressable e_models.Addressable) error
	Remove(id string) error
	RemoveByName(name string) error
	UpdateAdminState(id string, state e_models.AdminState) error
}

type deviceCache struct {
	dMap    map[string]*e_models.Device // key is Device name
	nameMap map[string]string           // key is id, and value is Device name
}

// ForName returns a Device with the given name.
func (d *deviceCache) ForName(name string) (e_models.Device, bool) {
	if device, ok := d.dMap[name]; ok {
		return *device, ok
	} else {
		return e_models.Device{}, ok
	}
}

// ForId returns a device with the given device id.
func (d *deviceCache) ForId(id string) (e_models.Device, bool) {
	name, ok := d.nameMap[id]
	if !ok {
		return e_models.Device{}, ok
	}
	if device, ok := d.dMap[name]; ok {
		return *device, ok
	}
	return e_models.Device{}, ok
}

// All() returns the current list of devices in the cache.
func (d *deviceCache) All() []e_models.Device {
	devices := make([]e_models.Device, len(d.dMap))
	i := 0
	for _, device := range d.dMap {
		devices[i] = *device
		i++
	}
	return devices
}

// Adds a new device to the cache. This method is used to populate the
// devices cache with pre-existing devices from Core Metadata, as well
// as create new devices returned in a ScanList during discovery.
func (d *deviceCache) Add(device e_models.Device) error {
	if _, ok := d.dMap[device.Name]; ok {
		return fmt.Errorf("device %s has already existed in cache", device.Name)
	}
	d.dMap[device.Name] = &device
	d.nameMap[device.Id] = device.Name
	return nil
}

// Update updates the device in the cache
func (d *deviceCache) Update(device e_models.Device) error {
	if err := d.Remove(device.Id); err != nil {
		return err
	}
	return d.Add(device)
}

// UpdateAddressable updates the device addressable in the cache
func (d *deviceCache) UpdateAddressable(add e_models.Addressable) error {
	found := false
	for _, device := range d.dMap {
		if device.Addressable.Id == add.Id {
			device.Addressable = add
			found = true
		}
	}

	if found == false {
		return fmt.Errorf("addressable %s does not exist in cache", add.Id)
	}

	return nil
}

// Remove removes the specified device by id from the cache.
func (d *deviceCache) Remove(id string) error {
	name, ok := d.nameMap[id]
	if !ok {
		return fmt.Errorf("device %s does not exist in cache", id)
	}

	return d.RemoveByName(name)
}

// RemoveByName removes the specified device by name from the cache.
func (d *deviceCache) RemoveByName(name string) error {
	device, ok := d.dMap[name]
	if !ok {
		return fmt.Errorf("device %s does not exist in cache", name)
	}

	delete(d.nameMap, device.Id)
	delete(d.dMap, name)
	return nil
}

// UpdateAdminState updates the device admin state in cache by id. This method
// is used by the UpdateHandler to trigger update device admin state that's been
// updated directly to Core Metadata.
func (d *deviceCache) UpdateAdminState(id string, state e_models.AdminState) error {
	name, ok := d.nameMap[id]
	if !ok {
		return fmt.Errorf("device %s cannot be found in cache", id)
	}

	d.dMap[name].AdminState = state
	return nil
}

func newDeviceCache(devices []e_models.Device) DeviceCache {
	defaultSize := len(devices) * 2
	dMap := make(map[string]*e_models.Device, defaultSize)
	nameMap := make(map[string]string, defaultSize)
	for i, d := range devices {
		dMap[d.Name] = &devices[i]
		nameMap[d.Id] = d.Name
	}
	dc = &deviceCache{dMap: dMap, nameMap: nameMap}
	return dc
}

func Devices() DeviceCache {
	if dc == nil {
		InitCache()
	}
	return dc
}
