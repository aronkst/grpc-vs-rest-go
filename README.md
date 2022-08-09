# gRPC vs REST

This repository is a test containing a simple implementation of gRPC and REST developed in Go that runs inside a Docker, in which it is possible to define the amount of requests sent via gRPC and REST and verify the performance of both.

Its operation is very simple and can be done entirely using the make commands of the Makefile file.

# How to Use

Run the application:

`make run`

Start the application:

`make start`

Stop the application:

`make stop`

View the application logs:

`make logs`

Test via REST with a count of requests:

(Replace NUMBER with the number of requests to be sent during the test.)

`make test-rest count=NUMBER`

Test via gRPC with a count of requests:

(Replace NUMBER with the number of requests to be sent during the test.)

`make test-grpc count=NUMBER`

Generate Proto File (gRPC):

`make generate-proto`
