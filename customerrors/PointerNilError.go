package customerrors

import (
	"errors"
	"fmt"
)

// PointerNilError
//
// is a custom error type for nil pointer errors.
//
// Typically used when a function is called with a
// nil pointer as an argument.
type PointerNilError struct {
	functionCallChain string
	sourceFile        string
	parentType        string
	thisFunction      string
	nilParameterName  string
	errContext        string
	message           string
	wrappedErr        error
}

/*
Example Code Usage
==================

package main

import (
	"fmt"
	"yourproject/customerrors" // Assuming your custom error package name
)

type Database struct{}

func ProcessUser(db *Database, userID string) error {
	if db == nil {
		// Use the constructor pattern safely
		err, initErr := errors.NewPointerNilError(
			"main() -> ProcessUser()",
			"user_service.go",
			"Database",
			"ProcessUser",
			"db",
			"Fetching active user record",
			"", // Leaving 'errorMessage' blank to leverage the default fallback message
		)
		if initErr != nil {
			return initErr // Handles constructor configuration failure
		}
		return err
	}

	return nil
}
*/

// NewPointerNilError
//
// This is the constructor function that enforces validation,
// handles encapsulation, and eagerly prepares the error state.
func NewPointerNilError(
	functionCallChain string,
	sourceFile string,
	parentType string,
	thisFunction string,
	nilParameterName string,
	errContext string,
	errorMessage string,
) (*PointerNilError, error) {

	// Enforce strict required validations immediately
	if nilParameterName == "" {
		return nil, fmt.Errorf("PointerNilError initialization failed: nilParameterName is required")
	}

	// Default fallback message if empty
	if errorMessage == "" {
		errorMessage = fmt.Sprintf("%s is a nil pointer!", nilParameterName)
	}

	// Construct the concrete struct pointer
	err := &PointerNilError{
		functionCallChain: functionCallChain,
		sourceFile:        sourceFile,
		parentType:        parentType,
		thisFunction:      thisFunction,
		nilParameterName:  nilParameterName,
		errContext:        errContext,
		message:           errorMessage,
	}

	// Eagerly compute the string error block and wrap it.
	// This makes Error() and Unwrap() completely read-only and thread-safe.
	err.wrappedErr = errors.New(err.buildErrorString())

	return err, nil
}

// buildErrorString acts as an internal helper to format the output block
func (e *PointerNilError) buildErrorString() string {
	errString := "\n"

	if e.functionCallChain != "" {
		errString += fmt.Sprintf("Function Call Chain: %s\n", e.functionCallChain)
	}
	if e.sourceFile != "" {
		errString += fmt.Sprintf("Source File: %s\n", e.sourceFile)
	}
	if e.parentType != "" {
		errString += fmt.Sprintf("Parent Type: %s\n", e.parentType)
	}
	if e.thisFunction != "" {
		errString += fmt.Sprintf("Calling Function: %s\n", e.thisFunction)
	}

	errString += fmt.Sprintf("Nil Parameter Name: %s\n", e.nilParameterName)

	if e.errContext != "" {
		errString += fmt.Sprintf("Error Context: %s\n", e.errContext)
	}
	if e.message != "" {
		errString += fmt.Sprintf("%s\n", e.message)
	}

	return errString
}

// Error satisfies Go's built-in error interface.
// Because e.wrappedErr was initialized eagerly, this is completely thread-safe.
func (e *PointerNilError) Error() string {
	if e.wrappedErr == nil {
		return ""
	}
	return e.wrappedErr.Error()
}

// Unwrap satisfies Go's standard unwrap interface (`errors.Unwrap`).
// Safe to call at any time, in any order relative to Error().
func (e *PointerNilError) Unwrap() error {
	return e.wrappedErr
}
