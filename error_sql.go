package main

import (
	"database/sql"
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

var ErrNotFound = errors.New("Not found!")

func main() {
	err := SomeFunc()
	fmt.Printf("main: %+v\n", err)
}

func SomeFunc() error {
	err := Query("select * from ...")
	if err != nil {
		return err
	}
	return nil
}

func Query(sql string) error {
	err := sqlError()
	if err != nil {
		return xerrors.Wrapf(ErrNotFound,"Query not found")
	}
	return nil
}

func sqlError() error {
	return sql.ErrNoRows
}
