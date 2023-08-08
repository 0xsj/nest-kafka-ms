package authentication

import (
	pb "gateway/generated"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(address string) pb.AuthenticationServiceClient {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return pb.NewAuthenticationServiceClient(connection)
}