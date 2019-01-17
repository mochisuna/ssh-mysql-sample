package ssh

import (
	"io/ioutil"
	"net"

	"github.com/mochisuna/ssh-mysql-sample/config"
	"golang.org/x/crypto/ssh"
)

const NET_TCP = "tcp"

type Client struct {
	*ssh.Client
}

func New(conf *config.SSH) (*Client, error) {
	sshKey, err := ioutil.ReadFile(conf.Key)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(sshKey)
	if err != nil {
		return nil, err
	}
	hostKeyCallbackFunc := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	sshConf := &ssh.ClientConfig{
		User: conf.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallbackFunc,
	}
	client, err := ssh.Dial(NET_TCP, conf.Host+conf.Port, sshConf)
	return &Client{client}, err
}

func (c *Client) DialFunc() func(addr string) (net.Conn, error) {
	return func(addr string) (net.Conn, error) {
		// dialはtcpでやりとり
		return c.Dial(NET_TCP, addr)
	}
}

func (c *Client) Close() {
	c.Client.Close()
}
