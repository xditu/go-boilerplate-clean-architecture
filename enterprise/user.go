package enterprise

import "errors"

type User struct {
	Name     string
	Lastname string
	Username string
}

func (o *User) Fullname() string {
	return o.Name + o.Username
}

func (o *User) ValidateName() error {
	if len(o.Name) < 3 {
		return errors.New("nombre inválido, debe ser mayor a 3 caracteres")
	}
	return nil
}
