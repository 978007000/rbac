package model

type GroupCache struct {
	Name       string ` json:"name"`
	Uuid       string ` json:"uuid"`
	CustomerID int    ` json:"-"`

	Users    []string `json:"users,omitempty"` // many to many relation
	Policies []Policy `json:"policies,omitempty"`
}
