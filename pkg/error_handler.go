package pkg

import (
	"erdemkayacomtr/constant"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseStatus constant.ResponseStatus) {
	PanicException_(responseStatus.GetResponseMessage(), responseStatus.GetResponseMessage())
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		//msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, false, Null()))
			c.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, false, Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, false, Null()))
			c.Abort()
		}
	}
}
