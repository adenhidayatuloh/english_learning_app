# #!/bin/sh

# echo "Creating Kafka topics..."

docker exec -it 16f396d13643 /bin/bash
kafka-topics --create --bootstrap-server kafka:9092 --replication-factor 1 --partitions 1 --topic progressupdate


# kafka-topics --create \
#   --topic my-topic \
#   --bootstrap-server localhost:9092 \
#   --partitions 1 \
#   --replication-factor 1


#   kafka-console-consumer --bootstrap-server localhost:9092 --topic progressupdate --group lesson-update-group --from-beginning