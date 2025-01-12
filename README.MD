# **Go Kafka Project**

**Este projeto é uma aplicação de exemplo utilizando Golang e Apache Kafka para produção e consumo de mensagens. Ele demonstra como configurar um ambiente Kafka com Docker e integrar produtores e consumidores em Go.**

## **Ferramentas Utilizadas**

### **Tecnologias**

- **Golang**: Linguagem de programação utilizada para implementar o produtor e o consumidor.
- **Apache Kafka**: Sistema de mensageria distribuído utilizado para troca de mensagens.
- **Docker**: Ferramenta para containerização das aplicações e serviços.

### **Dependências**

- [**confluent-kafka-go**](https://github.com/confluentinc/confluent-kafka-go): Biblioteca para interação com o Kafka em Go.

## **Configuração do Ambiente**

**Certifique-se de ter o Docker e Docker Compose instalados em sua máquina.**

### **Estrutura do Projeto**

```
.
├── docker-compose.yml
├── producer.go
├── consumer.go
├── README.md
└── config
    ├── zookeeper-data
    └── kafka-data
```

### **Configuração de Volumes**

As pastas `zookeeper-data` e `kafka-data` são utilizadas para persistir os dados do Zookeeper e Kafka. Se elas não forem criadas automaticamente, você pode criá-las manualmente:

```bash
mkdir -p config/zookeeper-data config/kafka-data
```

### **Subindo os Contêineres**

Para subir os contêineres com Kafka e Zookeeper, execute:

```bash
docker-compose up -d
```

Isso iniciará os seguintes serviços:

- **Zookeeper**
- **Kafka**

## **Como Utilizar**

### **Criando um Tópico**

Para criar um tópico no Kafka chamado `teste` com 3 partições e fator de replicação igual a 1, execute o seguinte comando:

```bash
docker compose exec -it kafka kafka-topics --create \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 \
  --partitions 3 \
  --topic teste
```

### **Executando o Produtor**

O arquivo `producer.go` é responsável por enviar mensagens para o tópico Kafka. Para executá-lo:

```bash
go run cmd/producer/main.go
```

### **Executando o Consumidor**

O arquivo `consumer.go` é responsável por consumir mensagens do tópico Kafka. Para executá-lo:

```bash
go run cmd/consumer/main.go
```

## **Exemplos de Logs**

### **Logs do Produtor**

```log
2025/01/12 12:49:51 Configuração Kafka: &map[acks:all bootstrap.servers:localhost:9092 delivery.timeout.ms:0 enable.idempotence:true]
2025/01/12 12:49:51 Produtor Kafka criado com sucesso
2025/01/12 12:49:51 Mensagem enviada com sucesso para o Kafka
```

### **Logs do Consumidor**

```log
2025/01/12 12:50:00 Mensagem recebida: transferencia - Partição: 0 - Offset: 10
```

## **Solução de Problemas**

### **Erro: "Coordinator load in progress"**

Se você encontrar o erro `Coordinator load in progress`, ele indica que o broker ainda está inicializando. Aguarde alguns segundos e tente novamente.

## **Contribuição**

Contribuições são bem-vindas! Por favor, crie uma issue ou um pull request com melhorias ou correções.
