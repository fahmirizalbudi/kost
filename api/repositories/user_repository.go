package repositories

import (
	// req "api/types/structs/requests"
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