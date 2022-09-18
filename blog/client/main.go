package main

import (
	"blog/x/blog/types"
	"context"
	"fmt"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"log"
)

func main() {
	addressPrefix := "blog"

	// Create a Cosmos client instance
	cosmos, err := cosmosclient.New(
		context.Background(),
		cosmosclient.WithAddressPrefix(addressPrefix),
	)
	if err != nil {
		log.Fatalln(err)
	}

	accountName := "alice"

	// Get account from the keyring
	account, err := cosmos.Account(accountName)
	if err != nil {
		log.Fatalln(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	// Define a message to create a post
	msg := &types.MsgCreatePost{
		Creator: addr,
		Title:   "Hello!",
		Body:    "This is the first post!",
	}

	txResp, err := cosmos.BroadcastTx(account, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)

	queryClient := types.NewQueryClient(cosmos.Context())

	queryResp, err := queryClient.Posts(context.Background(), &types.QueryPostsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)
}
