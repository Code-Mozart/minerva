package errors

import "fmt"

type MissingEnvironmentVariableError struct {
	Key string
}

func (err *MissingEnvironmentVariableError) Error() string {
	return fmt.Sprintf("Missing environment variable '%s'", err.Key)
}
