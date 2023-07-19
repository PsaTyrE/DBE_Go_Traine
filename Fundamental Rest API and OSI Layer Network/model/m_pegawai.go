package model

import (
	"myApp/db"
	"net/http"
)

type Pegawai struct {
	Id     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Email  string `json:"email"`
	No_hp  string `json:"no_hp"`
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	conn := db.CreateCon()

	query := "SELECT * FROM pegawai"

	rows, err := conn.Query(query)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Email, &obj.No_hp)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func StorePegawai(nama string, alamat string, email string, no_hp string) (Response, error) {
	var res Response

	conn := db.CreateCon()

	query := "INSERT pegawai (nama, alamat, email, no_hp) VALUES (?, ?, ?, ?)"

	stmt, err := conn.Prepare(query)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, email, no_hp)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdatePegawai(id int, nama string, alamat string, email string, no_hp string) (Response, error) {
	var res Response

	conn := db.CreateCon()

	query := "UPDATE pegawai SET nama=? , alamat=?, email=?, no_hp=? Where id=?"

	stmt, err := conn.Prepare(query)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, email, no_hp, id)

	if err != nil {
		return res, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success update data!"
	res.Data = map[string]int64{
		"rows_affected ": rowAffected,
	}

	return res, nil

}

func DeletePegawai(id int) (Response, error) {
	var res Response

	conn := db.CreateCon()

	query := "DELETE FROM pegawai WHERE id=?"
	stmt, err := conn.Prepare(query)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}
