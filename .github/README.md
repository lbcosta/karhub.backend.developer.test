<p align="center">
  <a href="https://github.com/lbcosta/karhub.backend.developer.test">
      <img height="125" alt="Karhub" src="logo.png">
  </a>
  <br>
</p>

<div align="center">
  <a href="https://github.com/Azure/gocover">
    <img src="https://img.shields.io/badge/%F0%9F%94%8Egocover-75.4%25-green">
  </a>
  <a href="https://github.com/securego/gosec">
    <img src="https://img.shields.io/badge/%F0%9F%94%91gosec-passing-green">
  </a>
  <a href="#-testes">
    <img src="https://img.shields.io/badge/%F0%9F%A7%AA%20tests-passing-green">
  </a>
    <a href="#-documentação-da-api">
    <img src="https://img.shields.io/badge/%F0%9F%93%83%20API-docs-informational">
  </a>
</div>

<br>

<p align="center">
  Esta é uma API feita em <a href="https://go.dev/">Go</a> como teste técnico para a posição de Desenvolvedor Pleno na <a href="https://www.karhub.com.br/">Karhub</a>. A aplicação se trata de uma API REST para cadastrar, listar, atualizar e deletar estilos de cervejas com suas devidas temperaturas ideais de consumo. Além disso, a aplicação possui integração com a API do Spotify para buscar as músicas baseadas em um determinado estilo de cerveja. A aplicação possui um banco de dados em <a href="https://www.postgresql.org/">Postgresql</a>, testes unitários e de integração com <a href="https://github.com/stretchr/testify">Testify</a> e é containerizada com <a href="https://www.docker.com/">Docker</a> e <a href="https://docs.docker.com/compose/">Docker Compose</a>.
</p>

<p align="center">
  <a href="#%EF%B8%8F-instalação">Instalação</a> •
  <a href="#%EF%B8%8F-inicialização">Inicialização</a> •
  <a href="#-testes">Testes</a> •
  <a href="#-seeding">Seeding</a> •
  <a href="#-documentação-da-api">Documentação da API</a>
</p>

# ⚙️ Instalação

**É necessário ter Docker, Docker Compose e Make instalados na sua máquina.**

Clone o projeto para sua máquina:

```bash
git clone https://github.com/lbcosta/karhub.backend.developer.test.git
```

Na raíz do projeto, crie um arquivo `.env` com os seguintes valores, substituindo os valores entre `< >` pelas suas credenciais do Spotify:

```
APP_ENV=development
APP_PORT=8080

SPOTIFY_CLIENT_ID=<client-id>
SPOTIFY_CLIENT_SECRET=<client-secret>

POSTGRES_HOST=beer_db
POSTGRES_HOST_SEED=localhost
POSTGRES_PORT=5432
POSTGRES_USER=karhub
POSTGRES_PASSWORD=karhub.b33r
POSTGRES_DB=karhub
```
<br>

> 💡 **Observação:** Make não é um requisito obrigatório, mas é recomendado para facilitar a execução dos comandos.
> Caso não queira utilizar o Make, basta executar os comandos presentes no arquivo `Makefile` manualmente.

# ⚡️ Inicialização

Basta estar na raíz do projeto e executar o seguinte comando no terminal:

```bash
make
```

Para parar a aplicação e todos os containers Docker e removê-los do sistema, execute:

```bash
make cleanup
```

# 🧪 Testes

Para executar os testes, execute na raíz do projeto:
```bash
make test
```

# 🌱 Seeding

Para popular o banco de dados com dados de exemplo, execute na raíz do projeto:
```bash
make seed
```

# 📃 Documentação da API

URL Base:

```
http://localhost:8080/api/v1
```

## 🔗 Endpoints

### 🔍 Listagem de Estilos de Cervejas

**URL**: `/beer` <br>
**Method**: GET<br>
**Request Body**: _Sem request_ <br>
**Response**: Lista de estilos cervejas disponíveis<br>

Exemplo de **Response**:

<img src="https://img.shields.io/badge/Status-200-green">

```json
[
  {
    "id": 21,
    "style": "Weissbier",
    "min_temperature": -1,
    "max_temperature": 3
  },
  {
    "id": 22,
    "style": "Pilsens",
    "min_temperature": -2,
    "max_temperature": 4
  },
  {
    "id": 23,
    "style": "Weizenbier",
    "min_temperature": -4,
    "max_temperature": 6
  }
]
```

