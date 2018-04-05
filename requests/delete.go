package requests

import (
  "github.com/go-resty/resty"
  "fmt"
  "github.com/hansbringert/opencontent-client/ochost"
)

type DeleteRequest struct {
  Host    ochost.OpenContentHost
  BaseUrl string
  Uuid    string
}

func NewDeleteRequest(host ochost.OpenContentHost) DeleteRequest {
  req := DeleteRequest{}
  req.Host = host
  return req
}

func (req *DeleteRequest) GetUrl() string {
  return fmt.Sprint(req.Host.CreateUrl(), "/objects/", req.Uuid)
}

func (req *DeleteRequest) SetUuid(uuid string) {
  req.Uuid = uuid
}

func (req *DeleteRequest) execute() (*resty.Response, error) {
  return resty.SetDisableWarn(true).R().SetBasicAuth(req.Host.User, req.Host.Password).Delete(req.GetUrl())
}

func (req *DeleteRequest) Delete() error{
  response, err := req.execute()
  if response.StatusCode() == 409 {
    // Retry when we have 409
    return req.Delete()
  } else if ! ochost.IsResponseOk(*response, err, 200) {
    return nil
  }
  return err
}

func DeleteUuid(host ochost.OpenContentHost, uuid string) error{
  req := NewDeleteRequest(host)
  req.SetUuid(uuid)
  return req.Delete()
}




