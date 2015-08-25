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
//		src := args[1]
//		fmt.Println("src file ",src)
//		srcf,err := os.Open(src)
//		if err != nil {
//			return
//		}
//		srcf.Close()
//
//		data := make([]byte,1024);
//		i,err := srcf.Read(data)
//		fmt.Println(data)
//		for ;i != 0 && err == nil;{
//			fmt.Println(data)
//			i,err = srcf.Read(data)
//		}
	}

//	srcFile,err :=  os.Open(args[0])
//	if err != nil{
//		fmt.Printf("%s is not file",args[0])
//		return
//	}
//	defer srcFile.Close()
//
//	destFile,err := os.Create(args[1])
//	if err != nil {
//		fmt.Printf("%s is not file",args[1])
//		return
//	}
//	defer destFile.Close()
//
//	if _,err := io.Copy(destFile,srcFile);err!=nil{
//		fmt.Println(err.Error())
//		return
//	}
}

func ReadFrom(reader io.Reader,num int)([]byte,error){
	data := make([]byte,num)
	p,err := reader.Read(data)
	if p > 0 {
		data = data[:p]
	}
	return data,err
}