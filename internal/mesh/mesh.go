package mesh

import (
	"net"
	"time"

	"github.com/hamanakam1/lineasflow/pkg/types"
)

type Broadcaster struct {
	conn *net.UDPConn
	addr *net.UDPAddr
}

func NewBroadcaster(group string, port int) (*Broadcaster, error) {
	addr, err := net.ResolveUDPAddr("udp4", group+":"+itoa(port))
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return nil, err
	}
	return &Broadcaster{conn: conn, addr: addr}, nil
}

func (b *Broadcaster) Send(r *types.MatchReceipt) error {
	buf := make([]byte, types.MatchReceiptSize) // allocation; hot path ideally reuse
	r.Encode(buf)
	_, err := b.conn.Write(buf)
	return err
}

func (b *Broadcaster) SendBytes(buf []byte) error {
	_, err := b.conn.Write(buf)
	return err
}

func (b *Broadcaster) Close() error { return b.conn.Close() }

func itoa(n int) string {
	// tiny
	if n == 0 {
		return "0"
	}
	digits := []byte{}
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}

var _ = time.Second
