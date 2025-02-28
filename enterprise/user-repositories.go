package enterprise

type UserRepositories interface {
	Add(u *User) error
	Remove(u User) error
	Update(u User) error
	Get(username string) (User, error)
}
