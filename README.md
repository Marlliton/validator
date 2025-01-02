# Validator

Uma biblioteca de validação para structs no Go, que utiliza regras definidas programaticamente, sem a necessidade de tags. Este pacote é ideal para quem busca flexibilidade e reutilização de lógica de validação em diversos contextos de uma aplicação.

## 📋 Funcionalidades

- Validação baseada em regras programáticas (sem tags).
- Suporte a validações customizadas.
- Mensagens de erro associadas aos campos.
- Compatível com projetos complexos e escaláveis.

## 🛠️ Instalação

Instale o pacote usando:

```bash
go get github.com/Marlliton/validator
```

## 🚀 Uso Básico

Veja um exemplo de como criar e validar uma struct usando o pacote:

### Definição e Validação

```go
package main

import (
	"fmt"

	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rules"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u *User) Validate() ([]*validator_error.ValidatorError, bool) {
	v := validator.New()

	v.Add("Name", rules.Rules{
		rules.Required(),
		rules.MinLength(3),
		rules.MaxLength(50),
	})
	v.Add("Email", rules.Rules{
		rules.Required(),
		rules.ValidEmail(),
	})
	v.Add("Age", rules.Rules{
		rules.MinValue(18),
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
```

### Regras de Validação Suportadas

O pacote oferece uma variedade de regras que podem ser aplicadas aos campos. Exemplos:

- `rules.Required()`: O campo é obrigatório.
- `rules.MinLength(n)`: Comprimento mínimo.
- `rules.MaxLength(n)`: Comprimento máximo.
- `rules.ValidEmail()`: Valida um email.
- `rules.MinValue(n)`: Valor mínimo permitido.
- `rules.MaxValue(n)`: Valor máximo permitido.

Você pode combinar essas regras em diferentes campos.

### Validações Personalizadas

Crie suas próprias validações passando uma função personalizada como regra:

```go
v.Add("CustomField", rules.Rules{
	func(fieldName string, value interface{}) *validator_error.ValidatorError {
		if value.(int)%2 != 0 {
			return &validator_error.ValidatorError{
				Field:   fieldName,
				Message: "O número deve ser par",
			}
		}
		return nil
	},
})
```

## ✅ Testes

Certifique-se de que tudo está funcionando corretamente rodando:

```bash
go test ./...
```

## 📚 Exemplos

- Exemplos de uso disponíveis no diretório [`examples`](./examples).

## 🛡️ Contribuindo

Contribuições são bem-vindas! Abra um issue para relatar problemas ou envie um pull request com melhorias. Certifique-se de seguir as diretrizes no arquivo [CONTRIBUTING.md](./CONTRIBUTING.md).

## 📄 Licença

Este projeto é distribuído sob a licença [MIT](./LICENSE).

---

Esse README deixa claro como usar o pacote, fornece exemplos práticos e destaca as funcionalidades disponíveis. Você pode incluir mais detalhes, como links para benchmarks ou exemplos avançados, conforme necessário.
