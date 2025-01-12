package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	if producer == nil {
		log.Fatal("Falha ao criar produtor Kafka")
	}

	err := Publish("Mensagem", "teste", producer, []byte("transferencia"), deliveryChan)
	if err != nil {
		log.Fatalf("Erro ao publicar mensagem: %v", err)
	}

	go DeliveryReport(deliveryChan)
	producer.Flush(5000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	log.Printf("Configuração Kafka: %+v", configMap)
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Fatalf("Erro ao criar producer Kafka: %v", err)
	}

	log.Println("Produtor Kafka criado com sucesso")
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	if producer == nil {
		return fmt.Errorf("produtor Kafka não inicializado")
	}

	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(message, deliveryChan)
	if err != nil {
		log.Printf("Erro ao produzir mensagem: %v", err)
		return err
	}

	log.Println("Mensagem enviada com sucesso para o Kafka")
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Printf("Erro ao enviar mensagem para o Kafka: %v", ev.TopicPartition.Error)
			} else {
				log.Printf("Mensagem enviada com sucesso para o Kafka: %v", ev.TopicPartition)
			}
		}
	}
}
