package db

import (
	"net/http"
	"proxy/internal/pkg/helpers"
)

const insertRequest = `
INSERT INTO requests(method, uri, proto) VALUES($1, $2, $3) RETURNING id`

const insertHeader = `
INSERT INTO headers(req_id, key, value) VALUES($1, $2, $3)`

func InsertRequest(r *http.Request, uri string) error {
	var id int
	err := connection.QueryRow(insertRequest, r.Method, uri, r.Proto).Scan(&id)
	if err != nil {
		helpers.LogMessage(err.Error())
		return err
	}

	for key, value := range r.Header {
		_, err := Exec(insertHeader, id, key, value[0])
		if err != nil {
			helpers.LogMessage(err.Error())
			return err
		}
	}

	return nil
}

const selectRequestByID = `
SELECT method, uri, proto FROM requests where id = $1`

const selectHeadersByRequestID = `
SELECT key, value FROM headers WHERE req_id = $1`

func SelectRequest(id int) (*http.Request, error) {
	var method, uri, proto string

	err := connection.QueryRow(
		selectRequestByID,
		id,
	).Scan(
		&method,
		&uri,
		&proto,
	)
	if err != nil {
		helpers.LogMessage(err.Error())
		return nil, err
	}

	request, err := http.NewRequest(method, uri, nil)
	if err != nil {
		helpers.LogMessage(err.Error())
		return nil, err
	}

	rows, err := connection.Query(
		selectHeadersByRequestID,
		id,
	)
	defer rows.Close()
	if err != nil {
		helpers.LogMessage(err.Error())
		return nil, err
	}

	for rows.Next() {

		var key, value string
		err := rows.Scan(
			&key,
			&value,
		)

		if err != nil {
			helpers.LogMessage(err.Error())
			return nil, err
		}

		if key != "If-None-Match" && key != "Accept-Encoding" && key != "If-Modified-Since" {
			request.Header.Add(key, value)
		}
	}

	return request, nil
}
