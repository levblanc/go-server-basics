package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// server is open on port 8080
	// use browser to make request at
	// localhost:8080
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Println(err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	defer conn.Close()

	i := 0

	for scanner.Scan() {
		line := scanner.Text()

		// line 1 of http request header is
		// the request line, in format of
		// <METHOD> <URI> <PROTOCOL>
		if i == 0 {
			reqInfo := strings.Fields(line)

			// print extracted info to connection
			fmt.Fprintf(conn, "method is: %s\n", reqInfo[0])
			fmt.Fprintf(conn, "uri is: %s\n\n", reqInfo[1])
			fmt.Fprintf(conn, "REQUEST HEADER: \n\n")

			// console print seperator
			fmt.Println("=========================")
		}

		if line == "" {
			break
		}

		i++

		// print full request header to connection
		fmt.Fprintf(conn, "%s\n", line)
		// print full request header to console
		fmt.Println(line)
	}
}
