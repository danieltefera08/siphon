package logic

import (
	"siphon/app/djigateway/djigateway"
	"siphon/common/djisdk"
)

func errRes(tid string, err error) *djigateway.CommonRes {
	if djiErr, ok := djisdk.IsDJIError(err); ok {
		return &djigateway.CommonRes{
			Code:       -1,
			Message:    djiErr.Message,
			Tid:        tid,
			ReasonCode: int32(djiErr.Code),
		}
	}
	return &djigateway.CommonRes{Code: -1, Message: err.Error(), Tid: tid}
}

func okRes(tid string) *djigateway.CommonRes {
	return &djigateway.CommonRes{Code: 0, Message: "success", Tid: tid}
}
