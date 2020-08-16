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

package deezer

import (
	"context"
	"io"

	"github.com/andy-js/beoutil/clients/deezer/models"
	"github.com/andy-js/beoutil/clients/rest"
)

type AlbumIter interface {
	Next(ctx context.Context) ([]models.Album, error)
	Read() int
}

type albumIterImpl struct {
	client   rest.Client
	endpoint string
	read     int
	total    int
	started  bool
}

func (i *albumIterImpl) Next(ctx context.Context) ([]models.Album, error) {
	var resp struct {
		Data  []models.Album `json:"data"`
		Total int            `json:"total"`
		Next  string         `json:"next"`
	}
	if i.started && i.read == i.total {
		return nil, io.EOF
	}
	if err := i.client.DoGet(ctx, i.endpoint, &resp); err != nil {
		return nil, err
	}
	i.read += len(resp.Data)
	if !i.started {
		i.total = resp.Total
		i.started = true
	}
	i.endpoint = resp.Next
	return resp.Data, nil
}

func (i *albumIterImpl) Read() int {
	return i.read
}
