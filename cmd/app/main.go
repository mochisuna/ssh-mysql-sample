// main.go
package main

import (
	"fmt"
	"os"

	"github.com/mochisuna/ssh-mysql-sample/application"
	"github.com/mochisuna/ssh-mysql-sample/config"
	"github.com/mochisuna/ssh-mysql-sample/domain"
	"github.com/mochisuna/ssh-mysql-sample/handler"
	"github.com/mochisuna/ssh-mysql-sample/infra"
	"github.com/mochisuna/ssh-mysql-sample/infra/mysql"
	"github.com/mochisuna/ssh-mysql-sample/infra/ssh"
	"github.com/urfave/cli/v2"
)

func initClient(env string) (*ssh.Client, *mysql.Client) {
	// setting file path
	confPath := fmt.Sprintf("./_tools/config/%s/conf.toml", env)

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
	}

	// init db client
	dbClient, err := mysql.New(&conf.DB, sshClient)
	if err != nil {
		fmt.Printf("erro in new db client. reason : %v\n", err)
		panic(err)
	}
	return sshClient, dbClient

}

func initHandler(dbClient *mysql.Client) *handler.Handler {
	// init service
	storeRepo := infra.NewStoreRepository(dbClient)
	storeService := application.NewStoreService(storeRepo)
	services := &handler.Services{
		StoreService: storeService,
	}
	return handler.New(services)
}

func act(c *cli.Context) error {
	env := c.String("e")
	storeID := c.Int("s")
	outPath := fmt.Sprintf("./_tools/output/%s.csv", env)
	sshClient, dbClient := initClient(env)
	defer sshClient.Close()
	defer dbClient.Close()
	serviceHandler := initHandler(dbClient)

	// handler action sample
	// sample A: save as csv file
	if err := serviceHandler.OutputCSV(outPath); err != nil {
		fmt.Printf("err: %v", err)
	}

	// sample B: get single store data
	val, err := serviceHandler.GetStore(domain.StoreID(storeID))
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
	return nil
}

func main() {
	en := cli.StringFlag{
		Name:    "env",
		Aliases: []string{"e"},
		Value:   "dev",
		Usage:   "parameter of environment. default value: dev",
	}
	si := cli.IntFlag{
		Name:    "store_id",
		Aliases: []string{"s", "sid"},
		Value:   1,
		Usage:   "parameter of store id. default value: 1",
	}
	app := &cli.App{
		Flags: []cli.Flag{
			&si,
			&en,
		},
	}
	app.Action = act
	app.Run(os.Args)
}
