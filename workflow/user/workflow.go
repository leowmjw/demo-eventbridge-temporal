package user

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
)

type SessionReq struct {
	Action int
}

type UserReq struct {
	Action int
}

type MembershipReq struct {
	Action int
}

// AnonUserWorkflow - workflow for those tracked via session only; goes away after 1 day
func AnonUserWorkflow(ctx workflow.Context, req SessionReq) error {
	// Check email; check phone?
	// fraud check if behavior for existing and now extra info does not match
	// then on-boarded ..
	return nil
}

// UserWorkflow - workflow for user whole lifecycle
func UserWorkflow(ctx workflow.Context, ur UserReq) error {

	// Sign up
	// from anonymous to user with identity; KYC .
	// now gets removed from the system ..
	// suspended from system ..

	return nil
}

// MembershipWorkflow - Workflow ...
func MembershipWorkflow(ctx workflow.Context, mr MembershipReq) error {
	fmt.Println("Hellow!!!")
	// On-Board free
	// On-Board direct to Paid ..
	// Free to Paid ..
	// Did not pay so downgrade to free again
	// Marked as risky; temporarily deactivated ..

	return nil
}

func privateFund() {

}
