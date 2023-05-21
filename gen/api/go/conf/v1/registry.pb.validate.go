// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: conf/v1/registry.proto

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

// Validate checks the field values on Registry with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Registry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegistryMultiError, or nil
// if none found.
func (m *Registry) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if all {
		switch v := interface{}(m.GetConsul()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Consul",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Consul",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConsul()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Consul",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEtcd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEtcd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Etcd",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetZookeeper()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Zookeeper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Zookeeper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetZookeeper()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Zookeeper",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNacos()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Nacos",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Nacos",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNacos()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Nacos",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetKubernetes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Kubernetes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Kubernetes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKubernetes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Kubernetes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEureka()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Eureka",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Eureka",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEureka()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Eureka",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPolaris()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Polaris",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Polaris",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPolaris()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Polaris",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetServicecomb()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Servicecomb",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Servicecomb",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetServicecomb()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Servicecomb",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegistryMultiError(errors)
	}

	return nil
}

// RegistryMultiError is an error wrapping multiple validation errors returned
// by Registry.ValidateAll() if the designated constraints aren't met.
type RegistryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegistryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegistryMultiError) AllErrors() []error { return m }

// RegistryValidationError is the validation error returned by
// Registry.Validate if the designated constraints aren't met.
type RegistryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistryValidationError) ErrorName() string { return "RegistryValidationError" }

// Error satisfies the builtin error interface
func (e RegistryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistryValidationError{}

// Validate checks the field values on Registry_Consul with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Registry_Consul) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Consul with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Registry_ConsulMultiError, or nil if none found.
func (m *Registry_Consul) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Consul) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Scheme

	// no validation rules for Address

	// no validation rules for HealthCheck

	if len(errors) > 0 {
		return Registry_ConsulMultiError(errors)
	}

	return nil
}

// Registry_ConsulMultiError is an error wrapping multiple validation errors
// returned by Registry_Consul.ValidateAll() if the designated constraints
// aren't met.
type Registry_ConsulMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_ConsulMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_ConsulMultiError) AllErrors() []error { return m }

// Registry_ConsulValidationError is the validation error returned by
// Registry_Consul.Validate if the designated constraints aren't met.
type Registry_ConsulValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_ConsulValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_ConsulValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_ConsulValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_ConsulValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_ConsulValidationError) ErrorName() string { return "Registry_ConsulValidationError" }

// Error satisfies the builtin error interface
func (e Registry_ConsulValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Consul.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_ConsulValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_ConsulValidationError{}

// Validate checks the field values on Registry_Etcd with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Registry_Etcd) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Etcd with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Registry_EtcdMultiError, or
// nil if none found.
func (m *Registry_Etcd) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Etcd) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Registry_EtcdMultiError(errors)
	}

	return nil
}

// Registry_EtcdMultiError is an error wrapping multiple validation errors
// returned by Registry_Etcd.ValidateAll() if the designated constraints
// aren't met.
type Registry_EtcdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_EtcdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_EtcdMultiError) AllErrors() []error { return m }

// Registry_EtcdValidationError is the validation error returned by
// Registry_Etcd.Validate if the designated constraints aren't met.
type Registry_EtcdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_EtcdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_EtcdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_EtcdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_EtcdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_EtcdValidationError) ErrorName() string { return "Registry_EtcdValidationError" }

// Error satisfies the builtin error interface
func (e Registry_EtcdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Etcd.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_EtcdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_EtcdValidationError{}

// Validate checks the field values on Registry_ZooKeeper with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Registry_ZooKeeper) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_ZooKeeper with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Registry_ZooKeeperMultiError, or nil if none found.
func (m *Registry_ZooKeeper) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_ZooKeeper) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Registry_ZooKeeperValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Registry_ZooKeeperValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Registry_ZooKeeperValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Registry_ZooKeeperMultiError(errors)
	}

	return nil
}

// Registry_ZooKeeperMultiError is an error wrapping multiple validation errors
// returned by Registry_ZooKeeper.ValidateAll() if the designated constraints
// aren't met.
type Registry_ZooKeeperMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_ZooKeeperMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_ZooKeeperMultiError) AllErrors() []error { return m }

// Registry_ZooKeeperValidationError is the validation error returned by
// Registry_ZooKeeper.Validate if the designated constraints aren't met.
type Registry_ZooKeeperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_ZooKeeperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_ZooKeeperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_ZooKeeperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_ZooKeeperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_ZooKeeperValidationError) ErrorName() string {
	return "Registry_ZooKeeperValidationError"
}

// Error satisfies the builtin error interface
func (e Registry_ZooKeeperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_ZooKeeper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_ZooKeeperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_ZooKeeperValidationError{}

// Validate checks the field values on Registry_Nacos with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Registry_Nacos) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Nacos with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Registry_NacosMultiError,
// or nil if none found.
func (m *Registry_Nacos) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Nacos) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Port

	// no validation rules for NamespaceId

	// no validation rules for LogLevel

	// no validation rules for CacheDir

	// no validation rules for LogDir

	// no validation rules for UpdateThreadNum

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Registry_NacosValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Registry_NacosValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Registry_NacosValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBeatInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Registry_NacosValidationError{
					field:  "BeatInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Registry_NacosValidationError{
					field:  "BeatInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBeatInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Registry_NacosValidationError{
				field:  "BeatInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for NotLoadCacheAtStart

	// no validation rules for UpdateCacheWhenEmpty

	if len(errors) > 0 {
		return Registry_NacosMultiError(errors)
	}

	return nil
}

