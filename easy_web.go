/*
这是一个简易web,访问 localhost:8081/hello 输出 hello, world! 
*/
package main
import (
	"io"
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
