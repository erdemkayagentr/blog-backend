package pkg

import (
	"erdemkayacomtr/constant"
	"erdemkayacomtr/domain/models"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T, message string) models.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseMessage(), responseStatus == constant.Success, data)
}

func BuildResponse_[T any](message string, succeeded bool, data T) models.ApiResponse[T] {
	return models.ApiResponse[T]{
		Message:   message,
		Succeeded: succeeded,
		Data:      data,
	}
}
