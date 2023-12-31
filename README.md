# Directory Fuzzer For Termux
 This is the directory fuzzer tool specially created for Termux Users.

## Written in Go for brazingly fast speed

## Changelog Version 1.2:

1. Now you add negative status code. By defaul 404 is set.
2. Colorfull output for better user experience.
3. Added banner as well.
4. Added a sample wordlist. 

## ToDo:

1. working on progress bar to show how much wordlist has been completed.

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
go build .
```

### Example usage:
```bash
fuzz <url> <wordlist> <threads> <negative status codes>
```

```bash
fuzz https://example.com wordlist.txt 40 404,500
```

Access Fuzzer from any directory

```bash
cp fuzz /data/data/com.termux/files/usr/bin/fuzz
```

Now you can access the Fuzzer from any directory in the Termux.

```bash
fuzz <url> <wordlist> <threads> <negative status codes>
```

## Want more features?
### Ping me on [Telegram](https://t.me/anonShrivastav)
