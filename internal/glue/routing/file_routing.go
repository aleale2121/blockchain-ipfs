package routing

import (
	"net/http"

	"github.com/aleale2121/fileverse/internal/handlers/rest"
	"github.com/aleale2121/fileverse/platform/routers"
	"github.com/gin-gonic/gin"
)

func FileRouting(handler rest.FileHandler, middlewares gin.HandlersChain) []routers.Router {
	return []routers.Router{
		{
			Method:      http.MethodPost,
			Path:        "/upload",
			Handle:      handler.UploadFile,
			MiddleWares: middlewares,
		},
		{
			Method:      http.MethodGet,
			Path:        "/file/:fileId",
			Handle:      handler.DownloadFileByCID,
			MiddleWares: middlewares,
		},
		{
			Method:      http.MethodGet,
			Path:        "/ping",
			Handle:      handler.Ping,
			MiddleWares: middlewares,
		},
	}
}
