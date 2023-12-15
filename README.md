# Directory Fuzzer For Termux
 This is the directory fuzzer tool specially created for Termux Users.

## Written in Go for brazingly fast speed

## How to use?

### Installation:
```bash
pkg install golang git
```

```bash
git clone https://github.com/sudityashrivastav/Directory-Fuzzer-For-Termux
```

```bash
cd Directory-Fuzzer-For-Termux
```

```bash
go build fuzzer.go
```

### Example usage:
```bash
./fuzzer <url> <wordlist> <threads> <status codes>
```

```bash
./fuzzer https://example.com wordlist.txt 40 200,206
```

Access Fuzzer from any directory

```bash
cp fuzzer /data/data/com.termux/files/usr/bin/fuzzer
```

Now you can access the Fuzzer from any directory in the Termux.

```bash
fuzzer <url> <wordlist> <threads> <status codes>
```

## Want more features?
### Ping me on [Telegram](https://t.me/anonShrivastav)
