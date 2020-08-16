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
// /BeoHome/trigger/timerList
//

type ActionType string

const (
	AddToPlayQueue  ActionType = "addToPlayQueue"
	SetActiveSource ActionType = "setActiveSource"
	SetPowerState   ActionType = "setPowerState"
)

type Action struct {
	PlayQueueItem PlayQueueItem `json:"playQueueItem,omitempty"`
	PowerState    PowerState    `json:"powerState,omitempty"`
}

type Persistent string

const (
	Yes Persistent = "yes"
)

type Day string

const (
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
	Sunday    Day = "sunday"
)

type Timer struct {
	Id           string     `json:"id,omitempty"`
	FriendlyName string     `json:"friendlyName"`
	Time         string     `json:"time"`
	Active       string     `json:"active"`
	Recurring    []Day      `json:"recurring"`
	Persistent   string     `json:"persistent"`
	ActionType   ActionType `json:"actionType"`
	ActionValue  Action     `json:"actionValue"`
}

type TimerList struct {
	Timer []Timer `json:"timer"`
}

type TimerListResponse struct {
	TimerList TimerList `json:"timerList"`
}
