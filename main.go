package main

import (
	"context"
	"github.com/MbungeApp/mbunge-core/config"
)

func main() {
	client := config.ConnectDB()
	defer client.Disconnect(context.Background())
}
