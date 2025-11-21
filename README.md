# AutoMQ PoC

Proof of Concept (PoC) para testar o AutoMQ com MinIO como storage backend e cliente Go para produção de mensagens.

## Estrutura do Projeto

```
automq-poc/
├── docker/
│   └── docker-compose.yaml    # Setup do AutoMQ + MinIO
└── clients/
    ├── main.go                # Cliente Go para produção de mensagens
    └── go.mod                 # Dependências Go
```

## Pré-requisitos

- Docker e Docker Compose
- Go 1.21+ (para executar o cliente)
- Acesso sudo (para configurar /etc/hosts)

## Setup Inicial

### 1. Configurar o /etc/hosts

Para que o cliente consiga se conectar ao broker, adicione a seguinte entrada no `/etc/hosts`:

```bash
echo "127.0.0.1 server1" | sudo tee -a /etc/hosts
```

Isso é necessário porque o broker AutoMQ anuncia `server1` nos metadados de conexão.

### 2. Iniciar o Cluster AutoMQ

```bash
cd docker
docker compose up -d
```

Isso irá iniciar:
- **MinIO**: Storage S3-compatible na porta 9000 (API) e 9001 (Console)
- **AutoMQ Broker**: Kafka-compatible broker na porta 9092
- **mc**: Cliente MinIO para criar buckets automaticamente

Aguarde ~30 segundos para o cluster inicializar completamente.

### 3. Verificar Status

```bash
# Verificar containers em execução
docker ps

# Verificar logs do broker
docker logs automq-single-server

# Listar tópicos
docker exec automq-single-server /opt/automq/kafka/bin/kafka-topics.sh \
  --bootstrap-server localhost:9092 --list
```

## Cliente Go

O cliente Go demonstra como:
- Conectar ao AutoMQ via franz-go
- Fazer ping no broker
- Consultar metadados (brokers e tópicos)
- Produzir mensagens de forma síncrona

### Executar o Cliente

```bash
cd clients
go run main.go
```

**Output esperado:**
```
Client created, pinging broker...
Successfully connected to broker!
Brokers: [{NodeID:0 Host:server1 Port:9092 Rack:<nil> UnknownTags:{keyvals:map[]}}]
Topic: quickstart-events, Partitions: 1
Topic: __auto_balancer_metrics, Partitions: 1
Producing record...
Successfully produced record to partition 0 at offset X
Done!
```

## Configurações do AutoMQ

### Listeners
- **PLAINTEXT**: `0.0.0.0:9092` (dentro do container)
- **Advertised**: `localhost:9092` (para clientes externos)
- **CONTROLLER**: `server1:9093` (para quorum interno)

### Storage (MinIO)
- **Data Bucket**: `automq-data` - Armazenamento de mensagens
- **Ops Bucket**: `automq-ops` - Metadados operacionais
- **Endpoint**: `http://minio:9000` (dentro da rede Docker)

### Recursos
- **Heap**: 1GB inicial, 4GB máximo
- **MetaSpace**: 96MB
- **Direct Memory**: 1GB

