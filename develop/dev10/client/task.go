package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

//Реализовать простейший telnet-клиент.
//
//Примеры вызовов:
//go-telnet --timeout=10s host port go-telnet mysite.ru 8080
//go-telnet --timeout=3s 1.1.1.1 123
//
//Требования:
//1. Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
//2. Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
//3. При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

type Client struct {
	host     string
	port     string
	connTime time.Duration
}

func NewClient(h, p, cTime string) *Client {
	t, err := strconv.Atoi(cTime)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		host:     h,
		port:     p,
		connTime: time.Duration(t) * time.Second,
	}
}

func dial(client *Client) {
	ch := make(chan string, 1)
	cr := make(chan string, 1)
	defer close(ch)
	defer close(cr)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := net.DialTimeout("tcp", client.host+":"+client.port, client.connTime)
	if err != nil {
		time.Sleep(client.connTime)
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(ctx context.Context, s chan string) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				rd := bufio.NewReader(os.Stdin)
				text, err := rd.ReadString('\n')
				if err == io.EOF {
					_, err := conn.Write([]byte("Ctrl+D"))
					if err != nil {
						log.Println(err)
					}
					s <- "Ctrl+D"
					return
				} else if err != nil {
					fmt.Fprintf(os.Stderr, "error reading from stdin: %v\n", err)
				}
				_, err = conn.Write([]byte(text))
				if err != nil {
					fmt.Fprintf(os.Stderr, "error writing to server: %v\n", err)
				}
			}

		}
	}(ctx, ch)

	go func(ctx context.Context, s chan string) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fb, err := bufio.NewReader(conn).ReadString('\n')
				if err == io.EOF {
					s <- "crash server"
					return
				} else if err != nil {
					log.Println(err)
				}
				fmt.Println("from server :" + fb)
			}
		}
	}(ctx, cr)

	select {
	case <-cr:
		cancel()
		fmt.Println("press enter to exit")
	case <-ch:
		cancel()
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}
	wg.Wait()
}

func main() {
	t := flag.String("timeout", "10", "время на работу с сервером")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("should be host and port")
	}
	port := args[1]
	host := args[0]

	client := NewClient(host, port, *t)

	dial(client)

	fmt.Println("Connection closed...")
}
