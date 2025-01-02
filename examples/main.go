package main

import (
	"fmt"

	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

type Address struct {
	Street string
	Zip    string
}

type User struct {
	Name    string
	Email   string
	Address Address
}

func New(name, email, street, zip string) (*User, []*validator_error.ValidatorError) {
	user := &User{
		Name:  name,
		Email: email,
		Address: Address{
			Street: street,
			Zip:    zip,
		},
	}

	if errs, ok := user.Validate(); !ok {
		return nil, errs
	}

	return user, nil
}

func (u *User) Validate() ([]*validator_error.ValidatorError, bool) {
	v := validator.New()

	v.Add("Name", rule.Rules{
		rule.Required(),
		rule.MinLength(3),
	})
	v.Add("Email", rule.Rules{
		rule.Required(),
		rule.ValidEmail(),
	})
	v.Add("Address.Street", rule.Rules{
		rule.Required(),
	})
	v.Add("Address.Zip", rule.Rules{
		rule.Required(),
		rule.ExactLength(5),
	})

	errs := v.Validate(*u)
	if len(errs) == 0 {
		return nil, true
	}
	return errs, false
}

func main() {
	user, errs := New("John", "example@gmail.com", "", "")
	if errs != nil {
		fmt.Println("Erros de validação:")
		for _, err := range errs {
			fmt.Printf("Campo: %s, Mensagem: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Tudo válido: ", user)
	}
}
