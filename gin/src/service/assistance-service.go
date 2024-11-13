package service

import (
	"database/sql"
	"gin-project/src/model"
	"gin-project/src/response"
)

func GetAssistanceTools(id string, db *sql.DB) (model.Bantuan, error) {
	var bantuan model.Bantuan

	query := `SELECT 
          assistance.id, assistance.nama, assistance.koordinator, 
          assistance.sumber_anggaran, assistance.total_anggaran, 
          assistance.tahun_pemberian, 
          assistance_tools.kuantitas, 
          tools.id, tools.nama_item, tools.harga, tools.deskripsi
        FROM assistance
        LEFT JOIN assistance_tools ON assistance.id = assistance_tools.assistance_id
        LEFT JOIN tools ON assistance_tools.tools_id = tools.id
        WHERE assistance.id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return bantuan, response.NewResponseError(400, err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return bantuan, response.NewResponseError(400, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var alat model.Alat
		if err := rows.Scan(
			&bantuan.Id,
			&bantuan.Name,
			&bantuan.Koordinator,
			&bantuan.Sumber_anggaran,
			&bantuan.Total_anggaran,
			&bantuan.Tahun_pemberian,
			&alat.Kuantitas,
			&alat.Id,
			&alat.Name,
			&alat.Harga,
			&alat.Deskripsi,
		); err != nil {
			return bantuan, response.NewResponseError(400, err.Error())
		}

		if alat.Id.Valid {
			bantuan.Alat = append(bantuan.Alat, alat)
		}
	}

	return bantuan, nil
}

func CreateAssistanceTools(request model.CreateAssistanceToolsRequest, db *sql.DB) error {
	query := `INSERT INTO assistance_tools (assistance_id, tools_id, kuantitas) VALUES (?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return response.NewResponseError(500, err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.Assistance_id, request.Tools_id, request.Kuantitas)
	if err != nil {
		return response.NewResponseError(400, err.Error())
	}
	return nil
}
