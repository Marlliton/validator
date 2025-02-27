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
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u *User) Validate() (bool, []*fail.Error) {
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
		return true, nil
	}
	return false, errs
}

func main() {
	user := User{Name: "Jo", Email: "invalid_email", Age: 15}
	if ok, errs := user.Validate(); !ok {
		fmt.Println("Erros de validação:")
		for _, err := range errs {
			fmt.Println(err.Message)
		}
	} else {
		fmt.Println("Tudo válido!")
	}
}
```

### Regras de Validação Suportadas

O pacote oferece uma variedade de regras que podem ser aplicadas aos campos. Exemplos:

- `rule.Required()`: O campo é obrigatório.
- `rule.MinLength(n)`: Comprimento mínimo.
- `rule.MaxLength(n)`: Comprimento máximo.
- `rule.ExactLength(n)`: Comprimento exatamente igual a "n".
- `rule.ValidURL()`: Valida uma URL.
- `rule.ValidEmail()`: Valida um email.
- `rule.ValidPhoneNumber()`: Valida um telefone no formato [e164](https://en.wikipedia.org/wiki/E.164).
- `rule.String(n)`: Verifica se é uma string.
- `rule.Int(n)`: Verifica se é um inteiro.
- `rule.Bool(n)`: Verifica se é um bool.
- `rule.MinValue(n)`: Valor mínimo permitido.
- `rule.MaxValue(n)`: Valor máximo permitido.

Você pode combinar essas regras em diferentes campos.

### Validações Personalizadas

Crie suas próprias validações passando uma função personalizada como regra:

```go
v.Add("CustomField", rule.Rules{
	func(fieldName string, value interface{}) *fail.Error {
		if value.(int)%2 != 0 {
			return fail.New(fieldName, "O número deve ser par")
		}
		return nil
	},
})
```

## 📚 Exemplos

- Exemplos de uso disponíveis no diretório [`examples`](./examples).

## 🛡️ Contribuindo

Contribuições são bem-vindas! Abra um issue para relatar problemas ou envie um pull request com melhorias. Certifique-se de seguir as diretrizes no arquivo [CONTRIBUTING.md](./CONTRIBUTING.md).

## 🗹 Itens a Fazer

- [ ] Permitir que o usuário forneça sua próprias mensagens customizadas
- [x] Validar uma URL
- [ ] Checar o tipo:
  - [x] string
  - [x] number
    - [x] bool
    - [ ] float
- [ ] Adicionar método in para validar se um dado esté presente em um slice (includes)
- [x] Resolver bug na regra Required com o valor zero (0)

## 📄 Licença

Este projeto é distribuído sob a licença [MIT](./LICENSE).
