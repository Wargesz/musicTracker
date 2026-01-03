package main

import "github.com/gorilla/websocket"

type ConnectionKind int
type ConnectionKindStr string

const (
	None ConnectionKind = iota
	Client
	Monitor
)

var connections = map[*websocket.Conn]ConnectionKind{}

func (k ConnectionKindStr) toConnectionKind() ConnectionKind {
	if k == "Monitor" {
		return Monitor
	}
	if k == "Client" {
		return Client
	}
	return None
}

func (k ConnectionKind) toStr() string {
	if k == Monitor {
		return "Monitor"
	}
	if k == Client {
		return "Client"
	}
	return "None"
}

func senderType(c *websocket.Conn) string {
	return connections[c].toStr()
}
