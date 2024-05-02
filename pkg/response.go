package pkg

import (
	"coffeeshop/config"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code        int         `json:"-"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data,omitempty"`
	Meta        interface{} `json:"meta,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
	ctx.Abort()
}

func NewRes(code int, data *config.Result) *Response {
	var response = Response{
		Code:   code,
		Status: getStatus(code),
	}

	if response.Code >= 400 {
		response.Description = data.Data
	} else if data.Message != nil {
		response.Description = data.Message
	} else {
		response.Data = data.Data
	}

	if data.Meta != nil {
		response.Meta = data.Meta
	}

	return &response
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 403:
		desc = "Forbidden"
	case 404:
		desc = "Not Found"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
