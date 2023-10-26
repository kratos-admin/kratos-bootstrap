// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: conf/v1/oss.proto

package conf

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on OSS with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *OSS) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OSS with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OSSMultiError, or nil if none found.
func (m *OSS) ValidateAll() error {
	return m.validate(true)
}

func (m *OSS) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMinio()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OSSValidationError{
					field:  "Minio",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OSSValidationError{
					field:  "Minio",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMinio()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OSSValidationError{
				field:  "Minio",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OSSMultiError(errors)
	}

	return nil
}

// OSSMultiError is an error wrapping multiple validation errors returned by
// OSS.ValidateAll() if the designated constraints aren't met.
type OSSMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OSSMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OSSMultiError) AllErrors() []error { return m }

// OSSValidationError is the validation error returned by OSS.Validate if the
// designated constraints aren't met.
type OSSValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OSSValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OSSValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OSSValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OSSValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OSSValidationError) ErrorName() string { return "OSSValidationError" }

// Error satisfies the builtin error interface
func (e OSSValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOSS.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OSSValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OSSValidationError{}

// Validate checks the field values on OSS_MinIO with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OSS_MinIO) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OSS_MinIO with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OSS_MinIOMultiError, or nil
// if none found.
func (m *OSS_MinIO) ValidateAll() error {
	return m.validate(true)
}

func (m *OSS_MinIO) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Endpoint

	// no validation rules for AccessKey

	// no validation rules for SecretKey

	// no validation rules for Token

	// no validation rules for UseSsl

	// no validation rules for UploadHost

	// no validation rules for DownloadHost

	if len(errors) > 0 {
		return OSS_MinIOMultiError(errors)
	}

	return nil
}

// OSS_MinIOMultiError is an error wrapping multiple validation errors returned
// by OSS_MinIO.ValidateAll() if the designated constraints aren't met.
type OSS_MinIOMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OSS_MinIOMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OSS_MinIOMultiError) AllErrors() []error { return m }

// OSS_MinIOValidationError is the validation error returned by
// OSS_MinIO.Validate if the designated constraints aren't met.
type OSS_MinIOValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OSS_MinIOValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OSS_MinIOValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OSS_MinIOValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OSS_MinIOValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OSS_MinIOValidationError) ErrorName() string { return "OSS_MinIOValidationError" }

// Error satisfies the builtin error interface
func (e OSS_MinIOValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOSS_MinIO.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OSS_MinIOValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OSS_MinIOValidationError{}