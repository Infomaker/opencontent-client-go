package main

import (
  "fmt"
  "github.com/Sirupsen/logrus"
  "os"
  "bufio"
  "github.com/Infomaker/opencontent-client-go/host"
  "github.com/Infomaker/opencontent-client-go/requests"
)

func main() {

  host := host.NewOpenContentHost()

  host.Host = "xlibris.oc.gota.infomaker.io"
  host.Password = "gurus2,Beaks"

  req := requests.NewPropertiesRequest(host)
  req.AddProperty("uuid")

  file, err := os.Open("/local/infomaker/go/work/src/github.com/Infomaker/opencontent-client-go/examples/exists/uuids.txt")
  if err != nil {
    logrus.WithError(err).Warn("Problem occurred when reading file")
  } else {
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      uuid := scanner.Text()
      req.SetUuid(uuid)
      _, err := req.GetProperties()
      if err != nil {
        fmt.Println("ERROR:", uuid, err.Error())
      }
      
    }
  }
}
