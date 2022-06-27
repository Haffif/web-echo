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
