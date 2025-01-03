package main

import (
	"fmt"

	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
)

func main() {
	product := struct {
		Name  string
		Price float64
	}{Name: "Phone", Price: -1.999}

	v := validator.New()
	v.Add("Name", rule.Rules{rule.Required()})
	v.Add("Price", rule.Rules{
		rule.Required(),
		rule.MinValue(0.0),
	})

	errs := v.Validate(product)
	if errs != nil {
		fmt.Println("Erros de validação:")
		for _, err := range errs {
			fmt.Printf("Campo: %s, Mensagem: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Usuário válido: ", product)
	}
}
