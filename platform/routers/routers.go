package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Routers interface {
	Serve()
}

type Router struct {
	Method      string
	Path        string
	Handle      gin.HandlerFunc
	MiddleWares gin.HandlersChain
}
type Routing struct {
	host    string
	port    string
	routers []Router
}

func NewRouting(host, port string, routers []Router) Routers {
	return &Routing{
		host,
		port,
		routers,
	}
}

func (r *Routing) Serve() {
	ginRouter := gin.Default()
	for _, router := range r.routers {
		if router.MiddleWares != nil {
			for _, middle := range router.MiddleWares {
				ginRouter.Use(middle)
			}
		}
		ginRouter.Use(gin.Logger())
		ginRouter.Use(gin.Recovery())
		ginRouter.Handle(router.Method, router.Path, router.Handle)
	}

	addr := fmt.Sprintf("%s:%s", r.host, r.port)
	fmt.Printf("Starting the server at %s \n", addr)

	// curDir, _ := os.Getwd()
	// ginRouter.Static("/file/*fileId", path.Join(curDir,"assets"))

	ginRouter.Run(addr)

}

type Server struct {
}

// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Access-Control-Allow-Origin", "*")
// s.r.ServeHTTP(w, r)
// }
