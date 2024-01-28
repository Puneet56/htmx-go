package users

type User struct {
	Id       int
	Name     string
	Email    string
	Role     string // Admin, Project Manager, Team Member
	Password string
}
