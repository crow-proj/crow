package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/taosdata/driver-go/v2/taosRestful"
)

func main() {
	var taosDSN = "root:taosdata@http(localhost:6041)/demo"
	taos, err := sql.Open("taosRestful", taosDSN)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return
	}

	defer taos.Close()

	rows, err := taos.Query("SELECT * FROM t")
	if err != nil {
		fmt.Println("failed to select from demo, err:", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var ts time.Time
		var speed int
		rows.Scan(&ts, &speed)

		fmt.Println("TS:", ts, "SPEED:", speed)
	}
}
