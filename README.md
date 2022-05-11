## Kafka practicing using go lanuage and go library sarama

find the ip address of the  machine using the following command
ifconfig
under the virbr0 you can find the ip address

To run the kafka-zookeeper run the following code
sudo docker run -p 2181:2181 -p 9092:9092 -e ADVERTISED_HOST=192.168.122.1  -e NUM_PARTITIONS=4 johnnypark/kafka-zookeeper

To run the UI tool for the kafka
sudo docker run -p 9000:8080 -e KAFKA_BROKERS=192.168.122.1:9092 quay.io/cloudhut/kowl:master

then you can see the kafka broker and topic details using 
localhost:9000 address


above both docker commands require the same IP which can be found using the ifconfig.

1. create the folder workspace (say kafka)

2. Change the GOPATH envioronment variable to the workspace folder as follows
   export GOPATH=../../../kafka

3. Create a folder called src inside that

4. Inside the src folder create another folder called whatever name (say "main")

5. Inside src/main folder create a go file for coding (sat "app.go")

6. Initialize the go.mod file using the follwing command
   go mod init app

7. Write whatever code you need

8. Install the third party libraries from github
   example - go get github.com/Shopify/sarama

9. Inorder to install the go tools you need to use the go install command
   much better if we can give the URL with preceding @latest wording
   go get does not work for me with sarama tools
   go install github.com/Shopify/sarama/tools/kafka-console-producer@latest

after installing the kafka-console-producer tool no need to change the GOROOT or PATH variable

Without changing the above variables we can run the kafka-console-producer using as follows 
bin/kafka-console-producer -topic=test -value=value -brokers=localhost:9092

Now  you can neglect the 10 and 11 steps

10. change the GOROOT envioronment variable to the workspace folder
No need to change the GOROOT variable only need to change the PATH variable
    export GOROOT=/home/asitha/Desktop/kafka

11. Change the PATH variable to the bin of the workspace
    export PATH=/home/asitha/Desktop/kafka/bin
    But now you cant run any go command within that terminal

12. Now you can run the command just installed into the bin folder
    kafka-console-producer -topic=test -value=value -brokers=localhost:9092

13. Check the kowl whether the message reached to the broker at
    localhost:9000

   
we can specify a key and value both using the terminal command in producer console

we can change the partitioner flag as well 
bin/kafka-console-producer -topic=test -key=key1 -value=value -brokers=localhost:9092 -partitioner=random

for the partitioning theres a key based method as well which is hash

bin/kafka-console-producer -topic=test -key=key1 -value=value -brokers=localhost:9092 -partitioner=hash

so same key values always stack up on one partition. 
bin/kafka-console-producer -topic=test -key=222 -value=value -brokers=localhost:9092 -partitioner=hash

14. To install kafka-console-consumer 
go install github.com/Shopify/sarama/tools/kafka-console-consumer@latest
Make sure GOPATH is set to the workspace 

15. You can listen to the kafka brokers at specific address using the console-consumer
bin/kafka-console-consumer -topic=test -brokers=localhost:9092

If you produce new message then it will show it under the terminal of consumer-console

-offset flag has a default value of 'newest' and we can set it to 'oldest' then we can access the older values produced to the broker 

