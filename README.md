Segment violation on ecdsa.Sign
```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x30 pc=0x4c895f]

goroutine 1 [running]:
crypto/ecdsa.Sign(0x51c6e0, 0xc0000a0180, 0xc000093f50, 0xc0000ca000, 0x40, 0x40, 0x80, 0x7, 0xc0000ca000, 0xc000093e50)
	/nix/store/wwsma3fs061cyyfpqc5zz8h2qnwgcfpg-go-1.16.9/share/go/src/crypto/ecdsa/ecdsa.go:204 +0x5f
crypto/ecdsa.(*PrivateKey).Sign(0xc000093f50, 0x51c6e0, 0xc0000a0180, 0xc0000ca000, 0x40, 0x40, 0x0, 0x0, 0x0, 0xc000093ed8, ...)
	/nix/store/wwsma3fs061cyyfpqc5zz8h2qnwgcfpg-go-1.16.9/share/go/src/crypto/ecdsa/ecdsa.go:116 +0x85
crypto/ecdsa.SignASN1(...)
	/nix/store/wwsma3fs061cyyfpqc5zz8h2qnwgcfpg-go-1.16.9/share/go/src/crypto/ecdsa/ecdsa.go:286
main.SignMessage(0x51c680, 0xc0000a01e0, 0xc000093f50, 0x0, 0x0, 0x0, 0x0, 0x0)
	/home/sammy/go/src/github.com/naveensrinivasan/ecdsafuzz/main.go:52 +0xd5
main.main()
	/home/sammy/go/src/github.com/naveensrinivasan/ecdsafuzz/main.go:23 +0x189
exit status 2
```
