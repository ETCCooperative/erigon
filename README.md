# Erigon +ETC 

This repository forks the upstream default `devel` branch, adding support for the Ethereum Classic network.\
This repository's default branch is `devel+classic`.
A comparison of the two branches is visible as [`devel..devel+classic`](https://github.com/ETCCooperative/erigon/compare/devel..devel+classic).

The upstream README has been relocated [here](./README.ledgerwatch.md).
See that document for more information about the project, including
build and usage instructions.

## Running

To run `erigon` on ETC use the following.
```bash
erigon --chain=classic [other options]
```

Chain data is located by default in `~/.local/share/erigon/classic/`.\
Ethash/Etchash DAGs are located by default in `~/.local/share/erigon/classic/ethash-dags/`.

With the default sync mode, ETC requires around 125GB of space as of January 23, 2023 (block ~16.81M). 

### Docker

To run `erigon` with Docker use the following.

```bash
DOCKER_BUILDKIT=1 docker build -t erigon .
docker run -p 30303:30303 [more options] erigon erigon --chain classic [more options]
```

## Developer Implementation

The Ethereum Classic chain configuration is defined in [./params/chainspecs/classic.json](./params/chainspecs/classic.json).

This repository depends heavily on `erigon-lib`, a "common" library that is rewritten from scratch and licensed under Apache 2.0.
The forked version of this lib on which this repository depends is located at [etccooperative/erigon-lib](https://github.com/etccooperative/erigon-lib). Chain configuration logic is, for the most part, defined there; including many of the additional ETC `*Config` methods.

### Upstream Maintenance

To maintain this branch, we periodically merge the upstream `devel` branch into `devel+classic`.
This has been automated by a GitHub Action, which runs nightly.

### File naming

- Added Go files are conventionally named `*_classic.go`.
- Added Go test files are conventionally named `classic_*_test.go`. (Go test files must have the suffix `_test.go`.)
