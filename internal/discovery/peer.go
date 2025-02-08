package discovery

import (
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
)

func init() {
	kadDHT, err := dht.New()
	if err != nil {
		log.Fatal(err)
	}
}

func create_node() (host.Host, error) {
	node, err := libp2p.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DHT node w/ID", node.ID())

	return node, err
}
