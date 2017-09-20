package main
import (
	"os"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-test/filetrans/pb"
	"net"
	"bytes"
	"encoding/binary"
)
var mapf map[string] *os.File
func main() {
	mapf = make(map[string] *os.File)
	listen_sock,err := net.Listen("tcp",":6000")
	checkErr(err)
	defer listen_sock.Close()
	for{
		new_conn,err := listen_sock.Accept()
		if err != nil {
			continue
		}
		go handleClient(new_conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	tmpBuf := make([]byte,0)
	buf := make ([]byte ,10244)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("rrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr",err,n)
			return
		}
		tmpBuf = Unpack(append(tmpBuf, buf[:n]...))
		if len(tmpBuf) > 102400 {
			fmt.Println("package too big")
			return
		}
	}
}
func Unpack(buf []byte) []byte {
	var i = 0
	for {
		buf = buf[i:]
		var mlen= len(buf)
		if mlen > 4 {
			value := binary.LittleEndian.Uint32(buf[0:4])
			if uint32(mlen) >= value {
				fmt.Println("value:",value)
				dealmsg(buf[4:value])
				i = int(value)
			}else{
				break
			}
		}else{
			break
		}
	}
	return buf[:]
}
func dealmsg(data []byte)  {
	// 进行解码
	ms := &pb.File{}
	err := proto.Unmarshal(data, ms)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}
	fmt.Println("name:",ms.GetName())
	fmt.Println("len:",len(ms.GetContent()))
	//fmt.Println("content:",ms.GetContent())
	if ms.GetType() == 0{
		dealdir(ms)
	}else{
		dealfile(ms)
	}
}
func dealdir( ms *pb.File)  {
	os.Mkdir(ms.GetName(),os.ModePerm)
}

func dealfile( ms *pb.File)  {
	if f, ok := mapf[ms.GetName()]; ok {
		//n, _ := f.Seek(0, os.SEEK_END)
		f.Write([]byte(ms.GetContent()))
		if ms.GetRemain() == 0{
			f.Close()
			delete(mapf,ms.GetName())
		}
	}else{
		f, err := os.OpenFile(ms.GetName(),os.O_CREATE | os.O_TRUNC,0666)
		if err != nil {
			fmt.Println("OpenFile error: ", err)
			return
		}

		f.Write( []byte(ms.GetContent()) )
		if ms.GetRemain() == 0{
			f.Close()
		}else {
			mapf[ms.GetName()] = f
		}

	}
}
func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}