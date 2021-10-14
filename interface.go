package main

import "fmt"

/**
  * refer: https://www.jianshu.com/p/06704ad058f1
 */

type Inter interface {
	Get() int
	Set(int)
}

type St struct {
	age int
}

// (s St) or (s *St) both OK
func (s St) Get() int {
	return s.age
}

func (s *St) Set(age int)  {
	s.age = age
}

func test(inter Inter)  {
	inter.Set(10)
	fmt.Println(inter.Get())
}

func main() {
	s := St{}
	test(&s)
}
