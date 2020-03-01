package main

import "fmt"

type  Stu struct {
	Name string
	Age  int
}

func(p *Stu) SetName(name  string) *Stu{
	p.Name=name
	return p
}

func (p *Stu) SetAge(age int) *Stu{
	p.Age=age
	return  p
}

func(p *Stu) Print() {
	fmt.Printf("Name:%s,Age:%d\n",p.Name,p.Age)
}
func main()  {
	p :=&Stu{}
	p.SetAge(18)
	p.SetName("wmy")
	p.Print()
	p.SetName("wmy").SetAge(12).SetName("xiaowang").Print()
}