// Registry_NacosMultiError is an error wrapping multiple validation errors
// returned by Registry_Nacos.ValidateAll() if the designated constraints
// aren't met.
type Registry_NacosMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_NacosMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_NacosMultiError) AllErrors() []error { return m }

// Registry_NacosValidationError is the validation error returned by
// Registry_Nacos.Validate if the designated constraints aren't met.
type Registry_NacosValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_NacosValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_NacosValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_NacosValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_NacosValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_NacosValidationError) ErrorName() string { return "Registry_NacosValidationError" }

// Error satisfies the builtin error interface
func (e Registry_NacosValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Nacos.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_NacosValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_NacosValidationError{}

// Validate checks the field values on Registry_Kubernetes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Registry_Kubernetes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Kubernetes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Registry_KubernetesMultiError, or nil if none found.
func (m *Registry_Kubernetes) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Kubernetes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Registry_KubernetesMultiError(errors)
	}

	return nil
}

// Registry_KubernetesMultiError is an error wrapping multiple validation
// errors returned by Registry_Kubernetes.ValidateAll() if the designated
// constraints aren't met.
type Registry_KubernetesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_KubernetesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_KubernetesMultiError) AllErrors() []error { return m }

// Registry_KubernetesValidationError is the validation error returned by
// Registry_Kubernetes.Validate if the designated constraints aren't met.
type Registry_KubernetesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_KubernetesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_KubernetesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_KubernetesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_KubernetesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_KubernetesValidationError) ErrorName() string {
	return "Registry_KubernetesValidationError"
}

// Error satisfies the builtin error interface
func (e Registry_KubernetesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Kubernetes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_KubernetesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_KubernetesValidationError{}

// Validate checks the field values on Registry_Eureka with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Registry_Eureka) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Eureka with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Registry_EurekaMultiError, or nil if none found.
func (m *Registry_Eureka) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Eureka) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHeartbeatInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Registry_EurekaValidationError{
					field:  "HeartbeatInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Registry_EurekaValidationError{
					field:  "HeartbeatInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHeartbeatInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Registry_EurekaValidationError{
				field:  "HeartbeatInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRefreshInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Registry_EurekaValidationError{
					field:  "RefreshInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Registry_EurekaValidationError{
					field:  "RefreshInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRefreshInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Registry_EurekaValidationError{
				field:  "RefreshInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Path

	if len(errors) > 0 {
		return Registry_EurekaMultiError(errors)
	}

	return nil
}

// Registry_EurekaMultiError is an error wrapping multiple validation errors
// returned by Registry_Eureka.ValidateAll() if the designated constraints
// aren't met.
type Registry_EurekaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_EurekaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_EurekaMultiError) AllErrors() []error { return m }

// Registry_EurekaValidationError is the validation error returned by
// Registry_Eureka.Validate if the designated constraints aren't met.
type Registry_EurekaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_EurekaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_EurekaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_EurekaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_EurekaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_EurekaValidationError) ErrorName() string { return "Registry_EurekaValidationError" }

// Error satisfies the builtin error interface
func (e Registry_EurekaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Eureka.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_EurekaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_EurekaValidationError{}

// Validate checks the field values on Registry_Polaris with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Registry_Polaris) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Polaris with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Registry_PolarisMultiError, or nil if none found.
func (m *Registry_Polaris) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Polaris) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Port

	// no validation rules for InstanceCount

	// no validation rules for Namespace

	// no validation rules for Service

	// no validation rules for Token

	if len(errors) > 0 {
		return Registry_PolarisMultiError(errors)
	}

	return nil
}

// Registry_PolarisMultiError is an error wrapping multiple validation errors
// returned by Registry_Polaris.ValidateAll() if the designated constraints
// aren't met.
type Registry_PolarisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_PolarisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_PolarisMultiError) AllErrors() []error { return m }

// Registry_PolarisValidationError is the validation error returned by
// Registry_Polaris.Validate if the designated constraints aren't met.
type Registry_PolarisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_PolarisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_PolarisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_PolarisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_PolarisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_PolarisValidationError) ErrorName() string { return "Registry_PolarisValidationError" }

// Error satisfies the builtin error interface
func (e Registry_PolarisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Polaris.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_PolarisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_PolarisValidationError{}

// Validate checks the field values on Registry_Servicecomb with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Registry_Servicecomb) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Servicecomb with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Registry_ServicecombMultiError, or nil if none found.
func (m *Registry_Servicecomb) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Servicecomb) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Registry_ServicecombMultiError(errors)
	}

	return nil
}

// Registry_ServicecombMultiError is an error wrapping multiple validation
// errors returned by Registry_Servicecomb.ValidateAll() if the designated
// constraints aren't met.
type Registry_ServicecombMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_ServicecombMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_ServicecombMultiError) AllErrors() []error { return m }

// Registry_ServicecombValidationError is the validation error returned by
// Registry_Servicecomb.Validate if the designated constraints aren't met.
type Registry_ServicecombValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_ServicecombValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_ServicecombValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_ServicecombValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_ServicecombValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_ServicecombValidationError) ErrorName() string {
	return "Registry_ServicecombValidationError"
}

// Error satisfies the builtin error interface
func (e Registry_ServicecombValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Servicecomb.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_ServicecombValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_ServicecombValidationError{}
