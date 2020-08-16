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

type Track struct {
	ID                    int           `json:"id"`
	Readable              bool          `json:"readable"`
	Title                 string        `json:"title"`
	TitleShort            string        `json:"title_short"`
	TitleVersion          string        `json:"title_version"`
	Unseen                bool          `json:"unseen"`
	ISrc                  string        `json:"isrc"`
	Link                  string        `json:"link"`
	Share                 string        `json:"share"`
	Duration              int           `json:"duration"`
	TrackPosition         int           `json:"track_position"`
	DiskNumber            int           `json:"disk_number"`
	Rank                  int           `json:"rank"`
	ReleaseDate           string        `json:"release_date"`
	ExplicitLyrics        bool          `json:"explicit_lyrics"`
	ExplicitContentLyrics int           `json:"explicit_content_lyrics"`
	ExplicitContentCover  int           `json:"explicit_content_cover"`
	Preview               string        `json:"preview"`
	BPM                   float64       `json:"bpm"`
	Gain                  float64       `json:"gain"`
	AvailableCountries    []string      `json:"available_countries"`
	Alternative           *Track        `json:"alternative"`
	Contributors          []Contributor `json:"contributors"`
	MD5Image              string        `json:"md5_image"`
	Artist                *Artist       `json:"artist"`
	Album                 *Album        `json:"album"`
}
