package sapitest

import (
	sapi "github.com/dnsge/gin-standard-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetResStatus(t *testing.T) {
	res := sapi.RawStatus(http.StatusCreated)
	assert.Equal(t, http.StatusCreated, GetResStatus(res))
	res = sapi.Status(http.StatusCreated)
	assert.Equal(t, http.StatusCreated, GetResStatus(res))
}

func TestGetResBody(t *testing.T) {
	res := sapi.Status(http.StatusCreated)
	assert.JSONEq(t, `{"status": 201}`, GetResBody(res))
}
