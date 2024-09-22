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

//
// /BeoZone/System/Products
//

type SourceType struct {
	Type string `json:"type"`
}

type SourceID string

type ShortProduct struct {
	Jid          Jid    `json:"jid"`
	FriendlyName string `json:"friendlyName"`
}

type Source struct {
	Id           SourceID     `json:"id"`
	FriendlyName string       `json:"friendlyName,omitempty"`
	SourceType   SourceType   `json:"sourceType,omitempty"`
	Category     string       `json:"category,omitempty"`
	InUse        bool         `json:"inUse,omitempty"`
	Profile      string       `json:"profile,omitempty"`
	Linkable     bool         `json:"linkable,omitempty"`
	Product      ShortProduct `json:"product,omitempty"`
}

type Jid string

type Integrated struct {
	Role string `json:"role"`
	Jid  Jid    `json:"jid"`
}

type State string

const (
	StateIdle      State = "idle"
	StatePreparing State = "preparing"
	StatePlay      State = "play"
	StatePause     State = "pause"
	StateStop      State = "stop"
)

type ProductPrimaryExperience struct {
	Source   Source `json:"source"`
	Listener []Jid  `json:"listener"`
	State    State  `json:"state"`
}

type Product struct {
	Jid               Jid                       `json:"jid"`
	FriendlyName      string                    `json:"friendlyName"`
	Online            bool                      `json:"online"`
	PrimaryExperience *ProductPrimaryExperience `json:"primaryExperience"`
	Source            []Source                  `json:"source"`
	Integrated        *Integrated               `json:"integrated"`
	Integrate         bool                      `json:"integrate"`
	BorrowSource      bool                      `json:"borrowSource"`
	IrReroute         string                    `json:"irReroute"`
}

type ProductsResponse struct {
	Products []Product `json:"products"`
}

//
// /BeoZone/Zone/Sound/Volume/Speaker/Level
//

type SpeakerLevel struct {
	Level int `json:"level"`
}

type SpeakerMuted struct {
	Muted bool `json:"muted"`
}

//
// /BeoZone/Zone/PlayQueue
//

type Behaviour string

const (
	Planned   Behaviour = "planned"
	Impulsive Behaviour = "impulsive"
)

type Deezer struct {
	Id           int    `json:"id"`
	Availability string `json:"availability,omitempty"`
}

type Size string

const (
	Small  = "small"
	Medium = "medium"
	Large  = "large"
)

type Image struct {
	URL       string `json:"url"`
	Size      Size   `json:"size"`
	MediaType string `json:"mediatype"`
}

type Artist struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	NameNormalized string  `json:"nameNormalized"`
	Deezer         Deezer  `json:"deezer"`
	Image          []Image `json:"image"`
}

type Dlna struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type Track struct {
	Deezer               *Deezer  `json:"deezer,omitempty"` // Used for instantPlay
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	TrackNumber          int      `json:"trackNumber"`
	Duration             int      `json:"duration"`
	ArtistName           string   `json:"artistName,omitempty"`
	ArtistNameNormalized string   `json:"artistNameNormalized,omitempty"`
	Artist               []Artist `json:"artist"`
	Dlna                 *Dlna    `json:"dlna,omitempty"`
	Image                []Image  `json:"image"`
}

type BeoRadio struct {
	StationId string `json:"stationId"`
}

type Station struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	BeoRadio BeoRadio `json:"beoradio"`
	Image    []Image  `json:"image"`
}

type PlayQueueItemID string

type PlayQueueItem struct {
	Id        PlayQueueItemID `json:"id,omitempty"` // Included in the Track for instantPlay
	Behaviour Behaviour       `json:"behaviour"`
	Track     *Track          `json:"track,omitempty"`
	Station   *Station        `json:"station,omitempty"`
}

type Repeat string

const (
	RepeatOff         Repeat = "off"
	RepeatCurrentItem Repeat = "repeatCurrentItem"
	RepeatAll         Repeat = "repeatAll"
	RepeatUnknown     Repeat = "unknown"
)

type Random string

const (
	RandomUnknown Random = "unknown"
	RandomRandom  Random = "random"
	RandomOff     Random = "off"
)

type Type string

const (
	DeezerPlaylist Type = "deezerPlayList"
)

type Container struct {
	Type   Type   `json:"type"`
	Deezer Deezer `json:"deezer"`
}

type PlayQueueID string

const (
	PlayQueueIDMusic  PlayQueueID = "music"
	PlayQueueIDDeezer PlayQueueID = "deezer"
)

type PlayQueue struct {
	Id            PlayQueueID     `json:"id,omitempty"`
	Offset        int             `json:"offset,omitempty"`
	Count         int             `json:"count,omitempty"`
	StartOffset   int             `json:"startOffset,omitempty"`
	Total         int             `json:"total,omitempty"`
	PlayNowId     PlayQueueItemID `json:"playNowId,omitempty"`
	Random        Random          `json:"random,omitempty"`
	Repeat        Repeat          `json:"repeat,omitempty"`
	PlayQueueItem []PlayQueueItem `json:"playQueueItem"`
	Container     Container       `json:"container,omitempty"` // Used for instantPlay
}

type PlayQueueResponse struct {
	PlayQueue PlayQueue `json:"playQueue"`
}

type PlayQueueItemRequest struct {
	PlayQueueItem PlayQueueItem `json:"playQueueItem"`
}

//
// /BeoZone/Zone/PlayQueue/PlayPointer
//

type PlayPointer struct {
	PlayQueueItemId string `json:"playQueueItemId"`
	Position        int    `json:"position"`
}

type PlayPointerRequest struct {
	PlayPointer PlayPointer `json:"playPointer"`
}

//
// /BeoZone/Zone/ActiveSources
//

type IrMapping struct {
	Format  int `json:"format"`
	Unit    int `json:"unit"`
	Command int `json:"command"`
}

type ContentProtection struct {
	SchemeList []string `json:"schemeList"`
}

type EmbeddedBinary struct {
	SchemeList []string `json:"schemeList"`
}

type Listener struct {
	Jid Jid `json:"jid"`
}

type ListenerList struct {
	Listener []Listener `json:"listener"`
}

type ActiveSources struct {
	Primary    SourceID `json:"primary"`
	PrimaryJid Jid      `json:"primaryJid"`
}

type PrimaryExperience struct {
	Source               Source            `json:"source"`
	Category             string            `json:"category,omitempty"`
	Profile              string            `json:"profile,omitempty"`
	InUse                bool              `json:"inUse,omitempty"`
	Linkable             bool              `json:"linkable,omitempty"`
	RecommendedIrMapping []IrMapping       `json:"recommendedIrMapping,omitempty"`
	ContentProtection    ContentProtection `json:"contentProtection,omitempty"`
	EmbeddedBinary       EmbeddedBinary    `json:"embeddedBinary,omitempty"`
	Product              ShortProduct      `json:"product,omitempty"`
	ListenerList         ListenerList      `json:"listenerList,omitempty"`
}

type ActiveSourcesRequest struct {
	PrimaryExperience `json:"primaryExperience"`
}

type ActiveSourcesResponse struct {
	PrimaryExperience PrimaryExperience `json:"primaryExperience"`
	ActiveSources     ActiveSources     `json:"activeSources"`
}

type PrimaryExperienceRequest struct {
	Listener Listener `json:"listener"`
}
