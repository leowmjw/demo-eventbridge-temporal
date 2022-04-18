package user

// - user - foundation service;
//		with sub-services:
//		membership, profile, linked-accounts, personalization
// 	Membership have different type based on free or paid.
//		Fees may be waived based on requirements.
//  Under Membership; benefits derived based on type and status.
// 		Each benefit has its own requirements.

//go:generate stringer -type=Type
type Type int

const (
	UNKNOWNT Type = iota
	FREE
	PAID
)

//go:generate stringer -type=Status
type Status int

const (
	UNKNOWNS Status = iota
	ACTIVE
	INACTIVE
)

type Membership struct {
	Type   Type
	Status Status
}

type Profile struct {
	FullName string
}

type LinkedAccounts struct {
	BankRef string
}
