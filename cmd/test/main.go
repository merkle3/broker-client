package main

import (
	"context"

	"github.com/merkle3/broker-client/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// this example code connect to merkle tx stream
	// and print out the tx hash

	// connect to tx broker
	conn, err := grpc.Dial("txs.usemerkle.com:80", grpc.WithInsecure())

	if err != nil {
		log.WithError(err).Fatal("error connecting to tx stream")
	}

	broker := proto.NewBrokerApiClient(conn)

	log.Info("connected to broker")

	stream, err := broker.StreamReceivedTransactions(context.Background(), &proto.TxStreamRequest{
		Name: "<API KEY>", // not required while in beta
	})

	log.Info("connected to tx stream")

	if err != nil {
		log.WithError(err).Fatal("error connecting to tx stream")
	}

	for {
		txPacket, err := stream.Recv()

		if err != nil {
			log.WithError(err).Fatal("error receiving tx")
		}

		tx, err := proto.ToTransaction(txPacket)

		if err != nil {
			log.WithError(err).Error("could not decode tx")
			continue
		}

		log.Info(tx.Hash().String())
	}
}
