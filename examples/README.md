# Validador de Estruturas em Go

Este projeto demonstra como validar estruturas (structs) no Go de forma modular e reutilizável, utilizando um sistema de regras customizáveis. Ele aborda exemplos para validação de campos simples, estruturas aninhadas e regras personalizadas.

---

## Exemplo: Validação Simples

Esse exemplo valida os campos da struct `Product` e aplica todas as validações ao campos `Name` e `Price`.

```go
package main

import (
	"fmt"

	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
)

func main() {
	// Representa a struct Product a ser validada
	product := struct {
		Name  string
		Price float64
	}{Name: "Phone", Price: -1.999}

	// Cria uma instância do validator
	v := validator.New()
	// Define as regras de validação para cada campo
	v.Add("Name", rule.Rules{rule.Required()})
	v.Add("Price", rule.Rules{
		rule.Required(),
		rule.MinValue(0.0),
	})

	// Valida os campos da struct
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
```

## Exemplo: Validação no Momento da Criação

Este exemplo valida os dados ao criar uma nova instância de `User`:

```go
package main

import (
	"fmt"

	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

// User representa uma entidade com os campos Name e Email.
type User struct {
	Name  string
	Email string
}

// New cria uma nova instância de User com validação automática.
func New(name, email string) (*User, []*validator_error.ValidatorError) {
	user := &User{name, email}

	// Valida os dados do usuário antes de criar a instância.
	if errs, ok := user.Validate(); !ok {
		return nil, errs
	}

	return user, nil
}

// Validate aplica as regras de validação nos campos da struct User.
func (u *User) Validate() ([]*validator_error.ValidatorError, bool) {
	v := validator.New()

	// Regras de validação para o campo Name.
	v.Add("Name", rule.Rules{
		rule.Required(),
		rule.MinLength(3),
		rule.MaxLength(50),
	})

	// Regras de validação para o campo Email.
	v.Add("Email", rule.Rules{
		rule.Required(),
		rule.ValidEmail(),
	})

	// Executa a validação.
	errs := v.Validate(*u)
	if len(errs) == 0 {
		return nil, true
	}
	return errs, false
}

func main() {
	// Testa a criação de um usuário com dados inválidos.
	user, errs := New("Jo", "invalid_email")
	if errs != nil {
		fmt.Println("Erros de validação:")
		for _, err := range errs {
			fmt.Printf("Campo: %s, Mensagem: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Usuário válido: ", user)
	}
}
```

---

## Exemplo: Validação de Estruturas Aninhadas

Para cenários onde uma struct contém outra struct, as validações também podem ser aplicadas:

```go
package main

import (
	"fmt"

	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

// Address representa o endereço do usuário.
type Address struct {
	Street string
	Zip    string
}

// User representa um usuário com endereço.
type User struct {
	Name    string
	Email   string
	Address Address
}

// New cria uma nova instância de User com validação automática, incluindo Address.
func New(name, email, street, zip string) (*User, []*validator_error.ValidatorError) {
	user := &User{
		Name:  name,
		Email: email,
		Address: Address{
			Street: street,
			Zip:    zip,
		},
	}

	// Valida os dados do usuário e do endereço.
	if errs, ok := user.Validate(); !ok {
		return nil, errs
	}

	return user, nil
}

// Validate aplica as regras de validação para User e Address.
func (u *User) Validate() ([]*validator_error.ValidatorError, bool) {
	v := validator.New()

	// Regras para User.
	v.Add("Name", rule.Rules{
		rule.Required(),
		rule.MinLength(3),
	})
	v.Add("Email", rule.Rules{
		rule.Required(),
		rule.ValidEmail(),
	})

	// Regras para Address.
	v.Add("Address.Street", rule.Rules{
		rule.Required(),
	})
	v.Add("Address.Zip", rule.Rules{
		rule.Required(),
		rule.ExactLength(5),
	})

	// Executa a validação.
	errs := v.Validate(*u)
	if len(errs) == 0 {
		return nil, true
	}
	return errs, false
}

func main() {
	// Testa a criação de um usuário com dados inválidos.
	user, errs := New("Jo", "invalid_email", "", "")
	if errs != nil {
		fmt.Println("Erros de validação:")
		for _, err := range errs {
			fmt.Printf("Campo: %s, Mensagem: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Usuário válido: ", user)
	}
}
```

---

## Exemplo: Validação Personalizada

Implemente regras personalizadas para campos específicos:

```go
package main

import (
	"fmt"

	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

// User representa um usuário com nome, email e idade.
type User struct {
	Name  string
	Email string
	Age   int
}

// New cria uma nova instância de User com validação personalizada.
func New(name, email string, age int) (*User, []*validator_error.ValidatorError) {
	user := &User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	// Valida os dados do usuário antes de retornar.
	if errs, ok := user.Validate(); !ok {
		return nil, errs
	}

	return user, nil
}

// Validate aplica regras de validação, incluindo regras personalizadas.
func (u *User) Validate() ([]*validator_error.ValidatorError, bool) {
	v := validator.New()

	// Regras personalizadas para o campo Age.
	v.Add("Age", rule.Rules{
		func(key string, value interface{}) *validator_error.ValidatorError {
			// Valida se o número é par.
			if value.(int)%2 != 0 {
				return &validator_error.ValidatorError{
					Field:   key,
					Message: "the number must be even",
				}
			}
			return nil
		},
	})

	// Executa a validação.
	errs := v.Validate(*u)
	if len(errs) == 0 {
		return nil, true
	}
	return errs, false
}

func main() {
	// Testa a criação de um usuário com dados inválidos.
	user, errs := New("Jo", "invalid_email", 25)
	if errs != nil {
		fmt.Println("Erros de validação:")
		for _, err := range errs {
			fmt.Printf("Campo: %s, Mensagem: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Usuário válido: ", user)
	}
}
```
