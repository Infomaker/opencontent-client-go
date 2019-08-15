package requests

import (
  "fmt"
  "gopkg.in/resty.v1"
  "strings"
  "encoding/json"
  "github.com/Infomaker/opencontent-client-go/host"
)

type EventRequest struct {
  Host  host.OpenContentHost
  Event int
}

func NewEventRequest(host host.OpenContentHost) EventRequest {
  req := EventRequest{}
  req.Host = host
  req.Event = -1
  return req
}

func (req *EventRequest) GetUrl() string {
  httpQuery := "?"
  httpQuery = fmt.Sprint(httpQuery, "event=", req.Event)

  return fmt.Sprint(req.Host.CreateUrl(), "/eventlog", strings.TrimSuffix(httpQuery, "&"))
}

func (req *EventRequest) SetEventId(eventid int) {
  req.Event = eventid
}

func (req *EventRequest) execute() (*resty.Response, error) {
  return resty.SetDisableWarn(true).R().SetBasicAuth(req.Host.User, req.Host.Password).Get(req.GetUrl())
}

type Events struct {
  Events [] Event `json:"events"`
}

func (events *Events) ToJson() ([]byte, error) {
  bytes, err := json.MarshalIndent(events, " ", " ", )
  return bytes, err
}

func (events *Events) LastEvent(eventId int) int {
  for _, currentEvent := range events.Events {
    if currentEvent.Id > eventId {
      eventId = currentEvent.Id
    }
  }
  return eventId
}


type Event struct {
  Id        int     `json:"id"`
  Uuid      string  `json:"uuid"`
  EventType string  `json:"eventType"`
  Created   string  `json:"created"`
  Content   Content `json:"content"`
}

type Content struct {
  Uuid        string `json:"uuid"`
  Version     int    `json:"version"`
  Created     string `json:"created"`
  Source      string `json:"source"`
  ContentType string `json:"contentType"`
}

func (req *EventRequest) GetEvents() (*Events, error) {
  response, err := req.execute()
  events := &Events{}
  if host.IsResponseOk(*response, err, 200) {
    err := json.Unmarshal(response.Body(), events)
    if err != nil {
      return events, err
    }
  }
  return events, nil
}

