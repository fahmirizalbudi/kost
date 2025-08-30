package repositories

import (
	req "api/types/structs/requests"
	res "api/types/structs/responses"
	"database/sql"
)

func GetAllDormitories(dbParam *sql.DB) (response []res.DormitoryResponse, err error) {
	sqlStatement := "SELECT * FROM dormitories"
	rows, err := dbParam.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var dormitory res.DormitoryResponse

		err = rows.Scan(&dormitory.ID, &dormitory.Name, &dormitory.Address, &dormitory.Description, &dormitory.Price, &dormitory.Facilities, &dormitory.GoogleMaps, &dormitory.CreatedAt, &dormitory.UpdatedAt)
		if err != nil {
			panic(err)
		}

		response = append(response, dormitory)
	}

	return
}

func CreateDormitory(dbParam *sql.DB, dormitoryRequest req.DormitoryRequest) (response res.DormitoryResponse, err error) {
	sqlStatement := "INSERT INTO dormitories (name, address, description, price, facilities, google_maps) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	err = dbParam.QueryRow(sqlStatement, dormitoryRequest.Name, dormitoryRequest.Address, dormitoryRequest.Description, *dormitoryRequest.Price, dormitoryRequest.Facilities, dormitoryRequest.GoogleMaps).Scan(&response.ID, &response.Name, &response.Address, &response.Description, &response.Price, &response.Facilities, &response.GoogleMaps, &response.CreatedAt, &response.UpdatedAt)
	return
}

func GetDormitoryByID(dbParam *sql.DB, id int) (response res.DormitoryResponse, err error) {
	sqlStatement := "SELECT * FROM dormitories WHERE id = $1"
	err = dbParam.QueryRow(sqlStatement, id).Scan(&response.ID, &response.Name, &response.Address, &response.Description, &response.Price, &response.Facilities, &response.GoogleMaps, &response.CreatedAt, &response.UpdatedAt)
	return
}

func UpdateDormitory(dbParam *sql.DB, id int, dormitoryRequest req.DormitoryRequest) (response res.DormitoryResponse, err error) {
	sqlStatement := "UPDATE dormitories SET name = $1, address = $2, description = $3, price = $4, facilities = $5, google_maps = $6, updated_at = NOW() WHERE id = $7 RETURNING *"
	err = dbParam.QueryRow(sqlStatement, dormitoryRequest.Name, dormitoryRequest.Address, dormitoryRequest.Description, *dormitoryRequest.Price, dormitoryRequest.Facilities, dormitoryRequest.GoogleMaps, id).Scan(&response.ID, &response.Name, &response.Address, &response.Description, &response.Price, &response.Facilities, &response.GoogleMaps, &response.CreatedAt, &response.UpdatedAt)
	return
}

func DeleteDormitory(dbParam *sql.DB, id int) error {
	sqlStatement := "DELETE FROM dormitories WHERE id = $1"
	result, err := dbParam.Exec(sqlStatement, id)
	row, _ := result.RowsAffected()
	if row == 0 {
		return sql.ErrNoRows
	}
	return err
}

func GetAllDormitoriesWithPreviews(dbParam *sql.DB) (response []res.DormitoryWithPreviewResponse, err error) {
	sqlStatement := "SELECT * FROM dormitories"
	rows, err := dbParam.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var dormitory res.DormitoryWithPreviewResponse

		err = rows.Scan(&dormitory.ID, &dormitory.Name, &dormitory.Address, &dormitory.Description, &dormitory.Price, &dormitory.Facilities, &dormitory.GoogleMaps, &dormitory.CreatedAt, &dormitory.UpdatedAt)
		if err != nil {
			panic(err)
		}

		dormitoryPreviews, err := GetDormitoryPreviewsByDormitoryID(dbParam, dormitory.ID)
		if err != nil {
			panic(err)
		}

		dormitory.Previews = dormitoryPreviews

		response = append(response, dormitory)
	}

	return
}