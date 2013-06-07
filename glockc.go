package glockc

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	conn net.Conn
	last string
}

func (self *Client) initialize() {
	self.conn = nil
}

func (self *Client) command(command ...string) (int, error) {
	var err error
	sendCommand := ""
	l := len(command)
	for i, s := range command {
		if (i + 1) == l {
			sendCommand = fmt.Sprintf("%s %s\n", sendCommand, strings.TrimSpace(s))
		} else {
			sendCommand = fmt.Sprintf("%s %s ", sendCommand, strings.TrimSpace(s))
		}
	}
	fmt.Fprintf(self.conn, sendCommand)
	self.last, err = bufio.NewReader(self.conn).ReadString('\n')
	self.last = strings.TrimSpace(self.last)
	if err != nil {
		return 0, err
	}
	results := strings.Split(self.last, " ")
	return strconv.Atoi(results[0])
}

func (self *Client) Get(lock string, shared bool) (int, error) {
	if shared {
		return self.command("sg", lock)
	} else {
		return self.command("g", lock)
	}
}

func (self *Client) Release(lock string, shared bool) (int, error) {
	if shared {
		return self.command("sr", lock)
	} else {
		return self.command("r", lock)
	}
}

func (self *Client) Inspect(lock string, shared bool) (int, error) {
	if shared {
		return self.command("si", lock)
	} else {
		return self.command("i", lock)
	}
}

func (self *Client) Name(name string) (int, error) {
	return self.command("iam", name)
}

func (self *Client) DebugLast() string {
	return self.last
}

func New(host string, port int) (Client, error) {
	var err error
	client := Client{}
	client.initialize()
	client.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	return client, err
}
