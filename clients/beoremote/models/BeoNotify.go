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
	NotificationTypeSource               NotificationType = "SOURCE"
	NotificationTypeNowPlayingEnded      NotificationType = "NOW_PLAYING_ENDED"
	NotificationTypeProgressInformation  NotificationType = "PROGRESS_INFORMATION"
	NotificationTypeVolume               NotificationType = "VOLUME"
	NotificationTypeSoftwareUpdateStatus NotificationType = "SOFTWARE_UPDATE_STATUS"
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
// type: PROGRESS_INFORMATION
// kind: playing
//

type ProgressInformationData struct {
	State string `json:"state"`
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
