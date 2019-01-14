/**
 * Author:  Kernel Huang
 * Mail:    kernelman79@gmail.com
 * Date:    1/12/19
 * Time:    5:56 PM
 */

package websocket

type Client struct {
	send    Send
	receive Receive
}

var Cli = new(Client)

func (c *Client) Send() map[string]interface{} {
	return c.send.to()
}

func (c *Client) Receive() {
	c.receive.from()
}

func (c *Client) Close() {
	conn.Ws.Close()
}
