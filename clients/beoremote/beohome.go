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

	"github.com/andy-js/beoutil/clients/beoremote/models"
	"github.com/andy-js/beoutil/clients/rest"
)

type BeoHome interface {
	GetTimers(ctx context.Context) ([]models.Timer, error)
	AddTimer(ctx context.Context, timer models.Timer) error
	ModifyTimer(ctx context.Context, timer models.Timer) error
	DeleteTimer(ctx context.Context, id string) error
}

type beoHome struct {
	baseURL string
	client  rest.Client
}

func (b *beoHome) GetTimers(ctx context.Context) ([]models.Timer, error) {
	var r models.TimerListResponse
	if err := b.client.DoGet(ctx, b.baseURL+"/BeoHome/trigger/timerList", &r); err != nil {
		return nil, err
	}
	return r.TimerList.Timer, nil
}

type TimerRequest struct {
	Timer models.Timer `json:"timer"`
}

func (b *beoHome) AddTimer(ctx context.Context, timer models.Timer) error {
	r := TimerRequest{Timer: timer}
	_, err := b.client.DoPost(ctx, b.baseURL+"/BeoHome/trigger/timerList", r)
	return err
}

func (b *beoHome) ModifyTimer(ctx context.Context, timer models.Timer) error {
	r := TimerRequest{Timer: timer}
	_, err := b.client.DoPut(ctx, b.baseURL+"/BeoHome/trigger/timerList/"+timer.Id, r)
	return err
}

func (b *beoHome) DeleteTimer(ctx context.Context, id string) error {
	_, err := b.client.DoDelete(ctx, b.baseURL+"/BeoHome/trigger/timerList/"+id)
	return err
}
