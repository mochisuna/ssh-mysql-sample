// main.go
package main

import (
	"flag"
	"fmt"

	"github.com/mochisuna/ssh-mysql-sample/application"
	"github.com/mochisuna/ssh-mysql-sample/config"
	"github.com/mochisuna/ssh-mysql-sample/domain"
	"github.com/mochisuna/ssh-mysql-sample/handler"
	"github.com/mochisuna/ssh-mysql-sample/infra"
	"github.com/mochisuna/ssh-mysql-sample/infra/mysql"
	"github.com/mochisuna/ssh-mysql-sample/infra/ssh"
)

func main() {
	// parse runtime options
	env := flag.String("e", "dev", "parameter of environment. default value: dev")
	storeID := flag.Int("s", 1, "parameter of store id. default value: 1")
	flag.Parse()

	// setting file path
	confPath := fmt.Sprintf("./_tools/config/%s/conf.toml", *env)
	outPath := fmt.Sprintf("./_tools/output/%s.csv", *env)

	// Setting Config
	conf := &config.Config{}
	if err := config.New(conf, confPath); err != nil {
		fmt.Printf("error in new-config. reason: %v\n", err)
		panic(err)
	}
	// init ssh client
	sshClient, err := ssh.New(&conf.SSH)
	if err != nil {
		fmt.Println("err in new ssh client. reason : " + err.Error())
	} else {
		defer sshClient.Close()
	}

	// init db client
	dbClient, err := mysql.New(&conf.DB, sshClient)
	if err != nil {
		fmt.Printf("erro in new db client. reason : %v\n", err)
		panic(err)
	}
	defer dbClient.Close()

	// init service
	storeRepo := infra.NewStoreRepository(dbClient)
	storeService := application.NewStoreService(storeRepo)
	services := &handler.Services{
		StoreService: storeService,
	}
	serviceHandler := handler.New(services)

	// handler action sample
	// sample A: save as csv file
	if err := serviceHandler.OutputCSV(outPath); err != nil {
		fmt.Printf("err: %v", err)
	}

	// sample B: get single store data
	val, err := serviceHandler.GetStore(domain.StoreID(*storeID))
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println(val)

	// sample B: get multiple store data set
	vals, err := serviceHandler.GetStores()
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println(vals)
}
