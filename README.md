# Documentação da API

## Visão Geral

Esta API foi desenvolvida utilizando o framework Gin para Go e o GORM como ORM para interagir com um banco de dados SQLite. A API permite a criação e listagem de posts e comentários.

### Endpoints

- **POST /posts**: Cria um novo post.
- **GET /posts**: Lista todos os posts, incluindo seus comentários.
- **POST /comments**: Cria um novo comentário para um post específico.

---

## Endpoints Detalhados

### 1. Criar um Post

- **Método:** `POST`
- **Endpoint:** `/posts`
- **Descrição:** Cria um novo post.
- **Exemplo de Requisição:**

  ```bash
  curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title": "Meu Primeiro Post", "content": "Este é o conteúdo do meu primeiro post"}'

    Corpo da Requisição:

json

    {
      "title": "Meu Primeiro Post",
      "content": "Este é o conteúdo do meu primeiro post"
    }

Exemplo de Resposta:

    Status: 200 OK

json

    {
      "ID": 1,
      "CreatedAt": "2024-08-27T12:34:56.789Z",
      "UpdatedAt": "2024-08-27T12:34:56.789Z",
      "DeletedAt": null,
      "title": "Meu Primeiro Post",
      "content": "Este é o conteúdo do meu primeiro post",
      "comments": []
    }

2. Listar todos os Posts

    Método: GET

    Endpoint: /posts

    Descrição: Retorna todos os posts com seus respectivos comentários.

    Exemplo de Requisição:

    ```bash

    curl -X GET http://localhost:8080/posts

Exemplo de Resposta:

    Status: 200 OK

json

    [
      {
        "ID": 1,
        "CreatedAt": "2024-08-27T12:34:56.789Z",
        "UpdatedAt": "2024-08-27T12:34:56.789Z",
        "DeletedAt": null,
        "title": "Meu Primeiro Post",
        "content": "Este é o conteúdo do meu primeiro post",
        "comments": [
          {
            "ID": 1,
            "PostID": 1,
            "Content": "Este é um comentário",
            "CreatedAt": "2024-08-27T12:34:56.789Z",
            "UpdatedAt": "2024-08-27T12:34:56.789Z",
            "DeletedAt": null
          }
        ]
      }
    ]

3. Criar um Comentário

    Método: POST

    Endpoint: /comments

    Descrição: Cria um novo comentário para um post existente.

    Exemplo de Requisição:

bash

    curl -X POST http://localhost:8080/comments \
    -H "Content-Type: application/json" \
    -d '{"post_id": 1, "content": "Este é um comentário"}'

Corpo da Requisição:

json

    {
      "post_id": 1,
      "content": "Este é um comentário"
    }

Exemplo de Resposta:

    Status: 200 OK

json

    {
      "ID": 1,
      "PostID": 1,
      "Content": "Este é um comentário",
      "CreatedAt": "2024-08-27T12:34:56.789Z",
      "UpdatedAt": "2024-08-27T12:34:56.789Z",
      "DeletedAt": null
    }

Modelo de Dados
Post

    ID: Identificador único do post (uint)
    Title: Título do post (string)
    Content: Conteúdo do post (string)
    Comments: Lista de comentários associados ao post ([]Comment)

Comment

    ID: Identificador único do comentário (uint)
    PostID: Identificador do post ao qual o comentário pertence (uint)
    Content: Conteúdo do comentário (string)

Configuração do Banco de Dados

A API utiliza o SQLite como banco de dados. A conexão é estabelecida no arquivo database/database.go e as tabelas são automaticamente migradas utilizando GORM.

go

    # Código Go para conectar ao banco de dados e migrar os modelos
    
    func ConnectDatabase() {
        database, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
        if err != nil {
            panic("Failed to connect to database!")
        }
    
        database.AutoMigrate(&models.Post{}, &models.Comment{})
        DB = database
    }

Observações

    CORS: A API permite requisições de qualquer origem através de um middleware CORS. 
    Esse middleware permite métodos GET, POST e OPTIONS.
    Middleware de CORS: Certifique-se de que o middleware de CORS
    está configurado corretamente no arquivo routes/routes.go para evitar
    problemas de segurança em produção.

    
