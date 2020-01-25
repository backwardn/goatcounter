package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	// TODO: Doesn't seem to request <script> tags?
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://www.goatcounter.localhost:8081`),
		chromedp.WaitVisible(`body > footer`), // wait for footer element is visible (ie, page is loaded)
	)
	if err != nil {
		log.Fatal(err)
	}
}
