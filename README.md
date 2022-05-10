Kafka practicing using go lanuage and go library sarama

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

10. change the GOROOT envioronment variable to the workspace folder
    export GOROOT=/home/asitha/Desktop/kafka

11. Change the PATH variable to the bin of the workspace
    export PATH=/home/asitha/Desktop/kafka/bin

12. Now you can run the command just installed into the bin folder
    kafka-console-producer -topic=test -value=value -brokers=localhost:9092

13. Check the kowl whether the message reached to the broker at
    localhost:9000
