// Package errors provides some functions for handling errors.
package errors

import (
	"log"
	"os"
)

// CheckLogFatal checks if err exists and logs it as a fatal error if it does.
func CheckLogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CheckExitFatal exits if err exists.
func CheckExitFatal(err error) {
	if err != nil {
		os.Exit(1)
	}
}
