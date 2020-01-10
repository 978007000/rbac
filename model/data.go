package model

import (
	"encoding/json"
	// "fmt"
)

type Data struct {
	CustomerNumber   string          `json:"customerNumber"`
	ThingName        string          `json:"thingName"`
	ThingDisplayName string          `json:"thingDisplayName"`
	ThingSerial      string          `json:"serial"`
	PropertyName     string          `json:"propertyName"`
	DataType         DataType        `json:"dataType"`
	Value            string          `json:"value"`
	PropertyId       int             `json:"propertyId"`
	ThingId          int             `json:"thingId"`
	Alerts           []AlerDataModel `json:"alerts"`
}

func (b Data) String() string {
	e, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(e)
}

type ObjectData struct {
	Value interface{} `json:"value"`
}

func (b ObjectData) String() string {
	e, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(e)
}

type ObjectParams map[string]interface{}

func (b ObjectParams) Convert2Json(value interface{}) string {
	e, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(e)
}

type AlerDataModel struct {
	Name         string    `json:"name" `
	Value        string    `json:"value" `
	SettingValue string    `json:"settingValue" `
	AlertModule  string    `json:"alertModule" `
	Status       int       `json:"status"`
	PriorityId   int       `json:"priorityId"`
	Triggers     []Trigger `json:"triggers"`
}

type Trigger struct {
	Namespace    string `json:"namespace,omitempty"`
	Micro    string `json:"micro,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	Body     string `json:"body,omitempty"`
	Params   string `json:"params,omitempty"`
}
