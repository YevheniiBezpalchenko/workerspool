package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"workerspool/db"
	"workerspool/models"
	repo "workerspool/repository"
	"workerspool/workerpool"
)

func main() {

	wc := func() worker_pool.Worker {
		return TestWorker{}
	}
	wpool := worker_pool.NewPool(4, wc)
	wpool.DataSource = make(chan interface{})

	go wpool.Run()

	files, err := ioutil.ReadDir("./data/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//wpool.DataSource <- readJson(file.Name())
		fmt.Println(readJson(file.Name()))
	}
	wpool.Stop()
}

type TestWorker struct{}

func (w TestWorker) Do(data interface{}, i int) {
	//fmt.Printf("Goroutine number %d, is running. Input %s \n", i, rest)

}
func (w TestWorker) Stop() {

}
func readJson(fname string) models.Rest {
	var rest models.Rest
	jsonFile, err := os.Open("data/" + fname)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfull")
	defer jsonFile.Close()
	json.NewDecoder(jsonFile).Decode(&rest)

	return rest
}
func dbWrite(rest *models.Rest) {
	db, _ := db.Connect()
	Rdb := repo.RestDB{}
	Rdb.Conn(db)
	Rdb.Create(rest)
	Pdb := repo.ProductDB{}
	Pdb.Conn(db)
	for _, menu := range rest.Menu {
		Pdb.Create(menu)

	}

}
