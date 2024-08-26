![client](https://www.imager200.io/assets/images/main-logo-compressed.png)
# Golang client for the imager200 API

the client provides bindings for interacting with the [imager200 API](https://www.imager200.io/). It is auto-generated from the official open api specs: https://www.imager200.io/imager200.json

To use the API, a (free) [registration](https://panel.imager200.io/) is needed in order to obtain the API key. 

## Installation

`go get -u github.com/imager200/go-sdk`

## Example usage:


```
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/imager200/go-sdk/client"
	"github.com/imager200/go-sdk/generated"
)

func main() {
	cl, err := client.NewImager200Client("Set you API key here")
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()
	fb, err := os.ReadFile("jpeg_image.jpg")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Example 1: /grayscale/sync using POST
	resp, err := cl.GrayscaleSyncPostWithBodyWithResponse(ctx, client.ApplicationOctetStreamContentType, bytes.NewBuffer(fb))
	if err != nil {
		log.Fatal(err.Error())
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("received non ok status: %d, response: %s", resp.StatusCode(), string(resp.Body))
	}

	if err := os.WriteFile("grayscaled.jpg", resp.Body, 0644); err != nil {
		log.Fatal(err.Error())
	}
}
```

More examples can be found in the examples [directory](/examples/).


## Request editors:

For each special header described in the API docs, a `RequestEditorFn` is available:
- [WithPostOperationIDHeader](https://github.com/imager200/go-sdk/blob/main/client/imager200_client.go#L34)
- [WithUploadFileNameHeader](https://github.com/imager200/go-sdk/blob/main/client/imager200_client.go#L44)
- [WithImgurImageTitleHeader](https://github.com/imager200/go-sdk/blob/main/client/imager200_client.go#L55)
- [WithImgurImageDescriptionHeader](https://github.com/imager200/go-sdk/blob/main/client/imager200_client.go#L66)
- [WithGPhotoImageDescriptionHeader](https://github.com/imager200/go-sdk/blob/main/client/imager200_client.go#L77)