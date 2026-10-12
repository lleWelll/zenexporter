# zenexporter

A small CLI that signs in to [ZenMoney](https://zenmoney.ru), downloads your transaction history, and saves it as a CSV file.

## Requirements

- Go (the version is set in `go.mod`)

## Installation

Clone the repository and build the binary:

```sh
git clone https://github.com/<your-user>/zenexporter.git
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
./main --login=you@example.com --password='your-password' --out=./exports
```

Using environment variables:

```sh
export ZENMONEY_LOGIN=you@example.com
export ZENMONEY_PASSWORD='your-password'
export ZENMONEY_DATA_DIR=./exports
./main
```

On success the tool prints the path of the saved file.

If the login or output directory is missing, the tool exits with an error that names the flags and variables to set.

