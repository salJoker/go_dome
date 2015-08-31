package main

import(
	"fmt"
	"os"
)

func main(){
//	var data string
//	ind,_ := fmt.Scan(&data)//到空白符处停止扫描
//	fmt.Printf("len:%d,data:%s\n",ind,data)


	var (
		arr1 string
		arr2 string
	)
	file,err := os.Open("resources/context.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
//	ind,err := fmt.Fscanln(file,&data)//Scanln，Fscanln和Sscanln在换行时结束读取，并要求数据连续出现；
	ind,err := fmt.Fscanf(file,"%s %s\n",&arr1,&arr2)//从文件中读取符合format格式的字符串
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("len:%d,data:%s,%s\n",ind,arr1,arr2)

	file,err = os.OpenFile("resources/context.txt",os.O_APPEND,0)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	_,err = fmt.Fprintf(file,"%s and %s\n","jack","tom")//向io.Writer中写入格式化后数据（覆盖写）
	if err != nil {
		fmt.Println(err.Error())
	}
	for i :=0 ; i < 10 ; i++{
		fmt.Fprintln(file,"abcd")
	}
}