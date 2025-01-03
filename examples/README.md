## Validação no Momento da Criação

Este exemplo mostra como usar uma função de fábrica `New` para criar uma instância de `User` com validação automática. Ele verifica se os dados fornecidos atendem às regras antes de retornar o objeto, garantindo consistência e segurança nos dados. Caso os dados sejam inválidos, uma lista de erros é retornada.

```go
type User struct {
	Name  string
	Email string
}

func New(name, email string) (*User, []*validator_error.ValidatorError) {
	user := &User{name, email}

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
		rule.MaxLength(50),
	})
	v.Add("Email", rule.Rules{
		rule.Required(),
		rule.ValidEmail(),
	})

	errs := v.Validate(*u)
	if len(errs) == 0 {
		return nil, true
	}
	return errs, false
}

func main() {
	user, errs := New("Jo", "invalid_email")
	if errs != nil {
		fmt.Println("Erros de validação:")
		for _, err := range errs {
			fmt.Printf("Campo: %s, Mensagem: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Tudo válido: ", user)
	}
}
```

### Validação de Estruturas Aninhadas

Este exemplo demonstra como validar uma estrutura principal (`User`) que contém outra estrutura aninhada (`Address`).
Regras de validação são aplicadas tanto aos campos de `User` quanto aos de `Address`, garantindo que todos os dados estejam corretos antes de criar a instância.
Campos como `Name`, `Email`, `Street` e `Zip` são validados com regras específicas.

```go
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
```

### Validação Personalizada

Este exemplo mostra como adicionar uma regra de validação personalizada para um campo. A validação verifica se o valor do campo é um número par. Se não for, um erro de validação é retornado, com a mensagem "the number must be even". Caso contrário, a validação é considerada bem-sucedida.

```go
type User struct {
	Name  string
	Email string
	Age   int
}

func New(name, email string, age int) (*User, []*validator_error.ValidatorError) {
	user := &User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	if errs, ok := user.Validate(); !ok {
		return nil, errs
	}

	return user, nil
}

func (u *User) Validate() ([]*validator_error.ValidatorError, bool) {
	v := validator.New()

	v.Add("Age", rule.Rules{
		func(key string, value interface{}) *validator_error.ValidatorError {
			if value.(int)%2 != 0 {
				return &validator_error.ValidatorError{
					Field:   key,
					Message: "the number must be even",
				}
			}
			return nil
		},
	})

	errs := v.Validate(*u)
	if len(errs) == 0 {
		return nil, true
	}
	return errs, false
}
```
