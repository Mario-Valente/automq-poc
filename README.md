# AutoMQ PoC

PoC para testar AutoMQ com MinIO (local) e AWS BYOC (cloud).

## Pré-requisitos

- Docker e Docker Compose
- Terraform (para deploy AWS)
- Go 1.21+ (cliente opcional)

## Setup Local (Docker)

### 1. Adicionar hostname ao /etc/hosts

```bash
echo "127.0.0.1 server1" | sudo tee -a /etc/hosts
```

### 2. Iniciar cluster

```bash
cd docker
make up
```

### 3. Comandos úteis

```bash
make help              # Ver todos os comandos
make status            # Ver status dos containers
make logs-follow       # Ver logs em tempo real
make test-connection   # Testar conexão com broker
```

### 4. Operações Kafka

```bash
# Criar tópico
make create-topic TOPIC=my-topic PARTITIONS=3

# Listar tópicos
make list-topics

# Produzir mensagens
make produce TOPIC=my-topic

# Consumir mensagens
make consume TOPIC=my-topic

# Teste completo
make test-produce-consume
```

### 5. Acessar MinIO Console

http://localhost:9001 (minioadmin/minioadmin)

```bash
make minio-console  # Mostrar credenciais
make minio-stats    # Ver estatísticas dos buckets
```

### 6. Parar cluster

```bash
make down       # Parar containers
make clean      # Parar e remover volumes
```

## Deploy AWS (Terraform)

### 1. Inicializar

```bash
cd aws/terraform
make init
```

### 2. Deploy

```bash
# Ver plano
make plan

# Aplicar
make apply

# Com região e perfil específicos
make apply REGION=us-west-2 PROFILE=my-profile
```

### 3. Ver credenciais

```bash
make show-credentials
```

### 4. Outros comandos

```bash
make help           # Ver todos os comandos
make validate       # Validar configuração
make fmt            # Formatar arquivos
make output         # Ver outputs
make destroy        # Destruir infraestrutura
```

## Cliente Go

```bash
cd clients
go run main.go
```

O cliente conecta ao broker local, lista metadados e produz uma mensagem de teste.

## Endpoints

- **Kafka Broker**: `localhost:9092`
- **MinIO API**: `localhost:9000`
- **MinIO Console**: `localhost:9001`
