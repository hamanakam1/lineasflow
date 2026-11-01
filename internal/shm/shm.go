package shm

// #include <sys/mman.h>
// #include <sys/stat.h>
// #include <fcntl.h>
// #include <unistd.h>
// #include <stdlib.h>
// #include <string.h>
// static int shm_open_ro(const char* name){ return shm_open(name, O_RDONLY, 0); }
import "C"

import (
	"unsafe"

	"github.com/hamanakam1/lineasflow/pkg/types"
)

const shmName = "/bumperfort_mesh"

type Bridge struct {
	ptr unsafe.Pointer
	fd  C.int
}

func Open() (*Bridge, error) {
	cname := C.CString(shmName)
	defer C.free(unsafe.Pointer(cname))
	fd := C.shm_open_ro(cname)
	if fd < 0 {
		return nil, shmErr("shm_open")
	}
	size := C.size_t(types.MatchReceiptSize)
	ptr := C.mmap(nil, size, C.PROT_READ, C.MAP_SHARED, fd, 0)
	if uintptr(ptr) == ^uintptr(0) {
		C.close(fd)
		return nil, shmErr("mmap")
	}
	return &Bridge{ptr: ptr, fd: fd}, nil
}

func (b *Bridge) Read() types.MatchReceipt {
	// zero-copy read of raw memory
	r := (*types.MatchReceipt)(b.ptr)
	// copy out to avoid torn reads? Dirty bit check minimal
	return *r
}

func (b *Bridge) Close() {
	C.munmap(b.ptr, C.size_t(types.MatchReceiptSize))
	C.close(b.fd)
}

func shmErr(s string) error { return &shmError{s} }

type shmError struct{ op string }

func (e *shmError) Error() string { return e.op + " failed" }
