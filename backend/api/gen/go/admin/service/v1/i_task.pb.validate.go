// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin/service/v1/i_task.proto

package servicev1

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

// Validate checks the field values on TaskOption with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TaskOption) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TaskOption with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TaskOptionMultiError, or
// nil if none found.
func (m *TaskOption) ValidateAll() error {
	return m.validate(true)
}

func (m *TaskOption) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.RetryCount != nil {
		// no validation rules for RetryCount
	}

	if m.Timeout != nil {

		if all {
			switch v := interface{}(m.GetTimeout()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TaskOptionValidationError{
						field:  "Timeout",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TaskOptionValidationError{
						field:  "Timeout",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskOptionValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Deadline != nil {

		if all {
			switch v := interface{}(m.GetDeadline()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TaskOptionValidationError{
						field:  "Deadline",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TaskOptionValidationError{
						field:  "Deadline",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeadline()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskOptionValidationError{
					field:  "Deadline",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ProcessIn != nil {

		if all {
			switch v := interface{}(m.GetProcessIn()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TaskOptionValidationError{
						field:  "ProcessIn",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TaskOptionValidationError{
						field:  "ProcessIn",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetProcessIn()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskOptionValidationError{
					field:  "ProcessIn",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ProcessAt != nil {

		if all {
			switch v := interface{}(m.GetProcessAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TaskOptionValidationError{
						field:  "ProcessAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TaskOptionValidationError{
						field:  "ProcessAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetProcessAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskOptionValidationError{
					field:  "ProcessAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TaskOptionMultiError(errors)
	}

	return nil
}

// TaskOptionMultiError is an error wrapping multiple validation errors
// returned by TaskOption.ValidateAll() if the designated constraints aren't met.
type TaskOptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TaskOptionMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TaskOptionMultiError) AllErrors() []error { return m }

// TaskOptionValidationError is the validation error returned by
// TaskOption.Validate if the designated constraints aren't met.
type TaskOptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskOptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskOptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskOptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskOptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskOptionValidationError) ErrorName() string { return "TaskOptionValidationError" }

// Error satisfies the builtin error interface
func (e TaskOptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskOptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskOptionValidationError{}

// Validate checks the field values on Task with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Task) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Task with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TaskMultiError, or nil if none found.
func (m *Task) ValidateAll() error {
	return m.validate(true)
}

func (m *Task) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.TypeName != nil {
		// no validation rules for TypeName
	}

	if m.TaskPayload != nil {
		// no validation rules for TaskPayload
	}

	if m.CronSpec != nil {
		// no validation rules for CronSpec
	}

	if m.TaskOptions != nil {

		if all {
			switch v := interface{}(m.GetTaskOptions()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TaskValidationError{
						field:  "TaskOptions",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TaskValidationError{
						field:  "TaskOptions",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTaskOptions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskValidationError{
					field:  "TaskOptions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Enable != nil {
		// no validation rules for Enable
	}

	if m.Remark != nil {
		// no validation rules for Remark
	}

	if m.CreateBy != nil {
		// no validation rules for CreateBy
	}

	if m.UpdateBy != nil {
		// no validation rules for UpdateBy
	}

	if m.CreateTime != nil {

		if all {
			switch v := interface{}(m.GetCreateTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TaskValidationError{
						field:  "CreateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TaskValidationError{
						field:  "CreateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.UpdateTime != nil {

		if all {
			switch v := interface{}(m.GetUpdateTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TaskValidationError{
						field:  "UpdateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TaskValidationError{
						field:  "UpdateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.DeleteTime != nil {

		if all {
			switch v := interface{}(m.GetDeleteTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TaskValidationError{
						field:  "DeleteTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TaskValidationError{
						field:  "DeleteTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeleteTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskValidationError{
					field:  "DeleteTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TaskMultiError(errors)
	}

	return nil
}

// TaskMultiError is an error wrapping multiple validation errors returned by
// Task.ValidateAll() if the designated constraints aren't met.
type TaskMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TaskMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TaskMultiError) AllErrors() []error { return m }

// TaskValidationError is the validation error returned by Task.Validate if the
// designated constraints aren't met.
type TaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskValidationError) ErrorName() string { return "TaskValidationError" }

// Error satisfies the builtin error interface
func (e TaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskValidationError{}

// Validate checks the field values on ListTaskResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTaskResponseMultiError, or nil if none found.
func (m *ListTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListTaskResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTaskResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTaskResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListTaskResponseMultiError(errors)
	}

	return nil
}

// ListTaskResponseMultiError is an error wrapping multiple validation errors
// returned by ListTaskResponse.ValidateAll() if the designated constraints
// aren't met.
type ListTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTaskResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTaskResponseMultiError) AllErrors() []error { return m }

// ListTaskResponseValidationError is the validation error returned by
// ListTaskResponse.Validate if the designated constraints aren't met.
type ListTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTaskResponseValidationError) ErrorName() string { return "ListTaskResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTaskResponseValidationError{}

// Validate checks the field values on GetTaskRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTaskRequestMultiError,
// or nil if none found.
func (m *GetTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetTaskRequestMultiError(errors)
	}

	return nil
}

// GetTaskRequestMultiError is an error wrapping multiple validation errors
// returned by GetTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskRequestMultiError) AllErrors() []error { return m }

// GetTaskRequestValidationError is the validation error returned by
// GetTaskRequest.Validate if the designated constraints aren't met.
type GetTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskRequestValidationError) ErrorName() string { return "GetTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskRequestValidationError{}

// Validate checks the field values on GetTaskByTypeNameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTaskByTypeNameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskByTypeNameRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTaskByTypeNameRequestMultiError, or nil if none found.
func (m *GetTaskByTypeNameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskByTypeNameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TypeName

	if len(errors) > 0 {
		return GetTaskByTypeNameRequestMultiError(errors)
	}

	return nil
}

// GetTaskByTypeNameRequestMultiError is an error wrapping multiple validation
// errors returned by GetTaskByTypeNameRequest.ValidateAll() if the designated
// constraints aren't met.
type GetTaskByTypeNameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskByTypeNameRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskByTypeNameRequestMultiError) AllErrors() []error { return m }

// GetTaskByTypeNameRequestValidationError is the validation error returned by
// GetTaskByTypeNameRequest.Validate if the designated constraints aren't met.
type GetTaskByTypeNameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskByTypeNameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskByTypeNameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskByTypeNameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskByTypeNameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskByTypeNameRequestValidationError) ErrorName() string {
	return "GetTaskByTypeNameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTaskByTypeNameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskByTypeNameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskByTypeNameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskByTypeNameRequestValidationError{}

// Validate checks the field values on CreateTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaskRequestMultiError, or nil if none found.
func (m *CreateTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTaskRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTaskRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTaskRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTaskRequestMultiError(errors)
	}

	return nil
}

// CreateTaskRequestMultiError is an error wrapping multiple validation errors
// returned by CreateTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaskRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaskRequestMultiError) AllErrors() []error { return m }

// CreateTaskRequestValidationError is the validation error returned by
// CreateTaskRequest.Validate if the designated constraints aren't met.
type CreateTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskRequestValidationError) ErrorName() string {
	return "CreateTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskRequestValidationError{}

// Validate checks the field values on UpdateTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTaskRequestMultiError, or nil if none found.
func (m *UpdateTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateTaskRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateTaskRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateTaskRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateTaskRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateTaskRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateTaskRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.AllowMissing != nil {
		// no validation rules for AllowMissing
	}

	if len(errors) > 0 {
		return UpdateTaskRequestMultiError(errors)
	}

	return nil
}

// UpdateTaskRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTaskRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTaskRequestMultiError) AllErrors() []error { return m }

// UpdateTaskRequestValidationError is the validation error returned by
// UpdateTaskRequest.Validate if the designated constraints aren't met.
type UpdateTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTaskRequestValidationError) ErrorName() string {
	return "UpdateTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTaskRequestValidationError{}

// Validate checks the field values on DeleteTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTaskRequestMultiError, or nil if none found.
func (m *DeleteTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteTaskRequestMultiError(errors)
	}

	return nil
}

// DeleteTaskRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTaskRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTaskRequestMultiError) AllErrors() []error { return m }

// DeleteTaskRequestValidationError is the validation error returned by
// DeleteTaskRequest.Validate if the designated constraints aren't met.
type DeleteTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskRequestValidationError) ErrorName() string {
	return "DeleteTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskRequestValidationError{}

// Validate checks the field values on RestartAllTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RestartAllTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RestartAllTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RestartAllTaskResponseMultiError, or nil if none found.
func (m *RestartAllTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RestartAllTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	if len(errors) > 0 {
		return RestartAllTaskResponseMultiError(errors)
	}

	return nil
}

// RestartAllTaskResponseMultiError is an error wrapping multiple validation
// errors returned by RestartAllTaskResponse.ValidateAll() if the designated
// constraints aren't met.
type RestartAllTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RestartAllTaskResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RestartAllTaskResponseMultiError) AllErrors() []error { return m }

// RestartAllTaskResponseValidationError is the validation error returned by
// RestartAllTaskResponse.Validate if the designated constraints aren't met.
type RestartAllTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RestartAllTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RestartAllTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RestartAllTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RestartAllTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RestartAllTaskResponseValidationError) ErrorName() string {
	return "RestartAllTaskResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RestartAllTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRestartAllTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RestartAllTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RestartAllTaskResponseValidationError{}

// Validate checks the field values on ControlTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ControlTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ControlTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ControlTaskRequestMultiError, or nil if none found.
func (m *ControlTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ControlTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ControlType

	// no validation rules for TypeName

	if len(errors) > 0 {
		return ControlTaskRequestMultiError(errors)
	}

	return nil
}

// ControlTaskRequestMultiError is an error wrapping multiple validation errors
// returned by ControlTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type ControlTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ControlTaskRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ControlTaskRequestMultiError) AllErrors() []error { return m }

// ControlTaskRequestValidationError is the validation error returned by
// ControlTaskRequest.Validate if the designated constraints aren't met.
type ControlTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ControlTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ControlTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ControlTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ControlTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ControlTaskRequestValidationError) ErrorName() string {
	return "ControlTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ControlTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sControlTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ControlTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ControlTaskRequestValidationError{}
