package user

import (
	"testing"
)

// Test HappyPath of AnonUser complete onBoarding becoming full User
// TestHappyPathUserOnBoarding goes through full success
func TestHappyPathUserOnBoarding(t *testing.T) {

	// This should use the moclsbut how ...
	// need the mock version ..
	//tc, myerr := client.NewClient(client.Options{})
	tc := &mockTemporalClient

	//if myerr != nil {
	//	t.Fatalf("")
	//}

	u := UserSvc{
		TemporalClient: tc,
	}

	// When below condition happen; then return the following ..
	// Is the variable callstill need their red daddy ..
	call := mockTemporalClient.On("ExecuteWorkflow",
		nil).Return("test")
	// what about this? no clue ..
	call.Return()

	err := u.OnBoardFromAnonymous()
	if err != nil {
		t.Fatalf("OnBoardFromAnonymous ERR: %s", err.Error())
	}
	// SIgnal should be sent; they will need double-checking ..
	t.Fatal("TODO")

}

// Test Fraudulent AnonUser trying to login for existing full User;
//	flagged and block

func TestMembershipStatus_String(t *testing.T) {
	tests := []struct {
		name string
		i    MembershipStatus
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembershipType_String(t *testing.T) {
	tests := []struct {
		name string
		i    MembershipType
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatus_String(t *testing.T) {
	tests := []struct {
		name string
		i    UserStatus
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSvc_OnBoardFromAnonymous(t *testing.T) {
	type fields struct {
		User User
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserSvc{
				User: tt.fields.User,
			}
			if err := u.OnBoardFromAnonymous(); (err != nil) != tt.wantErr {
				t.Errorf("OnBoardFromAnonymous() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserSvc_UpdateProfile(t *testing.T) {
	type fields struct {
		User User
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserSvc{
				User: tt.fields.User,
			}
			if err := u.UpdateProfile(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserSvc_UpdateUser(t *testing.T) {
	type fields struct {
		User User
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserSvc{
				User: tt.fields.User,
			}
			if err := u.UpdateUser(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
