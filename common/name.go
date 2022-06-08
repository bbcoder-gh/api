package common

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (n *Name) String() string {
	return n.FirstName + " " + n.LastName
}
