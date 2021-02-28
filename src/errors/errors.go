// Package errors provides some functions for handling errors.
package errors
import (
  "log"
)

// CheckFatal checks if err exists and logs it as a fatal error if it does.
func CheckFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
