package bootboi

import "structs"

// This is a 16 byte aligned size.
// 16 byte aligned values will always have 4 bit's 0 in the lower.
//
// E.g.
// 16 = 0001 0000
// 32 = 0010 0000
// 48 = 0011 0000
// 64 = 0100 0000
//
// stores in least significant tetrad/nibble (half byte) of size
// | size (upper 60 bits) | type (lower 4 bits) |
// |----------------------|---------------------|
// 63                    4|3                    0
//
// We can use this to store the type in the lower 4 bits since the value will always be preserved with the help of the setter methods
type MMapEnt struct {
	Ptr   uint64
	_size uint64
}

func (e *MMapEnt) Size() uint64 {
	return e._size & 0xFFFFFFFFFFFFFFF0 // 60 bits of 1 and 4 bits of 0s in the lower
}

func (e *MMapEnt) Type() uint64 {
	return e._size & 0xF // (... 0000 0000 0000 1111)
}

func (e *MMapEnt) IsFree() bool {
	return (e._size & 0xF) == 1
}

type BOOTBOI struct {
	structs.HostLayout          /* Makes sure to maintain the struct field order to maintain compatibility with platform since the order of these fields are important for the data to be correctly mapped in the memory */
	Magic              uint32   /* 'BOOT' magic */
	Size               uint32   /* length of bootboot structure, minimum 128 */
	Protocol           uint8    /* 1, static addresses, see PROTOCOL_* and LOADER_* above */
	FbType             uint8    /* framebuffer type, see FB_* above */
	Numcores           uint16   /* number of processor cores */
	Bspid              uint16   /* Bootstrap processor ID (Local APIC Id on x86_64) */
	Timezone           int16    /* in minutes -1440..1440 */
	Datetime           [8]uint8 /* in BCD yyyyMMddhhmmss UTC (independent to timezone) */
	InitrdPtr          uint64   /* ramdisk image position */
	InitrdSize         uint64   /* ramdisk image size */
	FbPtr              uint64   /* framebuffer pointer */
	FbSize             uint32   /* framebuffer size */
	FbWidth            uint32   /* framebuffer dimension */
	FbHeight           uint32
	FbScanline         uint32
	Arch               arch
	Mmap               MMapEnt /* Mmap field is and should be at the 128th byte of the struct */
}

type arch struct {
	AcpiPtr uint64
	SmbiPtr uint64
	EfiPtr  uint64
	MpPtr   uint64
	_       uint64
	_       uint64
	_       uint64
	_       uint64
}
