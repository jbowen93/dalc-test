package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/celestiaorg/dalc/proto/dalc"
	"github.com/celestiaorg/dalc/proto/optimint"
	"google.golang.org/grpc"
)

var (
	dalcClient dalc.DALCServiceClient
)

func main() {

	destIP := os.Getenv("DEST_IP")
	destPort := os.Getenv("DEST_PORT")

	conn, err := grpc.Dial(destIP+":"+destPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// set the global client
	dalcClient = dalc.NewDALCServiceClient(conn)
	height, err := strconv.Atoi(os.Getenv("BLOCK_HEIGHT"))
	if err != nil {
		panic(err)
	}
	submitBlock(uint64(height))
}

func getBlockAvailability(height uint64) {
	resp, err := dalcClient.RetrieveBlock(context.TODO(), &dalc.RetrieveBlockRequest{
		Height: height,
	})
	if err != nil {
		log.Fatal(err)
	}
	hash := resp.Block.Header.DataHash
	fmt.Printf("%v\n", hash)
}

func submitBlock(height uint64) *optimint.Block {
	id := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	block := &optimint.Block{
		Header: &optimint.Header{
			Height:      height,
			NamespaceId: id,
		},
		Data: &optimint.Data{
			Txs: [][]byte{bytes.Repeat([]byte{1, 2, 3, 4}, 2000), {2}, {3, 4}},
		},
		LastCommit: &optimint.Commit{
			Height: height,
		},
	}
	req := dalc.SubmitBlockRequest{
		Block: block,
	}
	resp, err := dalcClient.SubmitBlock(context.TODO(), &req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", resp)
	return block
}

func retrieveBlock(block *optimint.Block) {

}
