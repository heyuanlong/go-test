package main

import(
	"fmt"
	"net"
	"time"
	"os"
)

/*


*/

func main() {

	service := "127.0.0.1:1234"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	checkErr(err)
	conn.Write([]byte{1,2,3})
	time.Sleep(10000 * time.Second)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
