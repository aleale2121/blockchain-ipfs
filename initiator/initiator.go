package initiator

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aleale2121/fileverse/internal/constant/model"
	"github.com/aleale2121/fileverse/internal/glue/routing"
	"github.com/aleale2121/fileverse/internal/handlers/rest"
	localStorage "github.com/aleale2121/fileverse/internal/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/aleale2121/fileverse/platform/ipfs"
	"github.com/aleale2121/fileverse/platform/routers"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
)

func Init() {
	log.Println("--------Initiating----------")

	logger := logrus.New()
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}

	sh := shell.NewShell(config.IPFSServerAddress)
	if !sh.IsUp() {
		log.Fatalf("unable to find IPFS node running %s", config.IPFSServerAddress)
	}

	id, err := sh.ID()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get base path: %v", err)
	}
	p, err := filepath.Abs(basePath)
	if err != nil {
		log.Fatal(err)
	}

	store, err := localStorage.NewStorage(basePath)
	if err != nil {
		log.Fatalf("cannot create storage: %v", err)
	}

	fileDownloadPath := filepath.Join(p, config.FileDownloadDir)
	log.Println(fileDownloadPath)
	ipfsClient := ipfs.NewIPFSClient(fileDownloadPath, id.PublicKey, sh)
	fileHandler := rest.NewFileHandler(logger, store, ipfsClient)
	fileRouting := routing.FileRouting(fileHandler, gin.HandlersChain{})

	var routersList []routers.Router
	routersList = append(routersList, fileRouting...)

	server := routers.NewRouting(config.HTTPHost, config.HTTPPort, routersList)
	server.Serve()
}

func LoadConfig() (config model.Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
