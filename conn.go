/**
 * Author:  Kernel Huang
 * Mail:    kernelman79@gmail.com
 * Date:    1/12/19
 * Time:    5:56 PM
 */

package websocket

import (
	"golang.org/x/net/websocket"
	"log"
)

type Conn struct {
	Ws        *websocket.Conn
	InChan    chan []byte
	OutChan   chan []byte
	CloseChan chan byte
}

var conn = new(Conn)

func (c *Conn) ws() *Conn {
	config.GetConfig()

	ws, err := websocket.Dial(config.Ws, "", config.Url)
	if err != nil {
		log.Fatal(err)
	}

	c.Ws = ws
	return c
}
