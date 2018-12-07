package repository

import (
	"github.com/mkadiri/golang-microservice/database"
	"github.com/mkadiri/golang-microservice/model"
)

func GetUsers() (users []model.User, err error) {
	rows, err := database.Db.Query("select * from `user`")

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateOfBirth)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, err
}

func AddUser(user model.User) (savedUser model.User, err error) {
	stmt, err := database.Db.Prepare("insert into `user` (first_name, last_name, date_of_birth) values(?, ?, ?)")
	defer stmt.Close()
	
	if err != nil {
		return savedUser, err
	}
	
	res, err := stmt.Exec(user.FirstName, user.LastName, user.DateOfBirth)

	if err != nil {
		return savedUser, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return savedUser, err
	}

	savedUser, err = GetUserById(int(id))
	
	return savedUser, err
}

func GetUserById(id int) (user model.User, err error) {
	err = database.Db.QueryRow("select * from `user` where `id` = ?", id).
		Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateOfBirth)

	return user, err
}