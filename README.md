# MyWay Service
## Escopo
O serviço deve avaliar a viagem do solicitante, buscando o caminho mais barato sem se preocupar com o tamanho do trajeto.

## Definições do projeto
- A linguagem escolhida foi a Golang;
- Bibliotecas de suporte para facilitar a estruturação do projeto:
  - [Cobra](https://github.com/spf13/cobra) - para interação via CLI com a aplicação;
  - [Gorilla Mux](https://github.com/gorilla/mux) - um poderoso router para facilitar a interação via serviço HTTP;
  - [Mock](https://github.com/golang/mock) - pacote para geração e trabalho com Mocks
  - [Testify](https://github.com/stretchr/testify) - pacote de asserções e mocks para testes
- Arquitetura:
  - Estruturada em cima de conceitos da [Arquitetura Hexagonal](https://alistair.cockburn.us/hexagonal-architecture/) com objetivo de isolar o núcleo de negócio da aplicação das camadas externas (portas e adaptadores). Com isso podemos garantir um baixo acoplamento e uma boa coesão.
  - Princípios do Design Pattern SOLID também foram observados durante a implementação (segregação de interfaces, srp, lsp)
  - Adaptadores:
    - cli - define a forma de interação pela linha de comando
    - csv - define como o core fará persistências no arquivo fonte dos dados
    - dto - estrutura de comunicação (dado) entre serviço e aplicação
    - web - define a forma de interação pelo serviço web (API Rest)

## Rodando a aplicação

### Instalação
1. Clone o repositório:
```shell
git clone https://github.com/alissonphp/go-routes-csv
```
2. Sincronize os pacotes de dependências
```shell
go mod tidy
```

3. Caso prefira, você pode utilizar o docker para rodar o projeto
```shell
docker compose up -d
docker exec -it myway-app bash
```

### CLI
Basta chamar a aplicação pelo comando ``go run main.go input-routes.csv GRU-CDG`` seguindo a seguintes estrutura ``go run main.go [file.csv] FROM-TO``

O retorno esperado é o seguinte:

```shell
root@ce48214b1e7a:/go/src# go run main.go input-routes.csv GRU-CDG
best route: GRU-BRC-SCL-ORL-CDG > $40
```

### API Rest

O serviço web pode ser inicilizado pelo comando ``go run main.go http`` onde será servido na porta **:8080** por padrão (o docker-compose já está fazendo bind da porta caso queria acessar pelo host).

#### Endpoints:
- **[POST] /routes** - cria uma nova rota (persistindo no arquivo .csv)
```shell
curl --request POST \
  --url http://localhost:8080/routes \
  --header 'Content-Type: application/json' \
  --data '{
	"from": "FOR",
	"to": "BSB",
	"price": 80
}'
```

```json
{
      "from": "FOR",
      "to": "BSB",
      "price": 80
}
```

- **[GET] /routes/cheapest?from=FROM&to=TO** - busca a rota mais barata de acordo com o ponto de partida e de destino do cliente
```shell
curl --request GET \
  --url 'http://localhost:8080/routes/cheapest?from=GRU&to=CDG'
```

A resposta deve ser um objeto como o seguinte:

```json
{
	"fly_path": "GRU-BRC-SCL-ORL-CDG",
	"total_cost": 40
}
```

É possível testar as rotas utilizando algum cliente como o Insomnia ou Postman. Basta importar o [arquivo de coleções](./docs/Insomnia_2022-08-04.json) que está disponível no repositório.