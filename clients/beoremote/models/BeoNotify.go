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

import "encoding/json"

type NotificationType string

const (
	NotificationTypeSource                  NotificationType = "SOURCE"
	NotificationTypeSourceExperienceChanged NotificationType = "SOURCE_EXPERIENCE_CHANGED"
	NotificationTypeNowPlayingEnded         NotificationType = "NOW_PLAYING_ENDED"
	NotificationTypeNowPlayingStoredMusic   NotificationType = "NOW_PLAYING_STORED_MUSIC"
	NotificationTypeNowPlayingNetRadio      NotificationType = "NOW_PLAYING_NET_RADIO"
	NotificationTypePlayQueueChanged        NotificationType = "PLAY_QUEUE_CHANGED"
	NotificationTypeProgressInformation     NotificationType = "PROGRESS_INFORMATION"
	NotificationTypeVolume                  NotificationType = "VOLUME"
	NotificationTypeSoftwareUpdateStatus    NotificationType = "SOFTWARE_UPDATE_STATUS"
)

type NotificationKind string

const (
	NotificationKindSource   = "source"
	NotificationKindPlaying  = "playing"
	NotificationKindRenderer = "renderer"
	NotificationKindDevice   = "device"
)

type Notification struct {
	Timestamp string           `json:"timestamp"`
	Type      NotificationType `json:"type"`
	Kind      string           `json:"kind"`
	Data      json.RawMessage  `json:"data"`
}

type NotificationWrapper struct {
	Notification Notification `json:"notification"`
}

//
// type: SOURCE
// kind: source
//

type PrimaryExperienceNotification struct {
	Source               Source            `json:"source"`
	Category             string            `json:"category,omitempty"`
	Profile              string            `json:"profile,omitempty"`
	InUse                bool              `json:"inUse,omitempty"`
	Linkable             bool              `json:"linkable,omitempty"`
	RecommendedIrMapping []IrMapping       `json:"recommendedIrMapping,omitempty"`
	ContentProtection    ContentProtection `json:"contentProtection,omitempty"`
	EmbeddedBinary       EmbeddedBinary    `json:"embeddedBinary,omitempty"`
	Listener             []Jid             `json:"listener,omitempty"`
	LastUsed             string            `json:"lastUsed,omitempty"`
	State                State             `json:"state"`
}

type SourceData struct {
	Primary           SourceID                      `json:"primary"`           // Primary is the source ID.
	PrimaryJid        Jid                           `json:"primaryJid"`        // PrimaryJid is the product ID.
	PrimaryExperience PrimaryExperienceNotification `json:"primaryExperience"` // PrimaryExperience is the primary experience.
}

//
// type: SOURCE_EXPERIENCE_CHANGED
// kind: source
//

type SourceExperienceChangedData struct {
	PrimaryExperience PrimaryExperienceNotification `json:"primaryExperience"`
}

//
// type: NOW_PLAYING_STORED_MUSIC
// kind: playing
//

type NowPlayingStoredMusicData struct {
	Name            string          `json:"name"`            // Name is the track title.
	Album           string          `json:"album"`           // Album is the album.
	Artist          string          `json:"artist"`          // Artist is the artist.
	TrackID         string          `json:"trackId"`         // TrackID is the deezer track ID.
	TrackImage      []Image         `json:"trackImage"`      // TrackImage is a list of track images.
	PlayQueueID     PlayQueueID     `json:"playQueueId"`     // PlayQueueID is the name if the queue playing.
	PlayQueueItemID PlayQueueItemID `json:"playQueueItemId"` // PlayQueueItemID is the item ID in the play queue.
}

//
// type: NOW_PLAYING_NET_RADIO
// kind: playing
//

type NowPlayingNetRadioData struct {
	Name            string      `json:"name"`            // Name is the track title.
	LiveDescription string      `json:"liveDescription"` // LiveDescription is the description.
	StationID       string      `json:"stationId"`       // StationID is the station ID.
	Image           []Image     `json:"image"`           // Image is the logo for the station.
	PlayQueueID     PlayQueueID `json:"playQueueId"`     // PlayQueueID is the name if the queue playing.
}

//
// type: PROGRESS_INFORMATION
// kind: playing
//

type ProgressInformationData struct {
	State           State           `json:"state"`
	Position        int             `json:"position"`
	TotalDuration   int             `json:"totalDuration"`
	SeekSupported   bool            `json:"seekSupported"`
	PlayQueueID     PlayQueueID     `json:"playQueueId"`
	PlayQueueItemID PlayQueueItemID `json:"playQueueItemId"`
}

//
// type: PLAY_QUEUE_CHANGED
// kind: playing
//

type PlayQueueChangedData struct {
	Revision    int         `json:"revision"`
	PlayQueueID PlayQueueID `json:"playQueueId"`
}

//
// type: VOLUME
// kind: renderer
//

type Range struct {
	Minimum int `json:"minimum"`
	Maximum int `json:"maximum"`
}

type Speaker struct {
	Level int   `json:"level"`
	Muted bool  `json:"muted"`
	Range Range `json:"range"`
}

type VolumeData struct {
	Speaker Speaker `json:"speaker"`
}

//
// type: SOFTWARE_UPDATE_STATUS
// kind: device
//

type SoftwareUpdateStatusData struct {
	State string `json:"state"`
}
