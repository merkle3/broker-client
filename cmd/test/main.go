package main

import (
	"context"
	"time"

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

	chainId := int32(1)

	stream, err := broker.StreamReceivedTransactions(context.Background(), &proto.TxStreamRequest{
		ApiKey:  "0x0000000000000000000000000000000000000000", // not required while in beta, please use your ethereum address in the meantime
		ChainId: chainId,                                      // optional, defaults to 1 (mainnet)
	})

	log.WithFields(log.Fields{
		"chainId": chainId,
	}).Info("connected to tx stream")

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

		log.WithFields(log.Fields{
			"timestamp": time.Now(),
			"lag_ms":    time.Now().UnixMilli() - txPacket.ArrivalTime,
			"chainId":   tx.ChainId().Uint64(),
		}).Info(tx.Hash().String())
	}
}
