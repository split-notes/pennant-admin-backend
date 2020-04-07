package hello_requests

import (
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/split-notes/pennant-admin-backend/library/appcontext"
	"net/http"
)

type GetHelloRequest struct {
}

//Decode the Push notification Req.
func MakeGetHelloRequestDecoder(appCtx appcontext.Context) (kithttp.DecodeRequestFunc, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		//decode := json.NewDecoder(r.Body)
		var Req GetHelloRequest
		//err := decode.Decode(&Req)
		//if err != nil {
		//	return nil, errors.New("inconsistent mapping between route and handler")
		//}

		// GET requests don't have bodies
		return Req, nil
	}, nil
}
