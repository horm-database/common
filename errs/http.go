package errs

import (
	"net/http"
)

// ErrToHTTPStatus maps from framework errs retcode to http status code.
var ErrToHTTPStatus = map[int]int{
	ErrServerNoService: http.StatusNotFound,
	ErrServerNoFunc:    http.StatusNotFound,
	ErrServerTimeout:   http.StatusGatewayTimeout,
	ErrServerOverload:  http.StatusTooManyRequests,

	ErrAppCallLimited: http.StatusForbidden,

	// 权限异常
	ErrAuthFail:        http.StatusUnauthorized,
	ErrAppidNotFound:   http.StatusUnauthorized,
	ErrHasNoTableRight: http.StatusUnauthorized,
	ErrHasNoDBRight:    http.StatusUnauthorized,
	ErrTableVerify:     http.StatusUnauthorized,

	ErrServerDecode:         http.StatusBadRequest,
	ErrParamInvalid:         http.StatusBadRequest,
	ErrParamEmpty:           http.StatusBadRequest,
	ErrParamMiss:            http.StatusBadRequest,
	ErrParamType:            http.StatusBadRequest,
	ErrParamValue:           http.StatusBadRequest,
	ErrNotFindName:          http.StatusBadRequest,
	ErrUnitNameEmpty:        http.StatusBadRequest,
	ErrRepeatNameAlias:      http.StatusBadRequest,
	ErrRefererNotFound:      http.StatusBadRequest,
	ErrRefererUnitFailed:    http.StatusBadRequest,
	ErrRefererResultType:    http.StatusBadRequest,
	ErrRefererFieldNotExist: http.StatusBadRequest,
	ErrFormatData:           http.StatusBadRequest,
}
