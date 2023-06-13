/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/desertbit/wego"
)

func main() {
	// Create client.
	// The client will automatically login and renew its token regularly.
	c, err := wego.NewClient(wego.Options{
		RemoteAddr: "https://your.wekanboard.com",
		Username:   "user",
		Password:   "secure-password",
	})
	if err != nil {
		log.Fatal(err)
	}

	boards, err := c.GetPublicBoards(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Public boards: %+v\n", boards)

	self, err := c.GetCurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Self: %+v\n", self)

	other, err := c.GetUser(context.Background(), "user-id-of-somebody")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Other: %+v\n", other)
}
