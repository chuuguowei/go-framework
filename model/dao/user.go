package dao

import (
	"context"
	"go-framework/library/client"
	"go-framework/library/mysql"
)

type User struct {
	Uid      int64  `db:"uid"`
	Username string `db:"username"`
	Age      int64  `db:"age"`
	Ctime    int64  `db:"create_time"`
}

type UserModel struct {
	Model
}

func NewUserModel(ctx context.Context) UserModel {
	return UserModel{
		Model: NewModel(ctx, client.MySqlClient),
	}
}

func (that *UserModel) FindOne(pk uint) (*User, error) {
	var user User
	err := that.Query(&user, "select * from users where uid=?", pk)
	switch err {
	case nil:
		return &user, nil
	case mysql.ErrNotFound:
		return nil, mysql.ErrNotFound
	default:
		return nil, err
	}
}

func (that *UserModel) List(age int) ([]User, error) {
	var users []User
	err := that.Query(&users, "select age,username,create_time from users where age>?", age)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (that *UserModel) Insert(user *User) (int64, error) {
	const sql = `insert into users (username,age) values(?, ?)`
	// insert op
	res, err := that.Exec(sql, user.Username, user.Age)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (that *UserModel) Delete(pk uint64) error {
	const sql = `delete from users where uid=?`
	// insert op
	_, err := that.Exec(sql, pk)
	return err
}
