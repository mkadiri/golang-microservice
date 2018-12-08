package repository

import (
	"github.com/mkadiri/golang-microservice/database"
	"github.com/mkadiri/golang-microservice/model"
)

func AddUser(user model.User) (savedUser model.User, err error) {
	stmt, err := database.Db.Prepare("insert into `user` (first_name, last_name, date_of_birth) values(?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		return savedUser, err
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.DateOfBirth)

	if err != nil {
		return savedUser, err
	}

	return user, err
}