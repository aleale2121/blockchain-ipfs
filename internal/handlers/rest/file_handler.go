package rest

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aleale2121/fileverse/internal/pkg/storage"
	"github.com/aleale2121/fileverse/platform/ipfs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FileHandler struct {
	logger     *logrus.Logger
	store      *storage.Storage
	ipfsClient *ipfs.IPFSClient
}

func NewFileHandler(logger *logrus.Logger,
	store *storage.Storage,
	ipfsClient *ipfs.IPFSClient,
) FileHandler {
	return FileHandler{
		logger:     logger,
		store:      store,
		ipfsClient: ipfsClient,
	}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	currentTime := time.Now().Local()

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to open the file",
		})
		return
	}
	defer src.Close()

	// Add the file to IPFS
	cid, err := h.ipfsClient.AddFile(src)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add the file to IPFS",
			"error":   err.Error(),
		})
		return
	}

	response := gin.H{
		"fileId":    cid,
		"size":      file.Size,
		"timestamp": currentTime.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}

func (h *FileHandler) DownloadFileByCID(c *gin.Context) {
	// Get the CID from the URL parameters
	cid := c.Param("fileId")

	// Download the file from IPFS using the CID
	err := h.ipfsClient.DownloadFile(cid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to download the file from IPFS",
			"error":   err.Error(),
		})
		return
	}
	listDirectoryContents(h.ipfsClient.LocalPath)
	// Serve the downloaded file to the client
	c.File(filepath.Join(h.ipfsClient.LocalPath, cid))

	// Optionally, you can remove the file from your local storage after serving it
	// err = os.Remove(filepath.Join(h.ipfsClient.LocalPath, cid))
	// if err != nil {
	// 	h.logger.Errorf("Failed to remove downloaded file: %v", err)
	// }
}

func (h *FileHandler) Ping(c *gin.Context) {
	c.JSON(200, "ok")
}

func listDirectoryContents(directoryPath string) error {
	// Open the directory.
	dir, err := os.Open(directoryPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	// Read the directory contents.
	entries, err := dir.ReadDir(0) // 0 means no limit on the number of entries
	if err != nil {
		return err
	}
	fmt.Println(len(entries))
	// Iterate through the directory entries and print their names.
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	return nil
}
