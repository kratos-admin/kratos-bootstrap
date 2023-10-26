// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: conf/v1/data.proto

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

// Validate checks the field values on Data with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DataMultiError, or nil if none found.
func (m *Data) ValidateAll() error {
	return m.validate(true)
}

func (m *Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDatabase()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Database",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Database",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDatabase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Database",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRedis()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Redis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Redis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRedis()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Redis",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DataMultiError(errors)
	}

	return nil
}

// DataMultiError is an error wrapping multiple validation errors returned by
// Data.ValidateAll() if the designated constraints aren't met.
type DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataMultiError) AllErrors() []error { return m }

// DataValidationError is the validation error returned by Data.Validate if the
// designated constraints aren't met.
type DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataValidationError) ErrorName() string { return "DataValidationError" }

// Error satisfies the builtin error interface
func (e DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataValidationError{}

// Validate checks the field values on Data_Database with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_Database) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Database with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_DatabaseMultiError, or
// nil if none found.
func (m *Data_Database) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Database) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Driver

	// no validation rules for Source

	// no validation rules for Migrate

	// no validation rules for Debug

	// no validation rules for MaxIdleConnections

	// no validation rules for MaxOpenConnections

	if all {
		switch v := interface{}(m.GetConnectionMaxLifetime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_DatabaseValidationError{
					field:  "ConnectionMaxLifetime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_DatabaseValidationError{
					field:  "ConnectionMaxLifetime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConnectionMaxLifetime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_DatabaseValidationError{
				field:  "ConnectionMaxLifetime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Data_DatabaseMultiError(errors)
	}

	return nil
}

// Data_DatabaseMultiError is an error wrapping multiple validation errors
// returned by Data_Database.ValidateAll() if the designated constraints
// aren't met.
type Data_DatabaseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_DatabaseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_DatabaseMultiError) AllErrors() []error { return m }

// Data_DatabaseValidationError is the validation error returned by
// Data_Database.Validate if the designated constraints aren't met.
type Data_DatabaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_DatabaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_DatabaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_DatabaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_DatabaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_DatabaseValidationError) ErrorName() string { return "Data_DatabaseValidationError" }

// Error satisfies the builtin error interface
func (e Data_DatabaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Database.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_DatabaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_DatabaseValidationError{}

// Validate checks the field values on Data_Redis with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_Redis) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Redis with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_RedisMultiError, or
// nil if none found.
func (m *Data_Redis) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Redis) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Network

	// no validation rules for Addr

	// no validation rules for Password

	// no validation rules for Db

	if all {
		switch v := interface{}(m.GetDialTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDialTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_RedisValidationError{
				field:  "DialTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetReadTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "ReadTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "ReadTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReadTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_RedisValidationError{
				field:  "ReadTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWriteTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "WriteTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "WriteTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWriteTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_RedisValidationError{
				field:  "WriteTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Data_RedisMultiError(errors)
	}

	return nil
}

// Data_RedisMultiError is an error wrapping multiple validation errors
// returned by Data_Redis.ValidateAll() if the designated constraints aren't met.
type Data_RedisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_RedisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_RedisMultiError) AllErrors() []error { return m }

// Data_RedisValidationError is the validation error returned by
// Data_Redis.Validate if the designated constraints aren't met.
type Data_RedisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_RedisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_RedisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_RedisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_RedisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_RedisValidationError) ErrorName() string { return "Data_RedisValidationError" }

// Error satisfies the builtin error interface
func (e Data_RedisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Redis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_RedisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_RedisValidationError{}
