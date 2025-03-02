# myblockchain

**myblockchain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project.

For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/myblockchain@latest! | sudo bash
```

`username/myblockchain` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

## Consensus-Breaking Changes

A **consensus-breaking change** refers to a modification to the blockchain protocol that causes nodes running an older version of the software to be unable to understand new blocks or transactions produced by nodes running the new version. In a blockchain network, all nodes must agree on the state of the blockchain, and this agreement is what we call consensus. When a change is made that is not backward compatible, it breaks the consensus because the old nodes will fail to understand the new block structure or transaction format.

## Why This Change Breaks Consensus

In this update, the format of the `ID` field for the `Resource` has been changed from a **string** to an **integer**. This is a consensus-breaking change because:

- Older nodes (those running the previous version) expect the `ID` to be a string.
- When they receive a new block containing resources with `ID` as an integer, they won’t be able to decode the transaction properly.
- As a result, older nodes will fail to interpret the block or transaction, causing them to fall out of sync with the network.

Thus, nodes running the older version of the software will no longer be able to correctly process data produced by nodes running the updated version, breaking the consensus between the two.
# blockchain-project
