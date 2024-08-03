package howtocomposetypedfunctions

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Server struct{}

func (s *Server) handleRequest(filename string) error {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])

	fmt.Println("new filename:", newFilename)
	return nil
}

func Example1() {
	server := &Server{}

	server.handleRequest("someName.png")
}
