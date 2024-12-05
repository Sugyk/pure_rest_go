package databases

import (
	"database/sql"
	"errors"
	"fmt"
)

func HandleError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorObjNotFound{}
	}
	return err
}

type ErrorObjNotFound struct{}

func (ErrorObjNotFound) Error() string {
	return "object not found"
}

func (ErrorObjNotFound) Unwrap() error {
	return fmt.Errorf("object not found")
}
