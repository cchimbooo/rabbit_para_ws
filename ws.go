package main

import (
	"gopkg.in/olahol/melody.v1"
)

func NovoWs() WsStruct {
	return WsStruct{melody.New()}
}

type WsStruct struct {
	Ms *melody.Melody
}

func (w WsStruct) BroadCast(b []byte) error {
	return w.Ms.Broadcast(b)
}
