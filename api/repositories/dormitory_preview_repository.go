package repositories

import (
	req "api/types/structs/requests"
	res "api/types/structs/responses"
	"database/sql"
)

func GetDormitoryPreviewsByDormitoryID(dbParam *sql.DB, dormitoryId int) (response []res.DormitoryPreviewResponse, err error) {
	sqlStatement := "SELECT * FROM dormitory_previews WHERE dormitory_id = $1"
	rows, err := dbParam.Query(sqlStatement, dormitoryId)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var dormitoryPreview res.DormitoryPreviewResponse

		err = rows.Scan(&dormitoryPreview.ID, &dormitoryPreview.DormitoryID, &dormitoryPreview.Url, &dormitoryPreview.CreatedAt)
		if err != nil {
			panic(err)
		}

		response = append(response, dormitoryPreview)
	}

	return
}

func CreateDormitoryPreview(dbParam *sql.DB, dormitoryPreviewRequest req.DormitoryPreviewRequest) (response res.DormitoryPreviewResponse, err error) {
	sqlStatement := "INSERT INTO dormitory_previews (dormitory_id, url) VALUES ($1, $2) RETURNING *"
	err = dbParam.QueryRow(sqlStatement, dormitoryPreviewRequest.DomitoryID, dormitoryPreviewRequest.Url).Scan(&response.ID, &response.DormitoryID, &response.Url, &response.CreatedAt)
	return
}

func DeleteDormitoryPreview(dbParam *sql.DB, id int) error {
	sqlStatement := "DELETE FROM dormitory_previews WHERE id = $1"
	result, err := dbParam.Exec(sqlStatement, id)
	row, _ := result.RowsAffected()
	if row == 0 {
		return sql.ErrNoRows
	}
	return err
}