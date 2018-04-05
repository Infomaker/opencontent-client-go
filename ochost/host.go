package ochost

import (
  "os"
  "net/url"
  "fmt"
  "github.com/go-resty/resty"
)

type OpenContentHost struct {
  Scheme   string `json:"schema"`
  Host     string `json:"host"`
  User     string `json:"user"`
  Password string `json:"password"`
  Port     string  `json:"port"`
  Baseurl  string  `json:"baseurl"`
}

func getEnvironmentVariable(name, defaultValue string) string {
  value := os.Getenv(name)
  if len(value) == 0 {
    return defaultValue
  }
  return value
}


// Creates a new Open Content host, if environment variables set they are used
// Environment variables :
//
func NewOpenContentHost() OpenContentHost {

  ocHost := OpenContentHost{}
  ocHost.Scheme = getEnvironmentVariable("OC_SCHEME", "http")
  ocHost.Host = getEnvironmentVariable("OC_HOST", "localhost")
  ocHost.User = getEnvironmentVariable("OC_USER", "admin")
  ocHost.Password = getEnvironmentVariable("OC_PWD", "admin")
  ocHost.Port = getEnvironmentVariable("OC_PORT", "8080")
  ocHost.Baseurl = getEnvironmentVariable("OC_BASEURL", "/opencontent")
  return ocHost
}

func (host *OpenContentHost) CreateUrl() string {
  url := url.URL{}
  url.Scheme = host.Scheme
  url.Host = fmt.Sprint(host.Host, ":", host.Port)
  url.Path = host.Baseurl
  return url.String()
}

func statusOk(statusCode int, status []int) bool {
  for _, s := range status {
    if statusCode == s {
      return true
    }
  }
  return false
}

func IsResponseOk(response resty.Response, err error, okstatus... int) bool {
  if err != nil {
    fmt.Println("ERROR: ", err.Error())
    return false
  }

  if ! statusOk(response.StatusCode(), okstatus) {
    var body string
    if len(response.Body()) > 0 {
      body = fmt.Sprint("Body: ", string(response.Body()))
    }
    fmt.Println("ERROR: ", response.Status(), body)
    return false
  }

  return true
}
