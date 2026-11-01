package main

import (
	"log"
	"time"

	"github.com/hamanakam1/lineasflow/internal/mesh"
	"github.com/hamanakam1/lineasflow/internal/shm"
	"github.com/hamanakam1/lineasflow/pkg/types"
)

func main() {
	b, err := shm.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	bc, err := mesh.NewBroadcaster("239.255.42.42", 4242)
	if err != nil {
		log.Fatal(err)
	}
	defer bc.Close()

	var lastSeq uint64
	var buf [types.MatchReceiptSize]byte
	for {
		r := b.Read()
		if r.Dirty != 0 && r.SeqNum != lastSeq {
			lastSeq = r.SeqNum
			r.Encode(buf[:])
			bc.SendBytes(buf[:])
		}
		time.Sleep(time.Microsecond * 100)
	}
}
