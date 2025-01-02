## Validação no Momento da Criação

Este exemplo mostra como usar uma função de fábrica (New) para criar uma instância de User com validação automática. Ele verifica se os dados fornecidos atendem às regras antes de retornar o objeto, garantindo consistência e segurança nos dados. Caso os dados sejam inválidos, uma lista de erros é retornada.

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
