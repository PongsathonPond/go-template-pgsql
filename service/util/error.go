package util

import (
	"fmt"
)

var (
	UnknownErr                    = defineErr("UNKNOWN", "UNKNOWN")
	ConvertInputToDomainErr       = defineErr("CONVERSION", "INPUT_TO_DOMAIN")
	JsonDecoderErr                = defineErr("DECODE", "JSON")
	QueryStrDecoderErr            = defineErr("DECODE", "QUERY_STRING")
	ValidationErr                 = defineErr("VALIDATION", "ERROR")
	ValidationCreateErr           = defineErr("VALIDATION", "CREATE")
	ValidationUpdateErr           = defineErr("VALIDATION", "UPDATE")
	ValidationParamOptionErr      = defineErr("VALIDATION", "PARAM_OPTIONS")
	ValidationHeaderErr           = defineErr("VALIDATION", "HEADER")
	ValidationTokenExpired        = defineErr("VALIDATION", "TOKEN_EXPIRED")
	ValidationActivationErr       = defineErr("VALIDATION", "ACTIVATION")
	RepoListErr                   = defineErr("REPOSITORY", "LIST")
	RepoCreateErr                 = defineErr("REPOSITORY", "CREATE")
	RepoReadErr                   = defineErr("REPOSITORY", "READ")
	RepoUpdateErr                 = defineErr("REPOSITORY", "UPDATE")
	RepoDeleteErr                 = defineErr("REPOSITORY", "DELETE")
	RepoFindErr                   = defineErr("REPOSITORY", "Find")
	RepoCountErr                  = defineErr("REPOSITORY", "COUNT")
	ServiceGRPCErr                = defineErr("SERVICE", "GRPC")
	UsersUnverifiedErr            = defineErr("USERS", "UNVERIFIED")
	ConditionActivateExpiredAtErr = defineErr("CONDITION", "ACTIVATE_EXPIRED_AT")
	ConditionStillInUseErr        = defineErr("CONDITION", "STILL_IN_USE")
)

type Error struct {
	Cause   error  `json:"cause"`
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
}

type ErrorFunc func(cause error) (err *Error)

func defineErr(code string, subCode string) (errFn ErrorFunc) {
	return func(cause error) (err *Error) {
		return &Error{
			Cause:   cause,
			Code:    code,
			SubCode: subCode,
		}
	}
}

func TypeOfErr(err error) (typeOf *Error) {
	if e0, ok := err.(*Error); ok {
		return e0
	}
	return &Error{Cause: err}
}

func (e *Error) Error() (msg string) {
	causeStr := ""
	if e.Cause != nil {
		causeStr = e.Cause.Error()
	}
	return fmt.Sprintf("%s (%s:%s)", causeStr, e.Code, e.SubCode)
}

func (e *Error) IsType(errFn ErrorFunc) (is bool) {
	test := errFn(nil)
	return e.Code == test.Code && e.SubCode == test.SubCode
}
