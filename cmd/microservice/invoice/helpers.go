package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxByte := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))
	//TODO:
	//FIXME:
	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, header ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if len(header) > 0 {
		for k, v := range header[0] {
			w.Header()[k] = v
		}
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {

}
func (app *application) CreateDirIfNotExists(path string) error {
	const mode = 0755

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}
	return nil

}
