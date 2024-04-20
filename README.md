# GOLANG SCRATCH

<!--toc:start-->

- [GOLANG SCRATCH](#golang-scratch)
  - [tech and tools](#tech-and-tools)
  - [database](#database)
  <!--toc:end-->

## tech and tools

- chi router
- cors
- handlerWithJson
- commitlint & linter
- database
  - goose
  - sqlc
- docker
- auth
  - apikey
- relationship

## database

```mermaid.js
erDiagram
    users {
        id uuid
        name text
        apikey varchar(64)
        created_at timestamp
        updated_at timestamp
    }

    feeds {
        id uuid
        name text
        url text
        created_at timestamp
        updated_at timestamp
        user_id uuid
    }

    users ||--o{ feeds : contains

```
