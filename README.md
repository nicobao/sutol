# Sutol

s - u - t - o - l
l - o - t - u - s

`sutol` is a simple CLI test tool for Filecoin miner.

Its purpose is to test the miner implementation - see how it reacts with custom and more or less malformed deal proposal.

To use it, you need to have access to a *running lotus daemon* with *admin* access token.

For now this CLI only allows you to list the deal proposal already sent - and replay them.

## Usage

Run `sutol help` for usage info.

## Limitation

Ideally, we would like to be able to talk DIRECTLY to the miner software (i.e lotus-miner or another implementation), and NOT through lotus-dameon.
Because as is, we are limited to what lotus-daemon validates from our requests, so we don't have complete freedom as to what is actually sent to the miner.
Therefore the current version of `sutol` actually tests lotus-daemon FIRST and then, if the request passes, tests the miner.

## Further improvements

It would be nice to be able to send any JSON body to the ClientStartDeal RPC. There could be three modes:
- send a syntaxically incorrect JSON (most likely rejected by lotus-daemon)
- send a syntaxically correct JSON that is NOT respecting the [ClientStartDeal](https://lotus.filecoin.io/docs/apis/json-rpc/#clientstartdeal) JSON Schema (type mismatch or missing fields for example)
- send a syntaxically correct JSON respecting the ClientStartDeal Schema but with semantically incorrect data (wrong CID, miner price) 


