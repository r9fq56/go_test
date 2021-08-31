package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
	"math/rand"
)

// type idGenerator string
 
// func (p idGenerator) Read(bs []byte) (int, error){

// 	return 5, io.EOF
// }

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s [tcp|udp] host:port\n", os.Args[0])
		os.Exit(1)
	}
	prot := os.Args[1]
	serv := os.Args[2]  
	conn, _ := net.Dial(prot, serv) 

	go test1(conn)
	test2(os.Stdout, conn)

}

func test1(dst io.Writer) {
	for {
		rand.Seed(time.Now().UnixNano())

		reasons := []string{"1","2","3","4","5","6","7","8","9","10"}

		len := len(reasons)
		rnd := rand.Intn(len)

		value := reasons[rnd]

		src := strings.NewReader(value + "\n")
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}

}

func test2(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

