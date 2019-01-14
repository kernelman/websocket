/**
 * Author:  Kernel Huang
 * Mail:    kernelman79@gmail.com
 * Date:    1/12/19
 * Time:    5:56 PM
 */

package websocket

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Url string
	Ws  string
	Msg *simplejson.Json
}

var config = new(Config)

func (c *Config) GetConfig() {

	open, err := os.Open("websocket.json")
	if err != nil {
		log.Fatal(err)
	}

	file, err := ioutil.ReadAll(open)
	if err != nil {
		log.Fatal(err)
	}

	get, err := simplejson.NewJson(file)
	if err != nil {
		log.Fatal(err)
	}

	host, _ := get.Get("url").Get("host").String()
	port, _ := get.Get("url").Get("port").String()
	msg := get.Get("msg")

	c.Url = "http://" + host + ":" + port + "/"
	c.Ws = "ws://" + host + ":" + port + "/"
	c.Msg = msg
}
