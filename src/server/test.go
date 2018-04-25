package main

import (
	"fmt"
	"reflect"
)

type SU2 struct{
	id int
}


type SU1 struct{
	SU2
	string
}


func main()  {
	s:= 110

	ca:=&SU1{string:"fff"}
	cb:=new(SU2)
	v,ok := interface{}(s).(string)
	ca.id = 5
	cb.id = 3
	//t:= interface{}(s).(string)

	fmt.Println("t =",ca.string)
	fmt.Println(ca.id)
	fmt.Println("v =",v)

	fmt.Println("ok = ", ok)

	er:=reflect.ValueOf(s)
	cweq:=[]string{"asd","ea"}
	fmt.Println("sd",cweq)
	fmt.Println(er)


}