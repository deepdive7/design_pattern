package behavioral

type AbstractCustomer interface {
	isNil() bool
	GetName() string
}

type RealCustomer struct {
	name string
}

func (rc *RealCustomer) isNil() bool {
	return false
}

func (rc *RealCustomer) GetName() string {
	return rc.name
}

type NullCustomer struct{}

func (nc *NullCustomer) isNil() bool {
	return true
}

func (rc *NullCustomer) GetName() string {
	return "Not Available in Customer Database"
}

type CustomerFactory struct {
	names []string
}

func (cf *CustomerFactory) GetCustomer(name string) AbstractCustomer {
	for _, e := range cf.names {
		if e == name {
			return &RealCustomer{name}
		}
	}
	return &NullCustomer{}
}
