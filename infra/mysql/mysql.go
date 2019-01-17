package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/mochisuna/ssh-mysql-sample/config"
	"github.com/mochisuna/ssh-mysql-sample/infra/ssh"
)

const (
	NET_TCP       = "tcp"
	NET_MYSQL_TCP = "mysql+tcp"
)

type Client struct {
	*sql.DB
	*ssh.Client
}

func New(conf *config.DB, sshc *ssh.Client) (*Client, error) {
	// MySQLで使うプロトコルは一旦tcpで設定
	mysqlNet := NET_TCP
	if sshc != nil {
		// MySQLはプロトコルを更新
		mysqlNet = NET_MYSQL_TCP
		mysql.RegisterDial(NET_MYSQL_TCP, sshc.DialFunc())
	}
	dbConf := &mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Addr:                 conf.Host + conf.Port,
		Net:                  mysqlNet,
		DBName:               conf.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", dbConf.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &Client{db, sshc}, nil
}

func (c *Client) Close() {
	c.DB.Close()
}
