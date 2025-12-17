# ARPspoofer

A **reliable ARP spoofing / MITM tool written in Go**, focused on correctness, safety, and proper cleanup.
Designed for **educational use and authorized network testing only**.

This tool poisons ARP caches between a **target** and a **host/gateway**, maintains the MITM state, and optionally restores ARP tables on exit.

---

## Features

* Targeted ARP spoofing (host ↔ target)
* Automatic ARP table restoration on exit
* Interface validation (up, exists, owns IP)
* Hardware address discovery via ARP
* Sysctl safety checks (Linux & macOS)
* Graceful shutdown with signal handling
* Verbose packet-level logging
* Cross-platform (Linux, macOS; others skip sysctl checks)

---

## How It Works (high-level)

1. Validates input IPs and interface
2. Verifies required sysctl settings (forwarding, ARP behavior)
3. Opens the network interface using `pcap`
4. Discovers MAC addresses via ARP requests
5. Continuously sends forged ARP replies to:

   * Tell the **host** that you are the **target**
   * Tell the **target** that you are the **host**
6. On exit, restores legitimate ARP mappings (optional)

---

## Installation

### Requirements

* Go 1.20+
* Root privileges (raw packet access)
* `libpcap`

### Build

```bash
git clone https://github.com/Onyz107/arpspoofer
cd arpspoofer
go build -o arpspoofer ./cmd/arpspoofer
```

---

## Usage

```bash
sudo ./arpspoofer --target 192.168.1.50 --host 192.168.1.1 --interface eth0
```

### Flags

| Flag          | Alias | Description                                |
| ------------- | ----- | ------------------------------------------ |
| `--target`    | `-t`  | Target IP to spoof                         |
| `--host`      | `-g`  | Host / gateway IP                          |
| `--interface` | `-i`  | Network interface to use                   |
| `--interval`  | `-n`  | Interval between ARP packets (default: 1s) |
| `--restore`   | `-r`  | Restore ARP tables on exit (default: true) |
| `--verbose`   | —     | Enable verbose packet logging              |

---

## Example

```bash
sudo ./arpspoofer -t 192.168.1.42 -g 192.168.1.1 -i wlan0 -v
```

---

## Sysctl Requirements

On **Linux** and **macOS**, the tool enforces required sysctl values to avoid broken MITM states (e.g. no forwarding, ARP flux, strict RP filtering).

If these are misconfigured, the program **refuses to run** instead of half-working and lying to you.

---

## Project Structure

```
cmd/arpspoofer     CLI entrypoint
handle/            pcap interface abstraction
internal/arp       ARP packet construction
internal/hwid      MAC discovery via ARP
internal/sysctl    OS-specific sysctl validation
internal/banner    Startup banner
spoof/             Spoofing & restoration logic
```

---

## Safety & Ethics

⚠ **This tool performs ARP poisoning.**
Use it **only** on networks you own or have explicit permission to test.

You are responsible for misuse. The code makes no attempt to hide activity.