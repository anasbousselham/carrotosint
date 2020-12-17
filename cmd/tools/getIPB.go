package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cavaliercoder/grab"
)

var urlBlock = [1]string{"https://www.ipdeny.com/ipblocks/data/aggregated/"}

func main() {
	argsCountry := os.Args[1]
	client := grab.NewClient()
	req, _ := grab.NewRequest("_"+argsCountry+".db", urlBlock[0]+argsCountry+"-aggregated.zone")

	// start download
	fmt.Printf("Downloading Block IP %s...\n", strings.ToUpper(argsCountry))
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress())

		case <-resp.Done:
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)

}
