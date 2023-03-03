![Logo](public/logo.png)

**Merkle is building next generation MEV infrastructure**. [Join us on discord](https://discord.gg/Q9Dc7jVX6c).

# Merkle transactions

Merkle transactions is a service that streams transactions harvested by Merkle's private transaction network. We run a fleet of nodes listening to transactions on Ethereum mainnet.

On average, our network gets a block coverage of ~98% and catches transactions 6 seconds before they are mined. All transactions are guaranteed to be valid, no duplicates, no low nonces, no wrong chain id.

## How it works

The Merkle transaction broker is a gRPC service running at 

```
https://txs.usemerkle.com
```

To test it, clone this repo and run `make listen`. Checkout [the example on how to listen for transactions](cmd/test/main.go).

**Merkle transations is free while in beta, pricing will be announced soon. Beta users will get preferred pricing.**