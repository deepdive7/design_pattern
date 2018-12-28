package structural

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	CEO := &Employee{"John", "CEO", 30000, []*Employee{}}

	headSales := &Employee{"Robert", "Head Sales", 20000, []*Employee{}}

	headMarketing := &Employee{"Michel", "Head Marketing", 20000, []*Employee{}}

	clerk1 := &Employee{"Laura", "Marketing", 10000, []*Employee{}}
	clerk2 := &Employee{"Bob", "Marketing", 10000, []*Employee{}}

	salesExecutive1 := &Employee{"Richard", "Sales", 10000, []*Employee{}}
	salesExecutive2 := &Employee{"Rob", "Sales", 10000, []*Employee{}}

	CEO.add(headSales)
	CEO.add(headMarketing)

	headSales.add(salesExecutive1)
	headSales.add(salesExecutive2)

	headMarketing.add(clerk1)
	headMarketing.add(clerk2)

	//打印该组织的所有员工
	fmt.Println(CEO.toString())
	for _, headEmployee := range CEO.subordinates {
		fmt.Println(headEmployee.toString())
		for _, employee := range headEmployee.subordinates {
			fmt.Println(employee.toString())
		}
	}
}
