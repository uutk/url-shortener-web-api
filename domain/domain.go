package domain

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	pb "github.com/vladkampov/url-shortener/service"
	"google.golang.org/grpc"
	"os"
	"time"
)

var c pb.ShortenerClient

func GetUrl(hash string) (string, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUrl(ctx, &pb.HashedUrlRequest{Hash: hash})
	if err != nil {
		log.Warnf("could not greet: %v", err)
	}

	if len(r.Url) == 0 {
		err = fmt.Errorf("No URL for hash %s ", hash)
	}

	log.Printf("URL was successfully got by hash %s: %s", hash, r.Url)
	return r.Url, err
}

func SendUrl(url string) (string, error) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Shorten(ctx, &pb.URLRequest{Url: url})
	if err != nil {
		return "", err
	}
	log.Printf("URL was successfully shortened: %s", r.Url)
	return r.Url, nil
}

func InitDomainGRPCSession() pb.ShortenerClient {
	domainServiceUrl := os.Getenv("SHORTENER_DOMAIN_PORT")

	if len(domainServiceUrl) == 0 {
		domainServiceUrl = "localhost:50051"
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(domainServiceUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c = pb.NewShortenerClient(conn)
	return c
}
