// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: transaction/v1/transaction.proto

package transactionv1

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

// Validate checks the field values on Transaction with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Transaction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Transaction with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TransactionMultiError, or
// nil if none found.
func (m *Transaction) ValidateAll() error {
	return m.validate(true)
}

func (m *Transaction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return TransactionMultiError(errors)
	}

	return nil
}

// TransactionMultiError is an error wrapping multiple validation errors
// returned by Transaction.ValidateAll() if the designated constraints aren't met.
type TransactionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransactionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransactionMultiError) AllErrors() []error { return m }

// TransactionValidationError is the validation error returned by
// Transaction.Validate if the designated constraints aren't met.
type TransactionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionValidationError) ErrorName() string { return "TransactionValidationError" }

// Error satisfies the builtin error interface
func (e TransactionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransaction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionValidationError{}

// Validate checks the field values on SubmitRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubmitRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubmitRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubmitRequestMultiError, or
// nil if none found.
func (m *SubmitRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SubmitRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for From

	// no validation rules for Contract

	if all {
		switch v := interface{}(m.GetTransaction()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubmitRequestValidationError{
					field:  "Transaction",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubmitRequestValidationError{
					field:  "Transaction",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTransaction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubmitRequestValidationError{
				field:  "Transaction",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SubmitRequestMultiError(errors)
	}

	return nil
}

// SubmitRequestMultiError is an error wrapping multiple validation errors
// returned by SubmitRequest.ValidateAll() if the designated constraints
// aren't met.
type SubmitRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubmitRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubmitRequestMultiError) AllErrors() []error { return m }

// SubmitRequestValidationError is the validation error returned by
// SubmitRequest.Validate if the designated constraints aren't met.
type SubmitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitRequestValidationError) ErrorName() string { return "SubmitRequestValidationError" }

// Error satisfies the builtin error interface
func (e SubmitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitRequestValidationError{}

// Validate checks the field values on SubmitResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubmitResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubmitResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubmitResponseMultiError,
// or nil if none found.
func (m *SubmitResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SubmitResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Hash != nil {
		// no validation rules for Hash
	}

	if m.RejectReason != nil {
		// no validation rules for RejectReason
	}

	if len(errors) > 0 {
		return SubmitResponseMultiError(errors)
	}

	return nil
}

// SubmitResponseMultiError is an error wrapping multiple validation errors
// returned by SubmitResponse.ValidateAll() if the designated constraints
// aren't met.
type SubmitResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubmitResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubmitResponseMultiError) AllErrors() []error { return m }

// SubmitResponseValidationError is the validation error returned by
// SubmitResponse.Validate if the designated constraints aren't met.
type SubmitResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitResponseValidationError) ErrorName() string { return "SubmitResponseValidationError" }

// Error satisfies the builtin error interface
func (e SubmitResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitResponseValidationError{}
