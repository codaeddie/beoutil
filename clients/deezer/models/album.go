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

package models

type TrackData struct {
	Data []Track `json:"data"`
}

type Album struct {
	ID                    int           `json:"id"`
	Title                 string        `json:"title"`
	UPC                   string        `json:"upc"`
	Link                  string        `json:"link"`
	Share                 string        `json:"share"`
	Cover                 string        `json:"cover"`
	CoverSmall            string        `json:"cover_small"`
	CoverMedium           string        `json:"cover_medium"`
	CoverLarge            string        `json:"cover_large"`
	CoverXL               string        `json:"cover_xl"`
	MD5Image              string        `json:"md5_image"`
	GenreID               int           `json:"genre_id"`
	Genres                []Genre       `json:"genre"`
	Label                 string        `json:"label"`
	NbTracks              int           `json:"nb_tracks"`
	Duration              int           `json:"duration"`
	Fans                  int           `json:"fans"`
	ReleaseDate           string        `json:"release_date"`
	RecordType            string        `json:"record_type"`
	Available             bool          `json:"available"`
	Alternative           *Album        `json:"alternative"`
	TrackList             string        `json:"tracklist"`
	ExplicitLyrics        bool          `json:"explicit_lyrics"`
	ExplicitContentLyrics int           `json:"explicit_content_lyrics"`
	ExplicitContentCover  int           `json:"explicit_content_cover"`
	Contributors          []Contributor `json:"contributors"`
	Fallback              *Album        `json:"fallback"`
	Artist                *Artist       `json:"artist"`
	Tracks                TrackData     `json:"tracks"`
}
