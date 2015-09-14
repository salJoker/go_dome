package main

import(
	"regexp"
	"fmt"
)

func main(){
	if isOk,err := regexp.Match("(FC)\\d{7}",[]byte("FC7500020")); isOk && err ==nil{
		fmt.Printf("%v\n",isOk)
	}

	if isOk,err := regexp.MatchString("(FC)","FC7500020");isOk && err == nil {
		fmt.Printf("%v\n",isOk)
	}

	reg,err := regexp.Compile("^(FC)\\d{7}$")
	if err != nil {
		fmt.Printf("%v\n",err)
	} else if data := reg.Find([]byte("FC6650202"));data != nil{
		fmt.Printf("%s\n",data)
	}
}
