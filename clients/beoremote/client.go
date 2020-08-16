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

type Client struct {
	client      rest.Client
	baseURL     string
	BeoZone     BeoZone
	BeoDevice   BeoDevice
	BeoSecurity *BeoSecurity
	BeoHome     BeoHome
}

func NewClient(addr string) *Client {
	c := rest.NewJSONClient()
	baseURL := "http://" + addr + ":8080"
	return &Client{
		client:  c,
		baseURL: baseURL,
		BeoDevice: &beoDevice{
			client:  c,
			baseURL: baseURL,
		},
		BeoZone: &beoZone{
			client:  c,
			baseURL: baseURL,
		},
		BeoSecurity: &BeoSecurity{client: c},
		BeoHome: &beoHome{
			client:  c,
			baseURL: baseURL,
		},
	}
}

func (l *Client) GetBeoDevice(ctx context.Context) (*models.BeoDeviceInfo, error) {
	var r models.BeoDeviceResponse
	err := l.client.DoGet(ctx, "/BeoDevice", &r)
	if err != nil {
		return nil, err
	}
	return &r.BeoDevice, nil
}
