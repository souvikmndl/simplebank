package gapi

import (
	"fmt"

	db "github.com/souvikmndl/simplebank/db/sqlc"
	"github.com/souvikmndl/simplebank/pb"
	"github.com/souvikmndl/simplebank/token"
	"github.com/souvikmndl/simplebank/util"
)

// Server serves grpc requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new grpc server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
