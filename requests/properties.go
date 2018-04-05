package requests

import (
  "github.com/go-resty/resty"
  "fmt"
  "encoding/json"
  "strings"
  "strconv"
  "errors"
  "github.com/Infomaker/opencontent-client-go/host"
)

// Struct for the proeprties request
type PropertiesRequest struct {
  Host       host.OpenContentHost
  Uuid       string
  Version    int
  Name       []string // should not be used, use properties instead
  Properties []string
  Filters    []string
  UseIndex   bool
}

// Creates an empty Properties request i.e. no arguments set
func NewPropertiesRequest(host host.OpenContentHost) PropertiesRequest {
  req := PropertiesRequest{}
  req.Host = host
  req.Version = -1
  return req
}

// Sets the uuid for which properties is wanted
func (req *PropertiesRequest) SetUuid(uuid string) {
  req.Uuid = uuid
}

// Sets the uuid for which properties is wanted
func (req *PropertiesRequest) SetVersion(version int) {
  req.Version = version
}

func (req *PropertiesRequest) AddProperty(property string) {
  if len(property) == 0 {
    return
  }
  // for _, p := range strings.Split(property, ",") {
  req.Properties = append(req.Properties, property)
  // }
  return
}

func (req *PropertiesRequest) AddName(name string) {
  if len(name) == 0 {
    return
  }
  // for _, p := range strings.Split(property, ",") {
  req.Name = append(req.Name, name)
  // }
  return
}

func (req *PropertiesRequest) getUrl() string {
  httpQuery := "?"

  if len(req.Properties) > 0 {
    httpQuery = fmt.Sprint(httpQuery, "properties=")
    for _, p := range req.Properties {
        httpQuery = fmt.Sprint(httpQuery, p, ",")
    }
    // remove trailing comma
    httpQuery = strings.TrimSuffix(httpQuery, ",")
    // add ampersand
    httpQuery = fmt.Sprint(httpQuery, "&")
  }

  if len(req.Name) > 0 {
    for _, p := range req.Name {
      httpQuery = fmt.Sprint(httpQuery, "name=" + p + "&")
    }
  }

  if req.Version >= 0 {
    httpQuery = fmt.Sprint(httpQuery, "version=", strconv.Itoa(req.Version), "&")
  }

  return fmt.Sprintf("%s/objects/%s/properties%s", req.Host.CreateUrl(), req.Uuid, strings.TrimSuffix(httpQuery, "&"))
}

func (req *PropertiesRequest) execute() (*resty.Response, error) {
  return resty.SetDisableWarn(true).R().SetBasicAuth(req.Host.User, req.Host.Password).Get(req.getUrl())
}

func (req *PropertiesRequest) GetProperties() (PropertiesResponse, error) {
  response, err := req.execute()
  if err != nil {
    return PropertiesResponse{}, err
  }
  if response.StatusCode() != 200 {
    return PropertiesResponse{}, errors.New(response.Status() + " " + string(response.Body()))
  }
  var properties PropertiesResponse
  err = json.Unmarshal(response.Body(), &properties)
  if err != nil {
    return properties, err
  }
  return properties, nil
}

type PropertiesResponse struct {
  ContentType string `json:"contentType"`
  Editable    bool `json:"editable"`
  Properties  [] Property `json:"properties"`
}

type Property struct {
  Name        string `json:"name"`
  Type        string `json:"type"`
  MultiValued bool `json:"multiValued"`
  ReadOnly    bool `json:"readOnly"`
  Values      []interface{} `json:"values"`
}

func (hr *PropertiesResponse) ToJson() ([]byte, error) {
  bytes, err := json.MarshalIndent(hr, " ", " ", )
  return bytes, err
}

func PropertiesResponseFromBytes(jsonBytes []byte) (PropertiesResponse, error) {
  health := PropertiesResponse{}
  err := json.Unmarshal(jsonBytes, &health)
  return health, err
}

