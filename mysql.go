package main

import (
	"fmt"
	_ "github.com/aloxc/goice/mysql"
	"github.com/aloxc/goice/sql"
	"sync"
	"time"
)

type Area struct {
	Id   int
	Name string
}

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:33066)/lovehome")
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var count = 50
	wait := sync.WaitGroup{}
	wait.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {

			rows, err := db.Query("select * from area limit 1")

			var area Area
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				//row.Scan(...)
				rows.Scan(&area.Id, &area.Name)
				fmt.Println("1:", i, area, time.Now().Unix())
			}
			//time.Sleep(time.Second * 20)
			rows.Close()
			rows, err = db.Query("select * from area limit 1")
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				//row.Scan(...)
				rows.Scan(&area.Id, &area.Name)
				fmt.Println("2:", i, area, time.Now().Unix())
				//fmt.Println(area)
			}
			rows.Close()
			//fmt.Println("22:",i,time.Now().Unix())
			//time.Sleep(time.Second * 20)
			wait.Done()
		}(i)
	}
	wait.Wait()
}
