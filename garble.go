// garble is a daemon that does evil things over each port it listens on

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	lAddr = flag.String("a", "localhost", "listen address")
	fAddr = flag.String("f", "server:80", "forward address")
)

type Filter struct {
	Run func(f *Filter, conn net.Conn)

	Port        int    // port this filter listens on
	Name        string // name of the filter
	Description string // help text for user friendliness
}

var filters = []*Filter{
	fFastClose,
	fTimeout,
	fNeverClose,
	fSlowButSteady,
	fRandom,
	fSlowButSteadyProxy,
}

func startFilter(f *Filter) {
	addr := fmt.Sprintf("%s:%d", *lAddr, f.Port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(f.Name, err)
	}

	log.Printf("Running '%s' on port '%d'", f.Name, f.Port)

	for {
		conn, err := l.Accept()
		if err != nil {
			// There's not really a need to die here
			log.Fatal(err)
		}
		go f.Run(f, conn)
	}

}

func main() {
	flag.Parse()

	for _, f := range filters {
		go startFilter(f)
	}

	// Block forever while the above goroutines execute
	select {}
}

var fFastClose = &Filter{
	Run:         runFastClose,
	Port:        1901,
	Name:        "Fast Close",
	Description: "Accept a connection then close it immediately",
}

func runFastClose(f *Filter, conn net.Conn) {
	conn.Close()
}

var fTimeout = &Filter{
	Run:         runTimeout,
	Port:        1902,
	Name:        "Timeout",
	Description: "Accept a connection, do nothing, then close after 30 seconds",
}

func runTimeout(f *Filter, conn net.Conn) {
	select {
	case <-time.After(30 * time.Second):
		conn.Close()
	}
}

var fNeverClose = &Filter{
	Run:         runNeverClose,
	Port:        1903,
	Name:        "Never Close",
	Description: "Accept a connection and never close it",
}

func runNeverClose(f *Filter, conn net.Conn) {
	select {}
}

var fSlowButSteady = &Filter{
	Run:         runSlowButSteady,
	Port:        1904,
	Name:        "Slow but Steady",
	Description: "Accept a connection and return one byte every second",
}

func runSlowButSteady(f *Filter, conn net.Conn) {
	c := time.Tick(1 * time.Second)
	for _ = range c {
		conn.Write([]byte("g"))
	}
}

var fRandom = &Filter{
	Run:         runRandom,
	Port:        1905,
	Name:        "Random Response",
	Description: "Accept a connection and return random binary data",
}

func runRandom(f *Filter, conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	rand.Read(buf)
	conn.Write(buf)
}

var fSlowButSteadyProxy = &Filter{
	Run:         runSlowButSteadyProxy,
	Port:        1906,
	Name:        "Slow but Steady Proxy",
	Description: "Proxy to a backend server and return the response one byte per second",
}

func runSlowButSteadyProxy(f *Filter, conn net.Conn) {
	defer conn.Close()
	proxy, err := net.Dial("tcp", *fAddr)
	if err != nil {
		// Not really a good reason to die here either
		log.Fatal(err)
	}

	go io.Copy(proxy, conn)

	c := time.Tick(1 * time.Second)
	for _ = range c {
		io.CopyN(conn, proxy, 1)
	}
}
