package structural

import (
	"fmt"
	"testing"
)

func TestCriteria(t *testing.T) {
	persons := []person{
		person{"Robert", "Male", "Single"},
		person{"John", "Male", "Married"},
		person{"Laura", "Female", "Married"},
		person{"Diana", "Female", "Single"},
		person{"Mike", "Male", "Single"},
		person{"Bobby", "Male", "Single"},
	}
	male := CriteriaMale{}
	female := CriteriaFemale{}
	single := CriteriaSingle{}
	singleMale := CriteriaAnd{single, male}
	singleOrFemale := CriteriaOr{single, female}
	fmt.Println("Males:")
	PrintPersons(male.meetCriteria(persons))
	fmt.Println("Females:")
	PrintPersons(female.meetCriteria(persons))
	fmt.Println("Single:")
	PrintPersons(single.meetCriteria(persons))
	fmt.Println("SingleMales:")
	PrintPersons(singleMale.meetCriteria(persons))
	fmt.Println("Single Or Females:")
	PrintPersons(singleOrFemale.meetCriteria(persons))
}
