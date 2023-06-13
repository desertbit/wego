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
		RemoteAddr: "https://board.wahtari.io",
		Username:   "webhooker",
		Password:   "N1zHKfSMpEBqqNBzjXAAcXiW1Ey6rrzgRrvnRKhnn1U1dguzk4aGIiYfBvjDyOm1",
	})
	if err != nil {
		log.Fatal(err)
	}

	const boardID = "GXCary4RDoJqR8n3u"
	listIDs := []string{"8EXETWQ9YBThk7vjt", "kkhTPJDcLsSZ9hRYs"}
	_ = listIDs
	lists, err := c.GetAllLists(context.Background(), boardID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lists: %+v\n", lists)

	u, err := c.GetComment(context.TODO(), boardID, "fN556xuzJFSEDCrn48k", "tinaer")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user: %+v\n", u)

	/*for _, listID := range listIDs {
		cards, err := c.GetAllCards(context.Background(), boardID, listID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("cards: %+v\n", cards)
	}*/

	/*card, err := c.GetCard(context.Background(), boardID, listIDs[1], "fN556xuzJFSEDC48k")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("card: %+v\n", card)*/

	r, err := c.EditCard(context.Background(), boardID, listIDs[1], "fN556xuzJFSEDC48k", wego.EditCardOptions{
		ListID: listIDs[0],
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response: %+v\n", r)
}
