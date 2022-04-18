package user

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
)

type MembershipSvc struct {
	// stuff here like ..?
}

type MembershipReq struct {
}

// MembershipWorkflow - Workflow ...
func (m *MembershipSvc) MembershipWorkflow(ctx workflow.Context, mr MembershipReq) error {
	fmt.Println("Hellow!!!")

	return nil
}

func (m *MembershipSvc) privateFund() {

}
