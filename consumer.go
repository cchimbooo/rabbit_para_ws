package main

import "fmt"

func Consumir(m WsStruct) func(b []byte) error {
	return func(b []byte) error {
		fmt.Println(string(b))
		if err := m.BroadCast(b); err != nil {
			return err
		}
		return nil
	}
}

func Handle(e error) {
	fmt.Println(e.Error())
}
