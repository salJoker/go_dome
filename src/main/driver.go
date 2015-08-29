package main

import(
	//"os"
	"fmt"
	//"path/filepath"
	"io/ioutil"
	"os"
)

func main(){
//	for i,val := range os.Args {
//		fmt.Printf("%s : %d",val,i)
//	}

//	driver := os.Args[1]
//	drv,err := filepath.Abs(driver)
//	if err != nil {
//		return
//	}
//	fileindos,err := ioutil.ReadDir(drv)
//	if err != nil {
//		return
//	}
//	fmt.Printf("fileinfos length : %d\n",len(fileindos))
//	var fileCount int = 0
//	for _,fileinfo := range fileindos {
//		if !fileinfo.IsDir(){
//			fmt.Printf("filename:%s\n",fileinfo.Name())
//			fileCount++
//		}
//	}
//	fmt.Println("file count : ",fileCount)

	//name,err := ioutil.TempDir("","go_tmp")//使用系统默认临时目录，目录名称：go_temp+随机数
	name,err := ioutil.TempDir("resources","go_temp")//使用相对路径，在当前GOPATH目录下的resources中创建临时目录
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(name)

	//创建临时文件，与ioutil.TempDir类似
	//filename,err := ioutil.TempFile("","go_temp")//使用系统默认临时目录
	filename,err := ioutil.TempFile(name,"go_temp")
	if err != nil{
		return
	}
	fmt.Println(filename.Name())

	defer func(){
		filename.Close()
		os.Remove(filename.Name())
	}()
}
