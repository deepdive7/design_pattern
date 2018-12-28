package behavioral

import (
	"fmt"
	"testing"
)

func TestNullObj(t *testing.T) {
	CF := CustomerFactory{
		[]string{
			"Rob",
			"Bob",
			"Julia",
			"Laura",
		},
	}
	c1 := CF.GetCustomer("Rob")
	fmt.Println(c1.GetName())
	c2 := CF.GetCustomer("Bob")
	fmt.Println(c2.GetName())
	c3 := CF.GetCustomer("Julia")
	fmt.Println(c3.GetName())
	c4 := CF.GetCustomer("Laura")
	fmt.Println(c4.GetName())
	n1 := CF.GetCustomer("Toy")
	fmt.Println(n1.GetName())
}
