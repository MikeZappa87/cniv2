package libcni

import (
	cni "github.com/MikeZappa87/cniv2/pkg/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//unix:///tmp/cni.sock
func NewNetworkRuntimeService(endpoint string) (cni.CNIClient, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(endpoint, opts)
	if err != nil {
		return nil, err
	}
	
	return cni.NewCNIClient(conn), nil
}

