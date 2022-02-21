# Sutol

s - u - t - o - l - o - t - u - s

`sutol` is a simple CLI test tool for Filecoin miner.

Its purpose is to test the miner implementation - see how it reacts with custom and more or less malformed deal proposals.

For now this CLI:
- relies on [lotus JSON-RPC API](https://lotus.filecoin.io/docs/apis/json-rpc/)
- requires an accessible lotus-daemon node running, with an [admin token](https://lotus.filecoin.io/docs/developers/api-access/#obtaining-tokens).
- only allows to list the deal proposals already sent by the specific lotus-daemon - and replay them.

## Installation

Run `go get -u github.com/nicobao/sutol`

## Usage

Run `sutol -h` for usage info.

## Development

### Possible further improvements for lotus-daemon interaction

It would be nice to be able to send any JSON body to the ClientStartDeal RPC. There could be three modes:
- send a syntaxically incorrect JSON (most likely rejected by lotus-daemon)
- send a syntaxically correct JSON that is NOT respecting the [ClientStartDeal](https://lotus.filecoin.io/docs/apis/json-rpc/#clientstartdeal) JSON Schema (type mismatch or missing fields for example)
- send a syntaxically correct JSON respecting the ClientStartDeal Schema but with semantically incorrect data (wrong CID, miner price) 


### Task list

- [x] Use existing proposalCid as input of replay-deal instead of hard-coding one
- [ ] Make miner address optionally set by flag
- [x] Make the info we cannot know (or don't know how to know) from previous deals to be optionally set by flag
- [ ] Pretty-print info from list-deals ... or find a way to print the Cid from `lotus client list-deals`

## Limitations and long-term goal

Ideally, we would like to be able to talk DIRECTLY to the miner software (i.e lotus-miner or another implementation), and NOT through lotus-daemon.

Because as is, we are limited to what lotus-daemon validates from our requests, so we don't have complete freedom as to what is actually sent to the miner.

Therefore the current version of `sutol` actually tests lotus-daemon FIRST and then, if the request passes, tests the miner.

## License

This software is released under the [BSD-2-Clause Plus Patent](./LICENSE) license.
