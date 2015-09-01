package main

import(
	"bufio"
	"fmt"
//	"strings"
	"os"
	"time"
)

func main(){

//	read := bufio.NewReader(strings.NewReader("https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter01/01.4.md"))
	file,_ := os.Open("resources/context.txt")
//	read := bufio.NewReader(file)

//	line,_ := read.ReadSlice('T')
//	line,_,_ := read.ReadLine()//受bufio的最大缓存限制，但返回值为复制
//	fmt.Printf("This is %d \n",len(line))

//	n,_ := read.ReadSlice('T')
//	_,isPrefix,err := read.ReadLine()
//	fmt.Printf("This is %s \n",line)
//	fmt.Printf("This is prefix :%v err:%v\n",isPrefix,err)

//	line,err := read.ReadBytes('\n')
//	fmt.Printf("line :%s ; length:%d,err:%v",line,len(line),err)

//	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)//实际上，缓存大小是16！别问我为什么，看源码！
//	go Peek(reader)
//	go func(){
//		line,_ := reader.ReadBytes('\t')
//		fmt.Printf("Bytes : %s\n",line)
//	}()
//	time.Sleep(1e8)

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fmt.Printf("%s \n",scanner.Text())
	}

}

func Peek(reader *bufio.Reader) {
	line, _ := reader.Peek(15)//查看缓存中前15个字节，不取出
	fmt.Printf("1%s\n", line)
	time.Sleep(1)
	fmt.Printf("2%s\n", line)
}