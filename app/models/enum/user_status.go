package enum

//UserStatus is the status of a user
type UserStatus int

var (
	// UserApproved lets a user's posts be visible to all by default.
	UserApproved UserStatus = 1
	//UserDeleted is used for users that chose to delete their accounts
	UserDeleted UserStatus = 2
	//UserBlocked is used for users that have been blocked by staff members
	UserBlocked UserStatus = 3
	//UserActive is the default status for users; their posts will be hidden
	// from other users until the user is moved to "approved" status.
	UserActive UserStatus = 4
)

var userStatusIDs = map[UserStatus]string{
	UserDeleted:  "deleted",
	UserBlocked:  "blocked",
	UserApproved: "approved",
	UserActive:   "active",
}

var userStatusName = map[string]UserStatus{
	"approved": UserApproved,
	"deleted":  UserDeleted,
	"blocked":  UserBlocked,
	"active":   UserActive,
}

// MarshalText returns the Text version of the user status
func (status UserStatus) MarshalText() ([]byte, error) {
	return []byte(userStatusIDs[status]), nil
}

// UnmarshalText parse string into a user status
func (status *UserStatus) UnmarshalText(text []byte) error {
	*status = userStatusName[string(text)]
	return nil
}
