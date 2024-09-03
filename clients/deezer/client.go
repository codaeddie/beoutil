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
	"net/url"
	"strconv"

	"beoutil/clients/deezer/models"
	"beoutil/clients/rest"
)

type Client struct {
	client  rest.Client
	baseURL string
}

func NewClient() *Client {
	return &Client{
		client:  rest.NewJSONClient(),
		baseURL: "https://api.deezer.com",
	}
}

type SearchOptions struct {
	Q     string
	Index int
	Limit int
}

func (c *Client) SearchArtist(ctx context.Context, opts *SearchOptions) ([]models.Artist, error) {
	var resp struct {
		Data []models.Artist `json:"data"`
	}
	reqURL := "/search/artist?q=" + url.QueryEscape(opts.Q)
	if opts.Index != 0 {
		reqURL += "&index=" + strconv.Itoa(opts.Index)
	}
	if opts.Limit != 0 {
		reqURL += "&limit=" + strconv.Itoa(opts.Limit)
	}
	if err := c.client.DoGet(ctx, c.baseURL+reqURL, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) NewAlbumIter(artistID string) AlbumIter {
	return &albumIterImpl{
		client:   c.client,
		endpoint: c.baseURL + "/artist/" + artistID + "/albums",
	}
}

func (c *Client) GetAlbumTracks(ctx context.Context, albumID string) ([]models.Track, error) {
	var resp struct {
		Data []models.Track `json:"data"`
	}
	if err := c.client.DoGet(ctx, c.baseURL+"/album/"+albumID+"/tracks", &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) GetTrack(ctx context.Context, trackID string) (models.Track, error) {
	var resp models.Track
	if err := c.client.DoGet(ctx, c.baseURL+"/track/"+trackID, &resp); err != nil {
		return models.Track{}, nil
	}
	return resp, nil
}
