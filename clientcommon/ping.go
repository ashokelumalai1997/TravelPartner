package clientcommon

import "net/http"

func Ping(c IData) {
	c.Data(http.StatusOK, "text/plain", []byte("pong"))
}
