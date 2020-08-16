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
// /BeoDevice
//

type ProductId struct {
	ProductType  string `json:"productType"`
	TypeNumber   string `json:"typeNumber"`
	SerialNumber string `json:"serialNumber"`
	ItemNumber   string `json:"itemNumber"`
}

type ProductFriendlyName struct {
	ProductFriendlyName string `json:"productFriendlyName"`
}

type Software struct {
	Version                     string `json:"version"`
	SoftwareUpdateProductTypeId int    `json:"softwareUpdateProductTypeId"`
}

type Hardware struct {
	Version string `json:"version"`
	Type    string `json:"type"`
	Item    string `json:"item"`
	Serial  string `json:"serial"`
	Bom     string `json:"bom"`
	Mac     string `json:"mac"`
}

type BeoDeviceInfo struct {
	ProductId           ProductId           `json:"productId"`
	ProductFamily       string              `json:"productFamily"`
	ProductFriendlyName ProductFriendlyName `json:"productFriendlyName"`
	Software            Software            `json:"software"`
	Hardware            Hardware            `json:"hardware"`
	AnonymousProductId  string              `json:"anonymousProductId"`
}

type BeoDeviceResponse struct {
	BeoDevice BeoDeviceInfo `json:"beoDevice"`
}

//
// /BeoDevice/powerManagement/standby
//

type PowerState string

const (
	PowerStateOn         PowerState = "on"
	PowerStateStandby    PowerState = "standby"
	PowerStateAllStandby PowerState = "allStandby"
	PowerStateReboot     PowerState = "reboot"
)

type Standby struct {
	PowerState PowerState `json:"powerState"`
}

type StandbyRequest struct {
	Standby Standby `json:"standby"`
}
