// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fz

/*
#include <mupdf/fitz.h>
#include "hand.h"
*/
import "C"

import (
	"unsafe"
)

type DeviceHook struct {
	Device
	OriDevice *Device
	ImageHook func(*Image, *Matrix, float32)
}

func NewDeviceHook(device *Device) *DeviceHook {
	dev := &DeviceHook{Device: *device, OriDevice: device}
	dev.FillImage_ = C.fill_image_hook
	return dev
}

func (dev *DeviceHook) ToDevice() *Device {
	return (*Device)(unsafe.Pointer(dev))
}

func (dev *DeviceHook) Free() {
	dev.OriDevice.Free()
}

//export fillImageHookCallback_Go
func fillImageHookCallback_Go(dev_ *C.fz_device, image_ *C.fz_image, ctm_ *C.fz_matrix, alpha_ C.float) {
	dev := (*DeviceHook)(unsafe.Pointer(dev_))
	image := (*Image)(unsafe.Pointer(image_))
	ctm := (*Matrix)(unsafe.Pointer(ctm_))
	alpha := float32(alpha_)
	if dev.ImageHook != nil {
		dev.ImageHook(image, ctm, alpha)
	}
}
