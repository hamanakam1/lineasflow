package types

import "testing"

func TestEncodeDecodeRoundTrip(t *testing.T) {
	in := MatchReceipt{SeqNum: 7, StartingHash: 9, FinalX: 1.5, FinalY: 2.5, FinalZ: 3.5, Dirty: 1}
	buf := make([]byte, MatchReceiptSize)
	in.Encode(buf)
	var out MatchReceipt
	out.Decode(buf)
	if out != in {
		t.Fatalf("mismatch: %+v vs %+v", out, in)
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	in := MatchReceipt{SeqNum: 42, StartingHash: 0xfeedface, FinalX: 1.5, FinalY: 2.5, FinalZ: 3.5, Dirty: 1}
	buf := make([]byte, MatchReceiptSize)
	var out MatchReceipt
	b.ReportAllocs()
	b.SetBytes(MatchReceiptSize)
	for i := 0; i < b.N; i++ {
		in.Encode(buf)
		out.Decode(buf)
	}
}
