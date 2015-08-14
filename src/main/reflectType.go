package main

import(
	"fmt"
	"reflect"
)

type Human interface {
	SayName()string
	Run()
}

type Man struct {
	Name string
}

func NewMan()*Man{
	return &Man{Name:"tomc"}
}

func (m *Man) SayName()string{
	return m.Name
}

func (m *Man) Run(){
	fmt.Println("%s is runing",m.Name)
}

type Woman struct{
	Name string
}

func NewWoman()*Woman{
	return &Woman{Name:"ali"}
}

func (w *Woman) SayName()string{
	return w.Name
}

func (w *Woman) Run(){
	fmt.Println("%s is runing",w.Name)
}

func main(){
	m := NewMan()

	mType := reflect.TypeOf(m)
	fmt.Println("name:",mType.Name())
	fmt.Println("Kind:",mType.Kind())
	fmt.Println("string:",mType.String())
	fmt.Println("method num:",mType.NumMethod())

	if method,exists := mType.MethodByName("Run");exists{
		fmt.Println("method.Name:",method.Name)
		fmt.Println("method.Value:",method.Func)
		fmt.Println("method.Type.NumIn:",method.Type.NumIn())

	}else{
		fmt.Println("method not exists")
	}

	mo := reflect.PtrTo(mType)
	fmt.Println("ptr:",mo)

	m_val := reflect.ValueOf(m)
	fmt.Println("value:",m_val)

	fmt.Println(reflect.Indirect(m_val))
}