Exemplos de possíveis **Erros**:

<img src="https://img.shields.io/badge/Status-422-red">

```
failed to get all beers
```

### ✏️ Registro de um Estilo de Cerveja

**URL**: `/beer` <br>
**Method**: POST<br>
**Request Body**: Objeto JSON com informações da cerveja<br>
**Response**: Objeto criado<br>

Exemplo de **Request Body**:

```json
{
  "style": "IPA",
  "min_temperature": -7,
  "max_temperature": 10
}
```

Exemplo de **Response**:

<img src="https://img.shields.io/badge/Status-200-green">

```json
{
  "id": 31,
  "style": "IPA",
  "min_temperature": -7,
  "max_temperature": 10
}
```

Exemplos de possíveis **Erros**:

<img src="https://img.shields.io/badge/Status-400-red">

```
failed to create beer: min temperature is higher than max temperature
```

### 🔄 Edição de Cervejas

**URL**: `/beer/:id` <br>
**Paramêtros**: `id` - Id da cerveja a ser editada<br>
**Method**: PATCH<br>
**Request Body**: Objeto JSON com as informações a serem editadas<br>
**Response**: Objeto após ser editado<br>

Exemplo de **Request Body**:

```json
{
  "max_temperature": 5
}
```

Exemplo de **Response**:

<img src="https://img.shields.io/badge/Status-200-green">

```json
{
  "id": 22,
  "style": "Pilsens",
  "min_temperature": -2,
  "max_temperature": 5
}
```

Exemplos de possíveis **Erros**:

<img src="https://img.shields.io/badge/Status-400-red">

```
min temperature is higher or equal than max temperature
```

### 🗑️ Exclusão de Cervejas

**URL**: `/beer/:id` <br>
**Paramêtros**: `id` - Id da cerveja a ser excluída<br>
**Method**: DELETE <br>
**Request Body**: _Sem request_ <br>
**Response**: _Sem response - Status 204: No Content_ <br>


Exemplos de possíveis **Erros**:

<img src="https://img.shields.io/badge/Status-404-red">

```
beer of given id not found
```

### 🌡️ Obter um Estilo de Cerveja a partir de uma Temperatura

**URL**: `/beer/style` <br>
**Method**: GET<br>
**Request Body**: Objeto contendo a temperatura <br>
**Response**: Lista de cervejas e playlists relacionadas às cervejas <br>

Exemplo de **Request Body**:

```json
{
  "temperature": 3
}
```

Exemplo de **Response**:

<img src="https://img.shields.io/badge/Status-200-green">

```json
[
  {
    "beer_style": "Pilsens",
    "playlist": {
      "name": "pilsen’s ",
      "tracks": [
        {
          "name": "What's Luv? (feat. Ja-Rule & Ashanti)",
          "artist": "Fat Joe, Ja Rule, Ashanti",
          "link": "https://open.spotify.com/track/2mKouqwAIdQnMP43zxR89r"
        },
        {
          "name": "Always On Time",
          "artist": "Ja Rule, Ashanti",
          "link": "https://open.spotify.com/track/4hrae8atte6cRlSC9a7VCO"
        }
      ]
    }
  },
  {
    "beer_style": "IPA",
    "playlist": {
      "name": "Ipank full album",
      "tracks": [
        {
          "name": "Terlalu Sadis",
          "artist": "Ipank",
          "link": "https://open.spotify.com/track/3Rz2UQ9IwkhDDqmGhT7wXX"
        },
        {
          "name": "Gubuk Jadi Istana",
          "artist": "Ipank",
          "link": "https://open.spotify.com/track/1qJAh2JBuFF90os0MJwqbF"
        }
      ]
    }
  }
]
      
```

### 💓 Health Check

**URL**: `/health` <br>
**Method**: GET<br>
**Response**: Informações sobre o funcionamento da aplicação<br>

Exemplo de **Response**:

<img src="https://img.shields.io/badge/Status-200-green">

```json{
    "status": "OK",
    "version": "1.0.0",
    "uptime": "13.271318371s",
    "database_status": "OK"
}
```