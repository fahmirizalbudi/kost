package repositories

import (
	req "api/types/structs/requests"
	res "api/types/structs/responses"
	"database/sql"
)

func GetAllUsers(dbParam *sql.DB) (result []res.UserResponse, err error) {
	sql := "SELECT id, name, email, role, phone, address, created_at, updated_at FROM users"

	rows, err := dbParam.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var user res.UserResponse

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.Phone, &user.Address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}
		
		result = append(result, user)
	}
	return
}

func CreateUser(dbParam *sql.DB, userRequest req.UserRequest) (response res.UserResponse, err error) {
	sql := "INSERT INTO users (name, email, password, role, phone, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, email, role, phone, address, created_at, updated_at"
	err = dbParam.QueryRow(sql, userRequest.Name, userRequest.Email, userRequest.Password, userRequest.Role, userRequest.Phone, userRequest.Address).Scan(&response.ID, &response.Name, &response.Email, &response.Role, &response.Phone, &response.Address, &response.CreatedAt, &response.UpdatedAt)
	return
}