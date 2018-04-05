package main

import (
  "github.com/hansbringert/opencontent-client/ochost"
  "fmt"
  "github.com/hansbringert/opencontent-client/requests"
  "github.com/Sirupsen/logrus"
  "os"
  "bufio"
)

func main() {

  host := ochost.NewOpenContentHost()

  host.Host = "xlibris.oc.gota.infomaker.io"
  host.Password = "gurus2,Beaks"

  req := requests.NewPropertiesRequest(host)
  req.AddProperty("uuid")

  file, err := os.Open("/local/infomaker/go/work/src/github.com/hansbringert/opencontent-client/examples/exists/uuids.txt")
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
