#!/bin/sh

echo "Creating Kafka topics..."
kafka-topics.sh --create --bootstrap-server kafka:9092 --replication-factor 1 --partitions 1 --topic progressupdate
