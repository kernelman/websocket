/**
 * Author:  Kernel Huang
 * Mail:    kernelman79@gmail.com
 * Date:    1/12/19
 * Time:    5:56 PM
 */

package websocket

import "log"

type Receive struct {
}

var From = new(Receive)

func (r *Receive) from() {
	/*
		inChan := make(chan []byte,1000)
		closeChan := make(chan byte,1)
		conn.InChan = inChan
		conn.CloseChan = closeChan
	*/

	var msg = make([]byte, 1024)
	m, err := conn.Ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(msg[:m]))

	/*
		for{

			m, err := conn.Ws.Read(msg)
			if err != nil {
				goto ERR
			}

			log.Println(string(msg[:m]))

			// 阻塞并等待inChan有空闲位置
			select{
				case conn.InChan <- msg:
				case <- conn.CloseChan:		// closeChan 感知 conn断开
				goto ERR
			}

			ERR: conn.Ws.Close()
		}
	*/

	/*
		var msg = make([]byte, 512)
		c := conn.Ws
		m, err := c.Read(msg)
		if err != nil {
			log.Fatal(err)
		}

		return string(msg[:m])
	*/
}
