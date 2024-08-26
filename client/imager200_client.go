package client

import (
	"context"
	"net/http"

	"github.com/imager200/go-sdk/generated"
)

const (
	server                      = "https://api.imager200.io"
	apiKeyHeader                = "X-Imager-Key"
	postOperationIDHeader       = "X-PostOp-Id"
	uploadFileNameHeader        = "X-Upload-File-Name"
	imgurImageTitleHeader       = "X-Imgur-Title"
	imgurImageDescriptionHeader = "X-Imgur-Description"
	gphotoDescriptionHeader     = "X-GPhoto-Description"
)

func NewImager200Client(apiKey string) (*generated.ClientWithResponses, error) {
	return generated.NewClientWithResponses(server, apiKeyClientOption(apiKey))
}

var apiKeyClientOption = func(apiKey string) generated.ClientOption {
	return generated.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set(apiKeyHeader, apiKey)
		return nil
	})
}

// request editor that sets the post operation id header: X-PostOp-Id.
// post operation can be created in the control panel: https://panel.imager200.io/
// it can be used with async endpoints.
var WithPostOperationIDHeader = func(postOperationID string) generated.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set(postOperationIDHeader, postOperationID)
		return nil
	}
}

// request editor that sets the upload file name header: X-Upload-File-Name.
// The value should contain the file name only without extension
// it can be used with async endpoints.
var WithUploadFileNameHeader = func(filename string) generated.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set(uploadFileNameHeader, filename)
		return nil
	}
}

// request editor that sets the X-Imgur-Title header.
// The header set the image title in imgur
// it can be used when an imgur post operation is used.
// otherwise, it has no effect.
var WithImgurImageTitleHeader = func(imageTitle string) generated.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set(imgurImageTitleHeader, imageTitle)
		return nil
	}
}

// request editor that sets the X-Imgur-Description header.
// The header set the image title in imgur
// it can be used when an imgur post operation is used.
// otherwise, it has no effect.
var WithImgurImageDescriptionHeader = func(imageDescription string) generated.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set(imgurImageDescriptionHeader, imageDescription)
		return nil
	}
}

// request editor that sets the upload file name header: X-Upload-File-Name.
// The value should contain the file name only without extension
// it can be used when a Google Photo post operation is used.
// otherwise, it has no effect.
var WithGPhotoImageDescriptionHeader = func(gphotoDescription string) generated.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set(gphotoDescriptionHeader, gphotoDescription)
		return nil
	}
}
