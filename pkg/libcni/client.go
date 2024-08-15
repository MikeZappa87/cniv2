package libcni

import (
	"google.golang.org/grpc"
)

func NewNetworkRuntimeService(endpoint string) (cni.CNIClient, error) {
	conn, err := grpc.NewClient(endpoint)
	if err != nil {
		return nil, err
	}

	return cni.NewCNIClient(conn), nil
}
