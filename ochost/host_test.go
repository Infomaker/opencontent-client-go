package ochost

import (
  "testing"
  "github.com/stretchr/testify/assert"
)



func TestNewOpenContentHost(t *testing.T) {

  host := NewOpenContentHost()

  assert.Equal(t, "localhost", host.Host)
  assert.Equal(t, "admin", host.User)
  assert.Equal(t, "admin", host.Password)
  assert.Equal(t, "8080", host.Port)
  assert.Equal(t, "/opencontent", host.Baseurl)
}

func TestOpenContentHost_GetBaseUrl(t *testing.T) {
  host := NewOpenContentHost()
  assert.Equal(t, "http://localhost:8080/opencontent", host.CreateUrl())
}


