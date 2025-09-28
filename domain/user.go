package domain

type UserRole string

const (
	GuestUser  UserRole = "guest"
	MemberUser UserRole = "member"
)

type User struct {
	ID      string
	Role    UserRole
	Email   string
	Picture string
	Name    string
}

func (u User) IsGuest() bool {
	return u.Role == GuestUser
}

func (u User) IsMember() bool {
	return u.Role == MemberUser
}
