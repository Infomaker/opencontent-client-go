package main

import (
  "github.com/hansbringert/opencontent-client/ochost"
  "github.com/hansbringert/opencontent-client/requests"
  "fmt"
  "time"
)

func details(host ochost.OpenContentHost, uuid string, propertyFields... string) {
  req :=requests.NewPropertiesRequest(host)
  req.Uuid = uuid

  for _, property := range propertyFields {
    req.AddProperty(property)
  }

  properties, err := req.GetProperties()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    return
  }

  jsonBytes, err := properties.ToJson()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    return
  }

  fmt.Println(string(jsonBytes))
}


func EventTail(interval string, host ochost.OpenContentHost, eventId int) {
  duration, _ := time.ParseDuration(interval)
  ticker := time.NewTicker(duration)
  quit := make(chan struct{})
  var EventInProgress bool
  go func() {
    for {
      select {
      case <-ticker.C:
        if ! EventInProgress {
          EventInProgress = true
          req := requests.NewEventRequest(host)
          req.SetEventId(eventId)
          events, _ := req.GetEvents()

          for _, event := range events.Events {
            fmt.Println(event.Created, " ", event.Id, " ", event.Uuid, " ", event.Content.ContentType)
            details(host, event.Uuid, "uuid")
          }
          eventId = events.LastEvent(eventId)
          EventInProgress = false
        }
      case <-quit:
        return
      }
    }
  }()
}




func main() {
  host := ochost.NewOpenContentHost()

  host.Host = "xlibris.oc.gota.infomaker.io"
  host.Password = "gurus2,Beaks"

  EventTail("2s", host, -1)
  forever := make(chan bool)
  fmt.Println(" [*] Processing updates every ", "2s", " To exit press CTRL+C")
  <-forever

  //req := requests.NewEventRequest(host)
  //req.SetEventId(0)
  //
  //events, err := req.GetEvents()
  //if (err!= nil)  {
  //  fmt.Println("ERROR:", err.Error())
  //  os.Exit(1)
  //}
  //
  //jsonBytes, _ := events.ToJson()
  //fmt.Println(string(jsonBytes))

}
