package sapitest

import (
	sapi "github.com/dnsge/gin-standard-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetErrStatus(t *testing.T) {
	err := sapi.RawErrorStatus(http.StatusInternalServerError)
	assert.Equal(t, http.StatusInternalServerError, GetErrStatus(err))
	err = sapi.ErrorStatus(http.StatusInternalServerError)
	assert.Equal(t, http.StatusInternalServerError, GetErrStatus(err))
}

func TestGetErrBody(t *testing.T) {
	err := sapi.Error(http.StatusInternalServerError, "Oh no")
	assert.JSONEq(t, `{"status": 500, "error": "Oh no"}`, GetErrBody(err))
}
