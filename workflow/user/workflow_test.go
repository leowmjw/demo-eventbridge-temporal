package user

import (
	"app/internal/user"
	"go.temporal.io/sdk/workflow"
	"testing"
)

func TestMembershipSvc_MembershipWorkflow(t *testing.T) {
	type args struct {
		ctx workflow.Context
		mr  MembershipReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MembershipSvc{}
			if err := m.MembershipWorkflow(tt.args.ctx, tt.args.mr); (err != nil) != tt.wantErr {
				t.Errorf("MembershipWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMembershipSvc_privateFund(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MembershipSvc{}
			m.privateFund()
		})
	}
}

func TestUserSvc_UserWorkflow(t *testing.T) {
	type fields struct {
		User user.User
	}
	type args struct {
		ctx workflow.Context
		ur  UserReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSvc{
				User: tt.fields.User,
			}
			if err := u.UserWorkflow(tt.args.ctx, tt.args.ur); (err != nil) != tt.wantErr {
				t.Errorf("UserWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
