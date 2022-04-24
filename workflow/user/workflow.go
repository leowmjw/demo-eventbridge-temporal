package user

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
)

type UserReq struct {
	Action int
}

type MembershipSvc struct {
	// stuff here like ..?
}

type MembershipReq struct {
	Action int
}

// UserWorkflow - workflow for user whole lifecycle
func UserWorkflow(ctx workflow.Context, ur UserReq) error {

	// Sign up
	// from anonymous to user with identity .
	// now gets removed from the system ..
	// suspended from system ..

	return nil
}

// MembershipWorkflow - Workflow ...
func (m *MembershipSvc) MembershipWorkflow(ctx workflow.Context, mr MembershipReq) error {
	fmt.Println("Hellow!!!")
	// On-Board free
	// On-Board direct to Paid ..
	// Free to Paid ..
	// Did not pay so downgrade to free again
	// Marked as risky; temporarily deactivated ..

	return nil
}

func (m *MembershipSvc) privateFund() {

}
