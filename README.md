# zenexporter

A small CLI that signs in to [ZenMoney](https://zenmoney.ru), downloads your transaction history, and saves it as a CSV file

If you're looking for a web-based service to analyze your transaction history, check out my related project: [zenalitics](https://github.com/lleWelll/zenalitics)

## Download

Prebuilt binaries for Linux, macOS and Windows are on the [Releases page](https://github.com/lleWelll/zenexporter/releases/latest)

1. Download the archive for your system:
   - macOS (Apple Silicon): `zenexport_<version>_darwin_arm64.tar.gz`
   - macOS (Intel): `zenexport_<version>_darwin_amd64.tar.gz`
   - Linux (64-bit Intel/AMD): `zenexport_<version>_linux_amd64.tar.gz`
   - Linux (ARM64): `zenexport_<version>_linux_arm64.tar.gz`
   - Windows (64-bit): `zenexport_<version>_windows_amd64.zip`
2. Extract it. The archive contains the `zenexport` binary (`zenexport.exe` on Windows).
3. On macOS or Linux, make it executable if needed: `chmod +x zenexport`.

To verify the download, compare its SHA-256 hash with the matching line in `checksums.txt`, which is attached to each release.

## Build from source

Requirements: Go (the version is set in `go.mod`).

```sh
git clone https://github.com/lleWelll/zenexporter.git
cd zenexporter
make build
```

This creates `./main` in the project root.

## Configuration

The tool needs three values. Each can be passed as a flag or read from an environment variable. If both are set, the flag wins.

| Flag       | Environment variable   | Description                          |
|------------|------------------------|--------------------------------------|
| `--login`  | `ZENMONEY_LOGIN`       | Your ZenMoney login                  |
| `--password` | `ZENMONEY_PASSWORD`  | Your ZenMoney password               |
| `--out`    | `ZENMONEY_DATA_DIR`    | Directory where the CSV will be saved |

A trailing `/` on the output directory is removed automatically.

## Usage

Using flags:

```sh
./zenexport --login=you@example.com --password='your-password' --out=./exports
```

Using environment variables:

```sh
export ZENMONEY_LOGIN=you@example.com
export ZENMONEY_PASSWORD='your-password'
export ZENMONEY_DATA_DIR=./exports
./zenexport
```

If you built from source, use `./main` instead of `./zenexport`.

On success the tool prints the path of the saved file.

If the login or output directory is missing, the tool exits with an error that names the flags and variables to set.

