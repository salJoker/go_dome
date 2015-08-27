package main

import (
	"fmt"
//	"io"
	"os"
	"io"
	"strings"
)

func main(){
	args := os.Args
	if args == nil || len(args) < 1 {
		fmt.Println("iocpy src dest")
		return
	}

	fmt.Println("args length ",len(args))

	if len(args) == 2 {
		data,err := ReadFrom(strings.NewReader(args[1]),1024)
		if err != nil {
			fmt.Printf("%s is error",args[1])
		}
		fmt.Printf("%s",data)
	}
}

func ReadFrom(reader io.Reader,num int)([]byte,error){
	data := make([]byte,num)
	p,err := reader.Read(data)
	if p > 0 {
		data = data[:p]
	}
	return data,err
}