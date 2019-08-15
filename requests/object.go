package requests

import (
  "github.com/Infomaker/opencontent-client-go/host"
  "fmt"
  "gopkg.in/resty.v1"
  "strings"
)

type ObjectRequest struct {
  Host    host.OpenContentHost
  Uuid    string
  Version int
}

func (req *ObjectRequest) GetUrl() string {
  httpQuery := "?"
  if req.Version > 0 {
    httpQuery = fmt.Sprint(httpQuery, "version=", req.Version, "&")
  }

  return fmt.Sprint(req.Host.CreateUrl(), "/objects/", req.Uuid, strings.TrimSuffix(httpQuery, "&"))
}

func (req *ObjectRequest) Execute() (*resty.Response, error) {
  return resty.SetDisableWarn(true).R().SetBasicAuth(req.Host.User, req.Host.Password).Get(req.GetUrl())
}

func (req *ObjectRequest) GetObject() ([]byte, error) {
  response, err := req.Execute()

  fmt.Println(req.GetUrl())
  if host.IsResponseOk(*response, err, 200) {
    if err != nil {
      return []byte{}, err
    }
  }
  return response.Body(), nil
}
