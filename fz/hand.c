// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "_cgo_export.h"

void (*fill_image_hook)(fz_device *, fz_image *img, fz_matrix *ctm, float alpha) = _fill_image_hook;

void _fill_image_hook(fz_device *dev, fz_image *image, fz_matrix *ctm, float alpha) {
    custom_device *cd = (custom_device*)dev;
    fillImageHookCallback_Go(dev, image, ctm, alpha);
    cd->ori_device->fill_image(dev, image, ctm, alpha);
}
