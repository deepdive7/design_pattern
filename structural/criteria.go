package structural

import (
	"fmt"
	"strings"
)

type person struct {
	name, gender, maritalStatus string
}

func (p person) equal(other person) bool {
	return p.name == other.name && p.gender == other.gender && p.maritalStatus == other.maritalStatus
}

//Criteria ...
type Criteria interface {
	meetCriteria([]person) []person
}

//CriteriaMale ...
type CriteriaMale struct{}

func (cMale CriteriaMale) meetCriteria(persons []person) []person {
	filteredPerson := make([]person, 0, len(persons))
	for _, p := range persons {
		if strings.Compare(p.gender, "Male") == 0 {
			filteredPerson = append(filteredPerson, p)
		}
	}
	return filteredPerson
}

//CriteriaFemale ...
type CriteriaFemale struct{}

func (cFemale CriteriaFemale) meetCriteria(persons []person) []person {
	filteredPerson := make([]person, 0, len(persons))
	for _, p := range persons {
		if strings.Compare(p.gender, "Female") == 0 {
			filteredPerson = append(filteredPerson, p)
		}
	}
	return filteredPerson
}

//CriteriaSingle ...
type CriteriaSingle struct{}

func (cSingle CriteriaSingle) meetCriteria(persons []person) []person {
	filteredPerson := make([]person, 0, len(persons))
	for _, p := range persons {
		if strings.Compare(p.maritalStatus, "Single") == 0 {
			filteredPerson = append(filteredPerson, p)
		}
	}
	return filteredPerson
}

//CriteriaAnd ...
type CriteriaAnd struct {
	criteriaA, criteriaB Criteria
}

func (cAnd CriteriaAnd) meetCriteria(persons []person) []person {
	return cAnd.criteriaA.meetCriteria(cAnd.criteriaB.meetCriteria(persons))
}

//CriteriaOr ...
type CriteriaOr struct {
	criteriaA, criteriaB Criteria
}

func (cOr CriteriaOr) meetCriteria(persons []person) []person {
	firstCriteriaItems := cOr.criteriaA.meetCriteria(persons)
	secondCriteriaItems := cOr.criteriaB.meetCriteria(persons)
	for _, p := range secondCriteriaItems {
		b := false
		for _, e := range firstCriteriaItems {
			b = b || e.equal(p)
		}
		if !b {
			firstCriteriaItems = append(firstCriteriaItems, p)
		}
	}
	return firstCriteriaItems
}

// PrintPersons ...
func PrintPersons(persons []person) {
	for _, p := range persons {
		fmt.Println("Person : [ Name : ", p.name,
			", Gender : ", p.gender,
			", Marital Status : ", p.maritalStatus,
			" ]")
	}
}
