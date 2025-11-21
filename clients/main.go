package main

import (
	"context"
	"fmt"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kmsg"
)

func main() {

	seeds := []string{"localhost:9092"}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ClientID("my-client"),
		kgo.RequestTimeoutOverhead(10*time.Second),
		kgo.RequestRetries(3),
	)

	if err != nil {
		panic(err)
	}
	defer cl.Close()

	fmt.Println("Client created, pinging broker...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = cl.Ping(ctx)
	if err != nil {
		fmt.Printf("Failed to ping broker: %v\n", err)
		return
	}
	fmt.Println("Successfully connected to broker!")

	metadata, err := cl.Request(ctx, kmsg.NewPtrMetadataRequest())
	if err != nil {
		fmt.Printf("Failed to get metadata: %v\n", err)
	} else {
		resp := metadata.(*kmsg.MetadataResponse)
		fmt.Printf("Brokers: %+v\n", resp.Brokers)
		for _, topic := range resp.Topics {
			fmt.Printf("Topic: %s, Partitions: %d\n", *topic.Topic, len(topic.Partitions))
		}
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel2()

	record := &kgo.Record{Topic: "quickstart-events", Value: []byte("bar")}
	fmt.Println("Producing record...")

	results := cl.ProduceSync(ctx2, record)
	if err := results.FirstErr(); err != nil {
		fmt.Printf("Error producing record: %v\n", err)
		return
	}

	for _, r := range results {
		fmt.Printf("Successfully produced record to partition %d at offset %d\n", r.Record.Partition, r.Record.Offset)
	}

	fmt.Println("Done!")

}
