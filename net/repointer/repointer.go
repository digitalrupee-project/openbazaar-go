package net

import (
	"time"

	"gx/ipfs/QmXYjuNuxVzXKJCfWasQk1RqkhVLDM9jtUKhqc2WPQmFSB/go-libp2p-peer"

	"github.com/ipfs/go-ipfs/core"
	"github.com/op/go-logging"
	"github.com/digitalrupee-project/openbazaar-go/ipfs"
	"github.com/digitalrupee-project/openbazaar-go/repo"
	"golang.org/x/net/context"
)

var log = logging.MustGetLogger("service")

const kRepointFrequency = time.Hour * 12
const kPointerExpiration = time.Hour * 24 * 30

type PointerRepublisher struct {
	ipfsNode    *core.IpfsNode
	db          repo.Datastore
	pushNodes   []peer.ID
	isModerator func() bool
}

func NewPointerRepublisher(node *core.IpfsNode, database repo.Datastore, pushNodes []peer.ID, isModerator func() bool) *PointerRepublisher {
	return &PointerRepublisher{
		ipfsNode:    node,
		db:          database,
		pushNodes:   pushNodes,
		isModerator: isModerator,
	}
}

func (r *PointerRepublisher) Run() {
	tick := time.NewTicker(kRepointFrequency)
	defer tick.Stop()
	go r.Republish()
	for range tick.C {
		go r.Republish()
	}
}

func (r *PointerRepublisher) Republish() {
	republishModerator := r.isModerator()
	pointers, err := r.db.Pointers().GetAll()
	if err != nil {
		log.Error(err)
		return
	}
	ctx := context.Background()

	for _, p := range pointers {
		switch p.Purpose {
		case ipfs.MESSAGE:
			if time.Now().Sub(p.Timestamp) > kPointerExpiration {
				r.db.Pointers().Delete(p.Value.ID)
			} else {
				go ipfs.PublishPointer(ctx, r.ipfsNode, p)
				for _, peer := range r.pushNodes {
					go ipfs.PutPointerToPeer(r.ipfsNode, context.Background(), peer, p)
				}
			}
		case ipfs.MODERATOR:
			if republishModerator {
				go ipfs.PublishPointer(ctx, r.ipfsNode, p)
			} else {
				r.db.Pointers().Delete(p.Value.ID)
			}
		default:
			continue
		}
	}
}
