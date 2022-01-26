package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
	"time"
)

func ConsumerRabbit(processar func([]byte) error, tratarError func(error)) {

	url := gerarUrlConexao(
		os.Getenv("RABBITUSER"),
		os.Getenv("RABBITPWD"),
		os.Getenv("RABBITHOST"),
		os.Getenv("RABBITPORT"),
		os.Getenv("RABBITVHOST"),
	)
	fmt.Println(url)
	conn, errDial := amqp.Dial(url)
	if errDial != nil {
		// Dorme um pouco ante de reconectar
		time.Sleep(20 * time.Second)
		conn, errDial = amqp.Dial(url)
		if errDial != nil {
			log.Fatalln("falha ao se conectar")
		}
	}

	ch, errConn := conn.Channel()
	if errConn != nil {
		log.Fatalln(errConn.Error())
	}

	q, errDeclarar := ch.QueueDeclare(
		os.Getenv("RABBITFILA"),
		true,
		false,
		false,
		false,
		nil,
	)

	if errDeclarar != nil {
		log.Fatalln(errDeclarar.Error())
	}

	fmt.Println("Conectou . . .")

	defer conn.Close()
	defer ch.Close()

	msgs, errCons := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if errCons != nil {
		log.Fatalln(errCons.Error())
	}

	// Para semrpe
	forever := make(chan bool)

	go func() {
		for d := range msgs {

			// se sobrar tempo fazer uma sinaleira para implementar mais de um leitor
			// testar se tem como fazer isso sem dar problema

			if errProcess := processar(d.Body); errProcess != nil {
				tratarError(errProcess)
			}
			// da o acknoledge
			_ = d.Ack(false)
		}
	}()

	fmt.Println("Rodando...")
	<-forever
}

func gerarUrlConexao(usuario, senha, host, port, vhost string) string {
	b := strings.Builder{}
	b.WriteString("amqp://")
	b.WriteString(usuario)
	b.WriteString(":")
	b.WriteString(senha)
	b.WriteString("@")
	b.WriteString(host)
	if port != "" {
		b.WriteString(":")
	}
	b.WriteString(port)
	b.WriteString("/")
	b.WriteString(vhost)
	return b.String()
}
