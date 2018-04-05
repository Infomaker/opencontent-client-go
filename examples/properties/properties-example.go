package main

import (
  "fmt"
  "os"
  "github.com/Infomaker/opencontent-client-go/host"
  "github.com/Infomaker/opencontent-client-go/requests"
)

func main() {
  host := host.NewOpenContentHost()

  host.Host = "192.168.1.109"

  req := requests.NewPropertiesRequest(host)
  req.SetUuid("f29ee1d9-2e11-4fbd-b323-675f9289edab")
  req.AddProperty("uuid,version")

  properties, err  := req.GetProperties()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    os.Exit(1)
  }

  // Print values
  for _, property := range properties.Properties {
    fmt.Println(property.Name, "=", property.Values)
  }
  // propertiesJsonBytes, err := properties.ToJson()
  // fmt.Println(string(propertiesJsonBytes))
}
