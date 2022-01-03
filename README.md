# Desafio Lógico Letras 2022

Para esse desafio, o projeto usou Golang, Docker e Docker Compose, PostgreSQL e Unit tests todos seguindo os conceitos de Clean Architecture para desenvolvimento de software. 

O Docker Compose é usado para automatizar a criação do container usando Docker que cria o banco de dados em PostgreSQL em que está salvo os títulos da músicas usadas para calcular o score.


## Instalação dos requisitos:

Primeiramente, devemos garantir que seu sistema operacional tenha a versão mais recente do Golang. Para isso basta seguir os passos de instalação em https://go.dev/doc/install. Em seguida, devemos instalar o Docker para executar o container que implementa o banco de dados em PostgreSQL.

Para instalar o Docker  siga as instruções dependendo do seu sistema operacional: 
- Ubuntu: https://docs.docker.com/engine/install/ubuntu/
- Windows: https://docs.docker.com/desktop/windows/install/
- MacOS: https://docs.docker.com/desktop/mac/install/
- Usando binários estáticos: https://docs.docker.com/engine/install/binaries/

Por último, é preciso instalar o Docker Compose. Siga os passos em https://docs.docker.com/compose/install/ dependendo do seu sistema operaciona. 

Depois desses dois primeiros passos o seu sistema está apto a executar o projeto. 

## Execucação do Projeto

Faça o clone do projeto primeiramente
```
git clone https://github.com/pedrobemer/desafio_letras2021.git
```

Em seguida, entre no diretório `desafio_letras2021`:
```
cd desafio_letras2021
```

Estando no diretório do projeto, é preciso primeiramente criar dois arquivos de ambiente. O primeiro é `.env`. Esse arquivo terá a variável de ambiente que armazenará a senha de acesso ao banco de dados. Assim, o conteúdo desse banco de dados é:
```
POSTGRES_PASSWORD_TEST=<PASSWORD>
```

O outro arquivo tem as variáveis de ambiente necessárias para a conexão com o banco de dados usando nosso projeto em Golang. O arquivo tem o nome de `database.env` e o conteúdo desse arquivo é:
```
DB_USER=postgres
DB_PASSWORD=test
DB_NAME=letras
DB_PORT=8082
DB_HOST=localhost
```

Estando no diretório do projeto, execute o comando para gerar o arquivo binário:
```
go build -o desafio-letras
```

Execute o comando para levantar o container do banco de dados:
```
docker-compose up
```

Para executar os unit tests faça:
```
go test -v ./...
```


Desse modo é possível executar o binário:
```
./desafio-letras
```

Com esses passos é possivel testar a lógica do cálculo do score com base na entrada padrão.


## Organização do Software

Para esse projeto se usou os conceitos de Clean Architecture, onde temos uma hierarquia baseada em camadas com a camada mais interna sendo a entities seguida pela de use cases, controllers/presenters e database. Nesse conceito, as camadas mais externas podem usar as mais internas, mas o mesmo não é válido para as mais internas usarem as mais externas. Assim temos nesse projeto:
- Entities: Representada pelo diretório `entity` onde tem as estruturas bases que serão usadas por todas as camadas.
- Use Cases: Representada pelo diretório `usecases` que tem toda a lógica programável do projeto. É onde tem o cálculo de score e do retorno dos melhore títulos com base no socre.
- Controllers/Presenters: Permite que os Use Cases consigam obter os dados do banco de dados, por exemplo sem desobedecer a arquitetura. No caso do uso de Golang, não temos diretamente um diretório específico para essa camada, mas podemos representar essa camada por meio do conceito de interfaces. Assim, no diretório do music temos um arquivo chamada `interface.go` que cria uma interface que podemos usar para acessar o banco de dados sem ser de forma direta. A vantagem é que não desobedecemos os conceitos de Clean Architecture e permite uma flexibilidade do uso do banco de dados já que o uso de interface apenas obriga que as funções tenham os mesmos parametros de entrada e saída sem se importar com o que é feito internamente na função. Assim podemos usar qualquer banco de dados seja ele relacional ou não relacional.
- Database: A camada mais externa que é representada pelo diretório `database/postgresql`.


## Estrutura de Dados

Para armazenar as informações dos 10 melhores scores como proposto no desafio, o projeto usou um slice de estruturas de 11 elementos. Essa estrutura tem a forma de:
```
type MusicInfo struct {
	Name  string
	Score int
}

```

A intenção de ter 11 elementos é que toda vez que um novo score é calculado para o dado "Name", a lógica salva na 11 posição do slice, já que para o usuário final é apenas importante que se retorne as 10 primeiras posições que refletem os maiores scores. Assim, a 11 posição é totalmente descartável. Logo em seguida, a lógica usa a implemetação otimizada de sort do próprio Golang (para entender melhor veja o arquivo em `usecases/music/application.go`) para que temos os 10 maiores score e que quando os scores forem iguais a prioridade vá para aquele que esteja na frente de forma alfabética. Com isso, depois do sort sempre garantimos que a 11 posição do slice é que tem o menor score e consequentemente pode ser descartada.

