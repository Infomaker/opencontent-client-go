package requests

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/Infomaker/opencontent-client-go/host"
)

func TestGetUrlWithoutVersion(t *testing.T) {
  req := ObjectRequest{
    Host: host.NewOpenContentHost(),
    Uuid:"11111-1111",
    Version: -1,
  }

  assert.Equal(t, "http://localhost:8080/opencontent/objects?uuid=11111-1111", req.GetUrl())
}


func TestGetUrlWithVersion(t *testing.T) {
  req := ObjectRequest{
    Host: host.NewOpenContentHost(),
    Uuid:"11111-1111",
    Version: 10,
  }

  assert.Equal(t, "http://localhost:8080/opencontent/objects?uuid=11111-1111&version=10", req.GetUrl())
}