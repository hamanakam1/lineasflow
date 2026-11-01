package types

import (
	"encoding/binary"
	"math"
)

// MatchReceipt mirrors Bumperfort's C-ABI match_receipt_t.
type MatchReceipt struct {
	SeqNum      uint64
	StartingHash uint64
	FinalX      float64
	FinalY      float64
	FinalZ      float64
	Dirty       uint32
	_           uint32
}

const MatchReceiptSize = 48

func (m *MatchReceipt) Encode(b []byte) {
	binary.LittleEndian.PutUint64(b[0:], m.SeqNum)
	binary.LittleEndian.PutUint64(b[8:], m.StartingHash)
	binary.LittleEndian.PutUint64(b[16:], math.Float64bits(m.FinalX))
	binary.LittleEndian.PutUint64(b[24:], math.Float64bits(m.FinalY))
	binary.LittleEndian.PutUint64(b[32:], math.Float64bits(m.FinalZ))
	binary.LittleEndian.PutUint32(b[40:], m.Dirty)
}

func (m *MatchReceipt) Decode(b []byte) {
	m.SeqNum = binary.LittleEndian.Uint64(b[0:])
	m.StartingHash = binary.LittleEndian.Uint64(b[8:])
	m.FinalX = math.Float64frombits(binary.LittleEndian.Uint64(b[16:]))
	m.FinalY = math.Float64frombits(binary.LittleEndian.Uint64(b[24:]))
	m.FinalZ = math.Float64frombits(binary.LittleEndian.Uint64(b[32:]))
	m.Dirty = binary.LittleEndian.Uint32(b[40:])
}
