# miniaudio

## TODO

- Fix on Windows.

```
panic: compileCallback: expected function with one uintptr-sized result

goroutine 1 [running]:
syscall.compileCallback({0x9f5680?, 0xc000022190?}, 0xe0?)
        C:/Program Files/Go/src/runtime/syscall_windows.go:288 +0x50c
syscall.NewCallback(...)
        C:/Program Files/Go/src/syscall/syscall_windows.go:216
github.com/ebitengine/purego.NewCallback({0x9f5680, 0xc000022190})
        C:/Users/Sam/go/pkg/mod/github.com/ebitengine/purego@v0.8.2/syscall_windows.go:41 +0x133
github.com/samborkent/miniaudio.DeviceConfig.toMA({{0xa1c6e0?, 0xb65b40?}}, 0xc00011c000)
        C:/Users/Sam/git/miniaudio/device_config.go:68 +0x22b
github.com/samborkent/miniaudio.NewDevice(...)
        C:/Users/Sam/git/miniaudio/device.go:16
main.main()
        C:/Users/Sam/git/miniaudio/cmd/main.go:16 +0xe5
exit status 2
```

- Add miniaudio as git submodule.
- Generate intermittant C file after macros.
- Remove all parts of miniaudio API which are unused.
