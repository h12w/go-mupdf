// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

typedef struct custom_device_s
{
	fz_device device;
	fz_device *ori_device;
} custom_device;

extern void (*fill_image_hook)(fz_device *, fz_image *img, fz_matrix *ctm, float alpha);
extern void _fill_image_hook(fz_device *dev, fz_image *image, fz_matrix *ctm, float alpha);
