package user

import (
	"go.temporal.io/sdk/workflow"
	"testing"
)

func TestMembershipWorkflow(t *testing.T) {
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
			if err := MembershipWorkflow(tt.args.ctx, tt.args.mr); (err != nil) != tt.wantErr {
				t.Errorf("MembershipWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserWorkflow(t *testing.T) {
	type args struct {
		ctx workflow.Context
		ur  UserReq
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
			if err := UserWorkflow(tt.args.ctx, tt.args.ur); (err != nil) != tt.wantErr {
				t.Errorf("UserWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_privateFund(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			privateFund()
		})
	}
}

func TestAnonUserWorkflow(t *testing.T) {
	type args struct {
		ctx workflow.Context
		req SessionReq
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
			if err := AnonUserWorkflow(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("AnonUserWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMembershipWorkflow1(t *testing.T) {
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
			if err := MembershipWorkflow(tt.args.ctx, tt.args.mr); (err != nil) != tt.wantErr {
				t.Errorf("MembershipWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserWorkflow1(t *testing.T) {
	type args struct {
		ctx workflow.Context
		ur  UserReq
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
			if err := UserWorkflow(tt.args.ctx, tt.args.ur); (err != nil) != tt.wantErr {
				t.Errorf("UserWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_privateFund1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			privateFund()
		})
	}
}
