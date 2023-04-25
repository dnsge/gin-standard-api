# gin-standard-api

gin-standard-api is a utility library for standardizing API responses and error handling in [Gin](https://github.com/gin-gonic/gin) applications. The API structure is as follows:


### Successful Request (2xx)
```json5
{
  "status": 200, // HTTP status code, required
  "data": ... // Body data, optional (any JSON value)
}
```

### Unsuccessful Response (4xx/5xx)
```json5
{
  "status": 403, // HTTP status code, required
  "error": "Forbidden", // Generic error message, required
  "message": "You do not have permission to do that." // Detailed error message, optional.
}
```

## Internal API

Instead of using a regular `gin.HandlerFunc`, you can instead use an `sapi.Handler`. The difference is that an `sapi.Handler` returns `(sapi.Res, sapi.Err)`, indicating either a successful response or an error response.

Handler Example:

```go
package web

import (
	"github.com/dnsge/gin-standard-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HandleCurrentTime(c *gin.Context) (sapi.Res, sapi.Err) {
	now := time.Now()
	if now.Second() == 0 {
		// By returning a sapi.Error instance, we make control flow more explicit.
		return nil, sapi.ErrorStatus(http.StatusServiceUnavailable)
	}
	
	return sapi.Data(http.StatusOK, now.Unix()), nil
}

func Register(r gin.IRoutes) {
	r.GET("/now", sapi.Handle(HandleCurrentTime))
}
```

Should you need to write a response outside a `sapi.Handler`, use the API like so:

```go
package web

import (
	"github.com/dnsge/gin-standard-api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DoSomething(c *gin.Context) {
    sapi.Status(http.StatusNotFound).WriteResponse(c)
}
```