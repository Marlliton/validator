# Guia de Contribuição

Obrigado por considerar contribuir para o projeto **Validator**! Este guia irá ajudá-lo a colaborar de maneira eficiente e produtiva.

## 📋 Como Contribuir

### Relatando Problemas (Issues)
Se você encontrou um bug, tem uma dúvida ou gostaria de sugerir uma nova funcionalidade:
1. Verifique se o problema já foi relatado navegando nos [issues existentes](https://github.com/Marlliton/validator/issues).
2. Abra um novo issue com o máximo de detalhes possível, incluindo:
   - Descrição clara do problema.
   - Passos para reproduzir o erro, se aplicável.
   - Ambiente (versão do Go, sistema operacional, etc.).
   - Mensagens de erro relevantes ou logs.

### Enviando Pull Requests (PRs)
Se você deseja corrigir um problema ou implementar uma nova funcionalidade:
1. Faça um fork do repositório.
2. Crie um novo branch com um nome descritivo:
   ```bash
   git checkout -b fix/validacao-email
   ```
3. Faça suas alterações e adicione commits significativos:
   - Escreva mensagens de commit claras e descritivas.
4. Certifique-se de que os testes estejam passando:
   ```bash
   go test ./...
   ```
5. Abra um pull request para o branch `main`, explicando:
   - O problema que você está resolvendo.
   - Uma descrição das mudanças feitas.
   - Qualquer contexto adicional que ajude na revisão.

## 🛠️ Configurando o Ambiente de Desenvolvimento

1. Clone o repositório:
   ```bash
   git clone https://github.com/Marlliton/validator.git
   cd validator
   ```
2. Certifique-se de ter o Go instalado (versão mínima recomendada: (1.23.4).
3. Instale dependências, se necessário:
   ```bash
   go mod tidy
   ```
4. Rode os testes para garantir que tudo está funcionando:
   ```bash
   go test ./...
   ```

## 🧪 Adicionando ou Atualizando Testes

- **Todo código novo deve incluir testes apropriados.** Isso garante que as alterações sejam verificáveis e que o projeto mantenha sua qualidade.
- Os testes devem ser colocados no diretório do arquivo que está sendo testado segudo de `*_test.go`.
- Siga o padrão de nomenclatura e organização dos testes existentes.

## 📚 Estilo de Código

- Este projeto segue as práticas recomendadas da comunidade Go (verifique o [Effective Go](https://go.dev/doc/effective_go)).
- Use o `gofmt` para formatar seu código:
  ```bash
  gofmt -s -w .
  ```
- Nomeie variáveis, funções e métodos de forma clara e descritiva.

## ✨ Boas Práticas

- Mantenha os pull requests pequenos e focados. Grandes alterações são mais difíceis de revisar.
- Explique suas mudanças claramente no PR.
- Sempre atualize seu branch com as últimas alterações do branch principal:
  ```bash
  git pull origin main
  ```


## 💬 Dúvidas?

Se você tiver alguma dúvida sobre como contribuir, abra um issue ou entre em contato diretamente pelo repositório.
