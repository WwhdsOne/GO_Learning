package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Kafka 集群地址
	brokers := []string{"localhost:9092"}

	// 创建 Kafka 配置
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0 // 设置 Kafka 版本
	config.Producer.Return.Successes = true

	// 创建 Kafka 管理客户端
	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create cluster admin: %s", err)
	}
	defer admin.Close()

	// 创建 Kafka 主题
	topic := "test-topic"
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	err = admin.CreateTopic(topic, topicDetail, false)
	if err != nil {
		log.Printf("Failed to create topic: %s", err)
	} else {
		log.Printf("Topic %s created successfully", topic)
	}

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer producer.Close()

	// 发送消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %s", err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	// 等待中断信号以优雅地关闭生产者
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
