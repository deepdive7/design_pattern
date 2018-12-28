package structural

import (
	"fmt"
)

// Employee ...
type Employee struct {
	name, dept   string
	salary       int
	subordinates []*Employee
}

func (e Employee) equal(otherE *Employee) bool {
	return e.name == otherE.name && e.dept == otherE.dept
}
func (e *Employee) add(newE *Employee) {
	if len(e.subordinates) < cap(e.subordinates) {
		e.subordinates[len(e.subordinates)] = newE
	} else {
		e.subordinates = append(e.subordinates, newE)
	}
}

func (e *Employee) remove(delE *Employee) {
	for index, subE := range e.subordinates {
		if subE.equal(delE) {
			if index == len(e.subordinates) {
				e.subordinates = append(e.subordinates[:index])
			} else {
				e.subordinates = append(e.subordinates[:index], e.subordinates[index+1:]...)
			}
			break
		}
	}
}

func (e *Employee) toString() string {
	return fmt.Sprintln("Employee :[ Name : ", e.name,
		", dept : ", e.dept, ", salary :",
		e.salary, " ]")
}
