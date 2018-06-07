package main

import (
	"LearningGo/GoGrpcDemo/protobuff"
	"sync"
	"golang.org/x/net/context"
	"fmt"
	"flag"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/testdata"
	"google.golang.org/grpc/credentials"
	pb "LearningGo/GoGrpcDemo/protobuff"
)

type AccountServicesServer struct {

	savedAccounts []protobuff.Account

	mu sync.Mutex

}


func (server *AccountServicesServer) AddMoneyToAccount(ctx context.Context, account *protobuff.Account) (*protobuff.Account, error) {
	server.getAccountById(account.AccountId)
	server.mu.Lock()
	if accountFound := server.getAccountById(account.AccountId); accountFound != nil {
		accountFound.Amount += account.Amount
		server.mu.Unlock()
		return accountFound, nil
	}
	server.mu.Unlock()
	return account, fmt.Errorf("Account does not exists")
}

func (server *AccountServicesServer) getAccountById(accountId string) *protobuff.Account {
	for _, savedAccount := range server.savedAccounts {
		if savedAccount.AccountId == accountId {
			return &savedAccount
		}
	}
	return nil
}

func (server *AccountServicesServer) GetAccount(ctx context.Context, details *pb.RequestAccountDetails) (*pb.Account, error) {
	if accountFound := server.getAccountById(details.AccountId); accountFound != nil {
		return accountFound, nil
	}
	server.mu.Unlock()
	return &pb.Account{}, fmt.Errorf("Account does not exists with id: " + details.AccountId)
}


func (server *AccountServicesServer) CreateAccount(ctx context.Context, accountId *pb.AccountId) (*protobuff.Account,error) {
	newAccount := protobuff.Account{AccountId:accountId.GetAccountId()}
	server.mu.Lock()
	if accountFound := server.getAccountById(accountId.GetAccountId()); accountFound == nil {
		server.savedAccounts = append(server.savedAccounts, newAccount)
		server.mu.Unlock()
		log.Println("Account created with id: ", accountId)
		log.Println("NumberOfAccounts so far : ", len(server.savedAccounts))
		return &newAccount, nil
	} else {
		newAccount = *accountFound
		server.mu.Unlock()
		log.Println("Account creation failed for id: ", accountId)
		log.Println("NumberOfAccounts so far : ", len(server.savedAccounts))
		return accountFound, fmt.Errorf("Account ceartion failed. Account already exists for id: " + accountId.GetAccountId())
	}
}

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "testdata/route_guide_db.json", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)


func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	aa := AccountServicesServer{}
	pb.RegisterAccountServicesServer(grpcServer, &aa)
	grpcServer.Serve(lis)
}