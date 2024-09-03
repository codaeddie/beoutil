// Copyright (c) 2020-2024 Andrew Stormont
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package beoremote

import (
	"context"

	"beoutil/clients/beoremote/models"
	"beoutil/clients/rest"
)

type BeoDevice interface {
	GetState(ctx context.Context) (models.PowerState, error)
	Standby(ctx context.Context) error
	AllStandby(ctx context.Context) error
	PowerOn(ctx context.Context) error
	Reboot(ctx context.Context) error
}

type beoDevice struct {
	client  rest.Client
	baseURL string
}

func (d *beoDevice) GetState(ctx context.Context) (models.PowerState, error) {
	var r models.StandbyRequest
	err := d.client.DoGet(ctx, d.baseURL+"/BeoDevice/powerManagement/standby", &r)
	if err != nil {
		return "", err
	}
	return r.Standby.PowerState, nil
}

func (d *beoDevice) setPowerState(ctx context.Context, state models.PowerState) error {
	r := models.StandbyRequest{
		Standby: models.Standby{
			PowerState: state,
		},
	}
	_, err := d.client.DoPut(ctx, d.baseURL+"/BeoDevice/powerManagement/standby", r)
	return err
}

func (d *beoDevice) Standby(ctx context.Context) error {
	return d.setPowerState(ctx, models.PowerStateStandby)
}

func (d *beoDevice) AllStandby(ctx context.Context) error {
	return d.setPowerState(ctx, models.PowerStateAllStandby)
}

func (d *beoDevice) PowerOn(ctx context.Context) error {
	return d.setPowerState(ctx, models.PowerStateOn)
}

func (d *beoDevice) Reboot(ctx context.Context) error {
	return d.setPowerState(ctx, models.PowerStateReboot)
}
