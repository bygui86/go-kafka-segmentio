
# SegmentIO - kafka-go - Example notes

## Compatibility versions

### 1st test > OK

Zookeeper - wurstmeister/zookeeper:3.4.6
Kafka - wurstmeister/kafka
producer-random - go1.15 + kafka-go v0.4.15
consumer-logger - go1.16 + kafka-go v0.4.15

### 2st test > OK

Zookeeper - wurstmeister/zookeeper:3.4.6
Kafka - wurstmeister/kafka:2.12-2.3.1
producer-random - go1.15 + kafka-go v0.4.15
consumer-logger - go1.16 + kafka-go v0.4.15

### 3rd test > OK

Zookeeper - wurstmeister/zookeeper
Kafka - wurstmeister/kafka
producer-random - go1.15 + kafka-go v0.4.15
consumer-logger - go1.16 + kafka-go v0.4.15

### 4th test > OK

Zookeeper - wurstmeister/zookeeper
Kafka - wurstmeister/kafka
producer-random - go1.15 + kafka-go v0.4.18
consumer-logger - go1.16 + kafka-go v0.4.18

### 5th test > OK

Zookeeper - wurstmeister/zookeeper
Kafka - wurstmeister/kafka
producer-random - go1.17 + kafka-go v0.4.18
consumer-logger - go1.17 + kafka-go v0.4.18
