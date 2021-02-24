package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// PrivatePOST -
func PrivatePOST(r io.Reader, p string) error {
	var err error
	var plst M3U
	plst.Loads()
	defer plst.Dumps()

	switch p {
	case "toggle":
		var m struct {
			ID   string
			Hide bool
		}
		err = json.NewDecoder(r).Decode(&m)
		if err != nil {
			return err
		}
		for _, v := range plst {
			if v.ID == m.ID {
				v.Hide = m.Hide
				break
			}
		}
	case "save":
		var m []struct {
			ID    string
			Order int
		}
		err = json.NewDecoder(r).Decode(&m)
		if err != nil {
			return err
		}
		for _, v := range plst {
			for _, i := range m {
				if v.ID == i.ID {
					v.Order = i.Order
				}
			}
		}
	default:
		return errors.New("Method Not Allowed")
	}
	return nil
}

// PrivateXML -
func PrivateXML(w http.ResponseWriter) error {
	return transfer("ott.tv.planeta.tc/epg/program.xml.gz", "", w)
}

// PrivateGetJSON -
func PrivateGetJSON(w http.ResponseWriter) error {
	var plst M3U
	plst.Loads()
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(&plst)
}
