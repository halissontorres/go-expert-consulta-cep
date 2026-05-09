# **GO EXPERT** | [FULL CYCLE](https://goexpert.fullcycle.com.br) | 2026 <img align="absmiddle" src="./assets/go-logo.png" width="100" height="100" alt="Go logo"/>


---
## Go expert | desafio técnico | construir um sistema que busca o endereço mais rápido entre duas APIs de CEP distintas


Principais diretórios:

- `cmd/main.go`: ponto de entrada principal da aplicação.
- `pkg/brasilapi`: cliente e modelo de resposta da BrasilAPI.
- `pkg/viacep`: cliente e modelo de resposta da ViaCEP.
- `pkg/api`: estrutura comum de resposta usada na concorrência.
- `configs/config`: configurações globais, incluindo timeout padrão.
- `configs/log`: logger utilizado pela aplicação.

---

## Como executar

### 1. Clone o repositório

```bash
git clone https://github.com/halissontorres/go-expert-consulta-cep
```

### 2. Execute a aplicação

```bash
go run cmd/main.go
```

> Por padrão, a aplicação consulta o CEP 13330-250 da Full Cycle :D

A saída exibirá os dados retornados pela API que responder primeiro.

Exemplo de saída:

```terminaloutput
INFO: 2026/05/08 12:19:49 [BrasilApi] Cep: 13330250, State: SP, City: Indaiatuba, Neighborhood: Centro, Street: Rua Hércules Mazzoni, Service: open-cep
```

## Configurando o timeout

O timeout padrão da aplicação é de `1 segundo`, conforme exigido pelo desafio.

Também é possível sobrescrever esse valor usando a variável de ambiente `CONSULTA_CEP_TIMEOUT`.

Exemplo:

```bash
CONSULTA_CEP_TIMEOUT=5s
go run cmd/main.go
```

Se nenhuma API responder dentro do tempo configurado, será exibida uma mensagem semelhante a:

```terminaloutput
[WARN] Tempo excedido: sem resposta dos serviços
```

## Como testar

Para executar todos os testes do projeto:

```bash
go test -v ./...
```

## Funcionamento esperado

A aplicação deve:

1. Iniciar duas requisições HTTP em paralelo:
   - uma para a BrasilAPI;
   - outra para a ViaCEP.

2. Receber a primeira resposta disponível por meio de um channel.

3. Exibir no terminal:
   - o provedor que respondeu primeiro.
   - os dados do endereço;

4. Encerrar a espera caso o tempo limite seja atingido.

---

## APIs consultadas

- [BrasilAPI](https://brasilapi.com.br/api/cep/v1/{cep})
- [ViaCEP](https://viacep.com.br/ws/{cep}/json/)

## Requisitos do desafio

- Fazer requisições simultâneas para as duas APIs.
- Aceitar apenas a resposta mais rápida.
- Descartar a resposta mais lenta.
- Exibir os dados do endereço no terminal.
- Exibir qual API respondeu primeiro.
- Usar timeout de `1 segundo`.
- Exibir erro caso o timeout seja atingido.
