package main

import (
	"fmt"

	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u *User) Validate() ([]*validator_error.ValidatorError, bool) {
	v := validator.New()

	v.Add("Name", rule.Rules{
		rule.Required(),
		rule.MinLength(3),
		rule.MaxLength(50),
	})
	v.Add("Email", rule.Rules{
		rule.Required(),
		rule.ValidEmail(),
	})
	v.Add("Age", rule.Rules{
		rule.MinValue(18),
	})

	errs := v.Validate(*u)
	if len(errs) == 0 {
		return nil, true
	}
	return errs, false
}

func main() {
	user := User{Name: "Jo", Email: "invalid_email", Age: 15}
	if errs, ok := user.Validate(); !ok {
		fmt.Println("Erros de validação:")
		for _, err := range errs {
			fmt.Printf("Campo: %s, Mensagem: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Tudo válido!")
	}
}
