package errs

import (
	"net/http"
)

// ErrToHTTPStatus maps from framework errs retcode to http status code.
var ErrToHTTPStatus = map[int]int{
	RetServerNoService: http.StatusNotFound,
	RetServerNoFunc:    http.StatusNotFound,
	RetServerTimeout:   http.StatusGatewayTimeout,
	RetServerOverload:  http.StatusTooManyRequests,

	RetAppCallLimited: http.StatusForbidden,

	// 权限异常
	RetServerAuthFail:    http.StatusUnauthorized,
	RetNotFindAppid:      http.StatusUnauthorized,
	RetHasNoTableRight:   http.StatusUnauthorized,
	RetHasNoDBRight:      http.StatusUnauthorized,
	RetTableVerifyFailed: http.StatusUnauthorized,

	RetServerDecodeFail:     http.StatusBadRequest,
	RetParamInvalid:         http.StatusBadRequest,
	RetParamEmpty:           http.StatusBadRequest,
	RetParamMiss:            http.StatusBadRequest,
	RetParamType:            http.StatusBadRequest,
	RetParamValue:           http.StatusBadRequest,
	RetNotFindName:          http.StatusBadRequest,
	RetUnitNameEmpty:        http.StatusBadRequest,
	RetRepeatNameAlias:      http.StatusBadRequest,
	RetNotFindReferer:       http.StatusBadRequest,
	RetRefererUnitFailed:    http.StatusBadRequest,
	RetRefererResultType:    http.StatusBadRequest,
	RetRefererFieldNotExist: http.StatusBadRequest,
	RetFormatDataError:      http.StatusBadRequest,
}
