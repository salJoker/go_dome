package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
	strs := strings.FieldsFunc("abc-ndh-klj",func(r rune)bool{
		return r == '-'
	})
	fmt.Printf("%v\n",strs)


	num,err := strconv.ParseInt("370",0,0)
	if err != nil {
		fmt.Printf("%v",err)
	}
	fmt.Println(num)
	fmt.Printf("%d\n",32 << (^uint(0) >> 63))

	num,err = strconv.ParseInt("9223372036854775807",10,64)
	if err != nil {
		fmt.Printf("%v",err)
	}
	fmt.Println(num)

	dst := make([]byte,0)
	dst = strconv.AppendBool(dst,false)
	fmt.Printf("%v\n",dst)
}