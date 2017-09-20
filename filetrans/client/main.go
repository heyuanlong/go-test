package main
import (
	"path/filepath"
	"os"
	"io"
	"fmt"
	"flag"
	"github.com/golang/protobuf/proto"
	"go-test/filetrans/pb"
	"net"
	"encoding/binary"
)
var conn net.Conn
func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if f.IsDir() {
			dealDir(path)
		}else {
			dealFile(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
func dealDir(path string) {
	ms := &pb.File{
		Type:proto.Int32(0),
		Name:proto.String(path),
		Remain:proto.Int32(0),
		Content:proto.String(""),
	}
	data,err := proto.Marshal(ms)
	if err != nil{
		fmt.Println("marshaling error: ", err)
	}
	packsize := uint32ToByte( uint32(len(data)) + 4 )
	pack := append(packsize,data...)
	conn.Write(pack)
	fmt.Println(path,packsize)
}

func dealFile(path string) {

	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("dealFile error: ", err)
		return
	}
	defer fi.Close()
	fiinfo, err := fi.Stat()
	fiSize := fiinfo.Size()
	buff := make([]byte,10240)
	ms := &pb.File{
		Type:proto.Int32(1),
		Name:proto.String(path),
	}
	var i = 0
	for  {
		n, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			fmt.Println("Read error: ", err)
			return
		}
		if n == 0 && i!= 0{
			break
		}
		ms.Remain = proto.Int32( int32(fiSize) - int32(n) )
		fiSize -= int64(n)
		ms.Content = proto.String(string(buff[:n]))
		data,err := proto.Marshal(ms)
		if err != nil{
			fmt.Println("marshaling error: ", err)
		}
		packsize := uint32ToByte( uint32(len(data)) + 4 )
		pack := append(packsize,data...)
		conn.Write(pack)
		fmt.Println(path,packsize)

		i += 1
	}

}

func uint32ToByte(i uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, i)
	return bytes
}

func main(){
	flag.Parse()
	root := flag.Arg(0)

	var err error
	conn, err = net.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Println("dial error ",err)
		return
	}
	defer conn.Close()
	getFilelist(root)


}
