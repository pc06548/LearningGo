package main

import (
	"log"
	"time"
	"flag"
	"google.golang.org/grpc/testdata"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	pb "LearningGo/GoGrpcDemo/protobuff"

	"golang.org/x/net/context"
	"strconv"
	"math/rand"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "server.crt", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "0.0.0.0:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "a", "The server name use to verify the hostname returned by TLS handshake")
)

func createAccount(client pb.AccountServicesClient, accountId pb.AccountId, ctx context.Context) {
	createdAccount, err := client.CreateAccount(ctx, &accountId)
	if err != nil {
		log.Println("Account Creation error: %v.GetFeatures(_) = _, %v: ", client, err)
	} else {
		log.Println("Account created: ", createdAccount)
	}
}


func main() {
	execute()
}

func execute() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAccountServicesClient(conn)
	s1 := rand.NewSource(10)
	r1 := rand.New(s1)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for i := 0; i < 10; i++ {
		createAccount(client, pb.AccountId{AccountId: strconv.Itoa(r1.Intn(10))}, ctx)
	}

	/*for i := 0; i < 10; i++ {
		account, err := client.GetAccount(ctx, &pb.RequestAccountDetails{AccountId: strconv.Itoa(i)})
		if err == nil {
			log.Println("Reading account info for: ", account.AccountId, ". Balance is: ", account.Amount)
		} else {
			log.Println("Account reading error: %v.GetFeatures(_) = _, %v: ", client, err)
		}
	}*/
}
