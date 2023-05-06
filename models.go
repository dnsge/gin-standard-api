package sapi

import (
	"github.com/gin-gonic/gin"
)

type rawStatusResponse int

func (r rawStatusResponse) WriteResponse(c *gin.Context) {
	c.Header("Content-Length", "0")
	c.Status(int(r))
}

type rawErrorStatusResponse int

func (r rawErrorStatusResponse) WriteError(c *gin.Context) {
	c.Header("Content-Length", "0")
	c.Status(int(r))
}

type dataResponse struct {
	Status int `json:"status"`
	Data   any `json:"data,omitempty"`
}

func (d *dataResponse) WriteResponse(c *gin.Context) {
	c.JSON(d.Status, d)
}

func (d *dataResponse) GetData() any {
	return d.Data
}

type redirectResponse struct {
	Status   int
	Location string
}

func (r *redirectResponse) WriteResponse(c *gin.Context) {
	c.Redirect(r.Status, r.Location)
}

type errorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func (e *errorResponse) WriteError(c *gin.Context) {
	c.JSON(e.Status, e)
}

type stringResponse struct {
	Status int
	String string
}

func (s *stringResponse) WriteResponse(c *gin.Context) {
	c.String(s.Status, s.String)
}

func (s *stringResponse) WriteError(c *gin.Context) {
	c.String(s.Status, s.String)
}
