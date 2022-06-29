package models

import (
	"net/http"

	"github.com/Haffif/web-echo/db"
)

type mahasiswa struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Alamat     string `json:"alamat"`
	No_telepon string `json:"no_telepon"`
}

func FetchAllMahasiswa() (Response, error) {
	var obj mahasiswa
	var arrobj []mahasiswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM mahasiswa"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.No_telepon)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreMahasiswa(nama string, alamat string, no_telepon string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT mahasiswa (nama, alamat, no_telepon) VALUES (?,?,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, no_telepon)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertId,
	}

	return res, nil
}

func UpdateMahasiswa(id int, nama string, alamat string, no_telepon string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE mahasiswa SET nama = ?, alamat = ?, no_telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, no_telepon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteMahasiswa(id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "DELETE FROM mahasiswa WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
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
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
