package user

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/mocks"
	"net"
	"time"
)

// - user - foundation service;
//		with sub-services:
//		membership, profile, linked-accounts, personalization
// 	Membership have different type based on free or paid.
//		Fees may be waived based on requirements.
//  Under Membership; benefits derived based on type and status.
// 		Each benefit has its own requirements.

//go:generate stringer -type=MembershipType
type MembershipType int

const (
	UNKNOWN_MT MembershipType = iota
	FREE
	PAID
)

//go:generate stringer -type=MembershipStatus
type MembershipStatus int

const (
	UNKNOWN_MS MembershipStatus = iota
	ACTIVE
	INACTIVE
)

//go:generate stringer -type=UserStatus
type UserStatus int

const (
	UNKNOWN_US UserStatus = iota
	ACTIVE_US
	INACTIVE_US
	SUSPENDED_US
)

// Have heavy weight stuff like Temporal client one per package? svc ..
var temporalClient client.Client
var mockTemporalClient mocks.Client

type AnonUser struct {
	ID         string
	SessionID  string
	LastIP     net.Addr
	LastActive time.Time
}

type User struct {
	ID      string
	AnonIDs []AnonUser
	Status  UserStatus
}

type Membership struct {
	Type   MembershipType
	Status MembershipStatus
}

type Profile struct {
	FullName     string
	PrimaryEmail string
	PrimaryPhone string
}

type LinkedAccounts struct {
	BankRef string
}

type UserSvc struct {
	// Probably should not be here ..
	// HTTP REST endpoint; better elsewhere?
	// Internal User struct ..
	User
}

type MembershipSvc struct {
	// stuff here like ..?
}

// OnBoardFromAnonymous allows extraction from JWTs to full user info; tie to identity
func (u UserSvc) OnBoardFromAnonymous() error {
	// finish off the lifecycle of existing anon session (if not already expired) ..
	return nil
}

func (u UserSvc) UpdateUser() error {
	// add if non-existent via kicking off Lifecycle of User
	return nil
}

func (u UserSvc) UpdateProfile() error {
	// add if non existent
	return nil
}
