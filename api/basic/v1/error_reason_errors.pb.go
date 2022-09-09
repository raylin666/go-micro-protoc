// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsServerError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SERVER_ERROR.String() && e.Code == 500
}

func ErrorServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsDataValidateError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_VALIDATE_ERROR.String() && e.Code == 422
}

func ErrorDataValidateError(format string, args ...interface{}) *errors.Error {
	return errors.New(422, ErrorReason_DATA_VALIDATE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsDataSelectError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_SELECT_ERROR.String() && e.Code == 400
}

func ErrorDataSelectError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_SELECT_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsDataAlreadyExists(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_ALREADY_EXISTS.String() && e.Code == 400
}

func ErrorDataAlreadyExists(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_ALREADY_EXISTS.String(), fmt.Sprintf(format, args...))
}

func IsDataNotFound(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_NOT_FOUND.String() && e.Code == 400
}

func ErrorDataNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsNotLoginError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_LOGIN_ERROR.String() && e.Code == 401
}

func ErrorNotLoginError(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_NOT_LOGIN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsDataAddError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_ADD_ERROR.String() && e.Code == 400
}

func ErrorDataAddError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_ADD_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsDataUpdateError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_UPDATE_ERROR.String() && e.Code == 400
}

func ErrorDataUpdateError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_UPDATE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsIdInvalidValueError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ID_INVALID_VALUE_ERROR.String() && e.Code == 400
}

func ErrorIdInvalidValueError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_ID_INVALID_VALUE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsDataDeleteError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_DELETE_ERROR.String() && e.Code == 400
}

func ErrorDataDeleteError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_DELETE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsDataResourceNotFound(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_RESOURCE_NOT_FOUND.String() && e.Code == 400
}

func ErrorDataResourceNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_RESOURCE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsCommandInvalidNotFound(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_COMMAND_INVALID_NOT_FOUND.String() && e.Code == 400
}

func ErrorCommandInvalidNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_COMMAND_INVALID_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsDataUpdateFieldError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_UPDATE_FIELD_ERROR.String() && e.Code == 400
}

func ErrorDataUpdateFieldError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_UPDATE_FIELD_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsGenerateUuidError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GENERATE_UUID_ERROR.String() && e.Code == 400
}

func ErrorGenerateUuidError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_GENERATE_UUID_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsGenerateShortUrlError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GENERATE_SHORT_URL_ERROR.String() && e.Code == 400
}

func ErrorGenerateShortUrlError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_GENERATE_SHORT_URL_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsGetShortUrlError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_SHORT_URL_ERROR.String() && e.Code == 400
}

func ErrorGetShortUrlError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_GET_SHORT_URL_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsUploadFileError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UPLOAD_FILE_ERROR.String() && e.Code == 400
}

func ErrorUploadFileError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_UPLOAD_FILE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsCreateUploadMediaError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CREATE_UPLOAD_MEDIA_ERROR.String() && e.Code == 400
}

func ErrorCreateUploadMediaError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_CREATE_UPLOAD_MEDIA_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsGetFileResourceError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_FILE_RESOURCE_ERROR.String() && e.Code == 400
}

func ErrorGetFileResourceError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_GET_FILE_RESOURCE_ERROR.String(), fmt.Sprintf(format, args...))
}