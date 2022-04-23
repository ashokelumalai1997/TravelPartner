package clientcommon

type RespJSON interface {
	JSON(code int, obj interface{})
}

type IData interface {
	Data(code int, contentType string, data []byte)
}
