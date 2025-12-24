GOFLAGS = -trimpath -gcflags=gitlab.com/bztsrc/bootboot=-std -ldflags="-linkmode external -extld $(LD) -extldflags '$(LDFLAGS)'"
LDFLAGS = -nostdlib -n -v -static -m $(M) -T link.ld
STRIPFLAGS = -s -K mmio -K fb -K bootboot -K environment -K initstack

all: mykernel.aarch64.elf

mykernel.aarch64.elf: kernel.go
	$(eval LD := aarch64-elf-ld)
	$(eval M := aarch64elf)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=5 go build $(GOFLAGS) -o mykernel.aarch64.elf
	aarch64-elf-strip $(STRIPFLAGS) mykernel.aarch64.elf
	aarch64-elf-readelf -hls mykernel.aarch64.elf >mykernel.aarch64.txt

clean:
	rm *.o *.elf *.txt || true
