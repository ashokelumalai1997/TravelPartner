package clientcommon

import (
	"errors"
	"net/http"
)

func RespOK(c RespJSON, o interface{}) {
	c.JSON(http.StatusOK, o)
}

func RespISE(c RespJSON, err error) {
	o := errors.New("err")
	c.JSON(http.StatusInternalServerError, o)
}
