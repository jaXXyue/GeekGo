package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type User struct {
}

func dao() (*User, error) {
	return nil, errors.Wrap(sql.ErrNoRows, "dao query user failed")
}

func biz() (*User, error) {
	user, err := dao()
	if err != nil {
		return nil, errors.WithMessage(err, "biz query user failed")
	}
	return user, nil
}

func service() {
	user, err := biz()
	if errors.Is(err, sql.ErrNoRows) {
		// 404
		fmt.Printf("%+v\n", err)
		return
	} else if err != nil {
		// 500
		fmt.Printf("%+v\n", err)
		return
	}

	// 200
	fmt.Println("found user", user)
}

func main() {
	service()
}
