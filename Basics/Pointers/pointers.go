package pointers

import "fmt"

type Person struct {
    name string
    age  int
}

func changeAge(p *Person) {
    p.age = 30
}

func Pointers() {
    p := Person{"Gungun", 22}
    fmt.Println("Before:", p)

    changeAge(&p) 
    fmt.Println("After:", p)
}
