/**
 * Author:  Kernel Huang
 * Mail:    kernelman79@gmail.com
 * Date:    1/12/19
 * Time:    5:56 PM
 */

package websocket

import (
	"encoding/json"
	"log"
)

type Send struct {
	Message map[string]interface{}
}

var To = new(Send)

func (s *Send) to() map[string]interface{} {
	c := conn.ws()
	s.Message, _ = config.Msg.Map()

	msg, err := json.Marshal(s.Message)
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Ws.Write(msg)
	if err != nil {
		log.Fatal(err)
	}

	return s.Message
}
