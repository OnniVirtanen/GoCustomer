package auth

import "database/sql"

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type UserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Id int `json:"id"`
}

func ScanRow(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(&user.Id, &user.Email, &user.Hash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
