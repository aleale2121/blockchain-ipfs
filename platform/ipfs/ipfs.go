package ipfs

import (
	"time"
	"io"
	shell "github.com/ipfs/go-ipfs-api"
)

type IPFSClient struct {
	LocalPath string
	PublicKey string
	Sh        *shell.Shell
}

func NewIPFSClient(LocalPath string, PublicKey string, Sh *shell.Shell) *IPFSClient {
	ipfs := IPFSClient{
		LocalPath: LocalPath,
		PublicKey: PublicKey,
		Sh:        Sh,
	}
	return &ipfs
}

func (ipfs IPFSClient) AddFile(r io.Reader) (string, error) {
	return ipfs.Sh.Add(r)
}

func (ipfs IPFSClient) DownloadFile(cid string) error {
	return ipfs.Sh.Get(cid, ipfs.LocalPath)
}

func (ipfs IPFSClient) AddToIPNS(cid string) error {
	var lifetime time.Duration = 50 * time.Hour
	var ttl time.Duration = 1 * time.Microsecond

	_, err := ipfs.Sh.PublishWithDetails(cid, ipfs.LocalPath, lifetime, ttl, true)
	return err
}

func (ipfs IPFSClient) ResolveIPNS() (string, error) {
	return ipfs.Sh.Resolve(ipfs.PublicKey)
}
