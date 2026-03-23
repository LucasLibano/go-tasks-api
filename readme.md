# GoTasks

Esse é um projeto de API simples em **Go + SQLite** que cria, lista, atualiza e deleta tarefas.  
Usei Docker pra rodar a API e salvar os dados no banco mesmo quando fecha o container.

---

## Como usar

1. **Build da imagem Docker**
```bash
docker build -t gotasks:latest .
```

2. **Criar volume pro banco**
```bash
docker volume create gotasks_data
```

3. **Rodar a API**
```bash
docker run --name gotasks -p 8080:8080 -v gotasks_data:/data gotasks:latest
```

A API vai rodar na porta `8080`.

---

## Endpoints

- **GET /healthz** → Ver se tá ok  
- **POST /tasks** → Criar tarefa (`title` obrigatório, `done` opcional)  
- **GET /tasks** → Lista todas as tarefas  
- **GET /tasks/{id}** → Ver uma tarefa pelo id  
- **PUT /tasks/{id}** → Atualizar título ou done  
- **DELETE /tasks/{id}** → Deletar tarefa

Exemplo pra criar tarefa:
```bash
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"estudar","done":false}'
```

---

## Dependências

- Go 1.22+  
- Docker  
- SQLite  
- `github.com/mattn/go-sqlite3`  
- `github.com/gorilla/mux`  

---


