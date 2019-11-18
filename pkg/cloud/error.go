package cloud

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"io"
	"net"
	"net/url"
	"os"
	"syscall"
)

func IsAwsError(err error, awsCode string) bool {
	if err == nil {
		return false
	}

	aerr, ok := err.(awserr.Error)

	return ok && aerr.Code() == awsCode
}

func IsRequestCanceled(err error) bool {
	return IsAwsError(err, request.CanceledErrorCode)
}

func IsConnectionError(err error) bool {
	if IsUrlError(err, io.EOF) {
		return true
	}

	if IsSyscallError(err, syscall.ECONNREFUSED, syscall.ECONNRESET, syscall.EPIPE) {
		return true
	}

	return false
}

func IsUrlError(err error, targets ...error) bool {
	if err == nil {
		return false
	}

	aerr, ok := err.(awserr.Error)

	if !ok {
		return false
	}

	urlErr, ok := aerr.OrigErr().(*url.Error)

	if !ok {
		return false
	}

	for _, t := range targets {
		if errors.Is(urlErr, t) {
			return true
		}
	}

	return false
}

func IsSyscallError(err error, syscallErrors ...syscall.Errno) bool {
	if err == nil {
		return false
	}

	aerr, ok := err.(awserr.Error)

	if !ok {
		return false
	}

	urlErr, ok := aerr.OrigErr().(*url.Error)

	if !ok {
		return false
	}

	opErr, ok := urlErr.Err.(*net.OpError)

	if !ok {
		return false
	}

	for {
		if nextOpErr, ok := opErr.Err.(*net.OpError); ok {
			opErr = nextOpErr
		} else {
			break
		}
	}

	syscallErr, ok := opErr.Err.(*os.SyscallError)

	if !ok {
		return false
	}

	for _, sysErr := range syscallErrors {
		if syscallErr.Err == sysErr {
			return true
		}
	}

	return false
}

const RequestCanceledError = requestCanceledError("RequestCanceled")

type requestCanceledError string

func (e requestCanceledError) Error() string {
	return string(e)
}
