/**
 * Author:  Kernel Huang
 * Mail:    kernelman79@gmail.com
 * Date:    1/12/19
 * Time:    5:56 PM
 */

package tests

import (
	"log"
	"websocket"
)

func main() {

	send := websocket.Cli.Send()
	log.Println("Send: ", send)
	ch := make(chan int, 100)
	for i := 0; i <= 100; i++ {
		ch <- 1
		go websocket.Cli.Receive()
	}
}
