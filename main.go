package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	loadEnv()
	http.HandleFunc("/", echo)
	println("serving")
	log.Fatal(http.ListenAndServe(env["ip"]+":"+env["port"], nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer delete(connections, c)
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		response := genResponse(string(message), c)
		log.Printf("%s %s -> %s\n", senderType(c), message, response)
		err = c.WriteMessage(mt, []byte(response))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func genResponse(msg string, c *websocket.Conn) string {
	if msg == "?id" {
		return connections[c].toStr()
	}
	if strings.Index(msg, "!id:") == 0 {
		kind := ConnectionKindStr(msg[4:])
		connections[c] = kind.toConnectionKind()
		return "hello " + string(kind)
	}
	if connections[c] == None {
		return "identify yourself"
	}
	if msg == "?track" {
		return "!track:" + track.str()
	}
	if strings.Index(msg, "!track:") == 0 {
		p := strings.Split(msg[7:], "|")
		if len(p) != 5 {
			return "wrong number of parameters"
		}
		value, err := strconv.Atoi(p[4])
		if err != nil {
			return "failed to update"
		}
		if value == 0 {
			c.WriteMessage(websocket.TextMessage, []byte("?track"))
			return "no length"
		}
		track = Track{p[0], p[1], p[2], p[3], value}
		updateMonitors()
		return "updated"
	}
	if msg == "?time" {
		getTime()
		return "noop"
	}
	if strings.Index(msg, "!time:") == 0 {
		updateTime(msg)
	}
	return "noop"
}

func updateMonitors() {
	for k, v := range connections {
		if v == Monitor {
			k.WriteMessage(websocket.TextMessage, []byte("!track:"+track.str()))
		}
	}
}

func getTime() {
	for c := range connections {
		c.WriteMessage(websocket.TextMessage, []byte("?time"))
	}
}

func updateTime(s string) {
	for k, v := range connections {
		if v == Monitor {
			k.WriteMessage(websocket.TextMessage, []byte(s))
		}
	}
}
