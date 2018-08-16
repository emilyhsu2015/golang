package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	e := echo.New()
	e.GET("/show", getWeather)

	e.Logger.Fatal(e.Start(":1323"))
}

func getWeather(c echo.Context) error {
	r := redisdata()
	if r != "" {
		return c.String(http.StatusOK, r)
	}

	// d := getdb()
	// if d != "" {
	// 	return c.String(http.StatusOK, d)
	// }

	a := apidata()
	setRedis(a)
	//setdbdata(a)
	return c.String(http.StatusOK, a)
}

func setRedis(data string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "rweather", data, "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func redisdata() string {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return ""
	}
	defer c.Close()

	rweather, err := redis.String(c.Do("GET", "rweather"))
	if err != nil {
		return ""
		// fmt.Println("redis get failed:", err)
	} else {
		return rweather
		// fmt.Printf("Get mykey: %v \n", rweather)
	}
}
func checkErr(err error) {

	if err != nil {

		panic(err)

	}

}

func getdb() string {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/SQL_Test?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}

	rows, err := db.Query("SELECT * FROM weather2")
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	defer db.Close()
	return ""
}

func setdbdata(data string) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/SQL_Test?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}

	stmt, err := db.Prepare(`INSERT weather2 (data) values (?)`)
	checkErr(err)
	res, err := stmt.Exec(data)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	defer db.Close()
}

func apidata() string {
	url := "http://weather.json.tw/api"

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	var result string

	json.Unmarshal(body, &result)

	// fmt.Printf("%s\n", string(body[:]))

	defer resp.Body.Close()
	return string(body[:])
}
