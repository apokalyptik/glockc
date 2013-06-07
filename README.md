glockc
======

A client for the glockd network locking server (http://code.svn.wordpress.org/glockd/)

Installation
------------
go get github.com/apokalyptik/glockc

Example
-------

```go
package main

import(
	"github.com/apokalyptik/glockc"
	"log"
)

func main() {
	var locks glockc.Client
	var err error
	var lock int
	locks, err = glockc.New("127.0.0.1", 9999)
	if err != nil {
		log.Printf( "Unable to connect to locking server: %+v", err )
		return
	}
	lock, err = locks.Get("somenamedlock", false)
	if err != nil {
		log.Printf( "Error aquiring lock: %+v", err )
		return
	}
	if lock == 1 {
		log.Printf( "Aquired exclusive lock" )
	} else {
		log.Printf( "Another client already has this exclusive lock" )
	}
}
```

Package Docs
------------

type Client struct {
	    // contains filtered or unexported fields
}

func New(host string, port int) (Client, error)

func (self *Client) DebugLast() string

func (self *Client) Get(lock string, shared bool) (int, error)

func (self *Client) Inspect(lock string, shared bool) (int, error)

func (self *Client) Name(name string) (int, error)

func (self *Client) Release(lock string, shared bool) (int, error)
