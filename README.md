## Go bubble

Welcome to the bubble source code repository! This is an Ethereum-based、high-performance and high-security implementation of the bubble protocol.
Most of peculiarities according the bubble's **whitepaper**([English](https://www.bubble.network/pdf/en/Bubble_A_High-Efficiency_Trustless_Computing_Network_Whitepaper_EN.pdf)|[中文](https://www.bubble.network/pdf/zh/Bubble_A_High-Efficiency_Trustless_Computing_Network_Whitepaper_ZH.pdf)) has been developed.

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/bubblenet/bubble?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/bubblenet/bubble)](https://goreportcard.com/report/github.com/bubblenet/bubble)
[![Build Status](https://github.com/bubblenet/bubble/workflows/unittest/badge.svg)](https://github.com/bubblenet/bubble/actions)
[![codecov](https://codecov.io/gh/bubblenet/bubble/branch/feature-mainnet-launch/graph/badge.svg)](https://codecov.io/gh/bubblenet/bubble)
[![version](https://img.shields.io/github/v/tag/bubblenet/bubble)](https://github.com/bubblenet/bubble/releases/latest)
[![GitHub All Releases](https://img.shields.io/github/downloads/bubblenet/bubble/total.svg)](https://github.com/bubblenet/bubble)

## Building the source
The requirements to build bubble are:

- OS:Windows10/Ubuntu18.04
- [Golang](https://golang.org/doc/install) :version 1.16+
- [cmake](https://cmake.org/) :version 3.0+
- [g++&gcc](http://gcc.gnu.org/) :version 7.4.0+
> 'cmake' and 'gcc&g++' are usually built-in with Ubuntu

In addition, the following libraries needs to be installed manually

```
sudo apt install libgmp-dev libssl-dev
```
Then, clone the repository and download dependency

```
git clone https://github.com/bubblenet/bubble.git --recursive

cd bubble && go mod download
```

Ubuntu:

```
make all
```

Windows:

```
go run build\ci.go install 
```

The resulting binary will be placed in 'bubble/build/bin' .

## Getting Started

The project comes with several executables found in the `build/bin` directory.

| Command    | Description |
|:----------:|-------------|
| **`bubble`** | Our main bubble CLI client. It is the entry point into the Bubble network |
| `keytool`    | a key related tool. |

### Generate the keys

Each node requires two pairs of public&private keys, the one is called node's keypair, it's generated based on the secp256k1 curve for marking the node identity and signning the block, and the other is called node's blskeypair, it's based on the BLS_12_381 curve and is used for consensus verifing. These two pairs of public-private key need to be generated by the keytool tool.

Switch to the directory where contains 'keytool.exe'(Windows) or 'keytool'(Ubuntu).
Node's keypair(Ubuntu for example):

```
keytool genkeypair
PrivateKey:  1abd1200759d4693f4510fbcf7d5caad743b11b5886dc229da6c0747061fca36
PublicKey :  8917c748513c23db46d23f531cc083d2f6001b4cc2396eb8412d73a3e4450ffc5f5235757abf9873de469498d8cf45f5bb42c215da79d59940e17fcb22dfc127
```
Node's blskeypair:：

```
keytool genblskeypair
PrivateKey:  7747ec6876bbf8ca0934f05e45917b4213afc5814639355868bbf06d0b3e0f19
PublicKey :  e5eb9915ed2b5fd52cf5ff760873a75a8562956e176968f3cbe5ea2b22e03a7b5efc07fdd5ad66d433b404cb880b560bed6295fa79f8fa649588be02231de2e70a782751dc28dbf516b7bb5d52053b5cdf985d8961a5baafa467e8dda55fe981
```

> Note: The PublicKey generated by the 'genkeypair' command is the ***NodeID*** we needed, the PrivateKey is the corresponding ***node private key***, and the PublicKey generated by the 'genblskeypair' command is the node ***BLS PublicKey***, used in the staking and consensus process, PrivateKey is the ***Node BLS PrivateKey***, these two keypairs are common in different operating systems, that is, the public and private keys generated in Windows above, can be used in Ubuntu.

store the two private keys in files:

```
mkdir -p ./data
touch ./data/nodekey 
echo "{your-nodekey}" > ./data/nodekey
touch ./data/blskey
echo "{your-blskey}" > ./data/blskey
```

### Generate a wallet

```
bubble --datadir ./data account new
Your new account is locked with a password. Please give a password. Do not forget this password.
Passphrase:
Repeat passphrase:
Address: {lat1anp4tzmdggdrcf39qvshfq3glacjxcd5k60wg9}
```

> Do remember the password

### Connect to the Bubble network

| Options | description |
| :------------ | :------------ |
| --identity | Custom node name |
| --datadir  | Data directory for the databases and keystore |
| --rpcaddr  | HTTP-RPC server listening interface (default: "localhost") |
| --rpcport  | HTTP-RPC server listening port (default: 6789) |
| --rpcapi   | API's offered over the HTTP-RPC interface |
| --rpc      | Enable the HTTP-RPC server |
| --nodiscover | Disables the peer discovery mechanism (manual peer addition) |
| --nodekey | P2P node key file |
| --cbft.blskey | BLS key file |
| --op.prikey | The Bubble sub-chain operates the node address private key, which is used to interact with the main chain |
| --proxy.rpc.port | rpc proxy port of Bubble sub-chain operator node on main chain operator node |
| --allow_ports | An open port range is allowed for port allocation when micro node nat holes are punched |

Run the following command to launch a Bubble node connecting to the Bubble's mainnet:

```bash
bubble --identity "bubble" --datadir ./data --port {your-p2p-port} --rpcaddr 127.0.0.1 --rpcport {your-rpc-port} --rpcapi "bubble,net,web3,admin,personal" --rpc --nodiscover --nodekey ./data/nodekey --cbft.blskey ./data/blskey --proxy.rpc.port 20000 --allow_ports 30000-35000
```

OK, it seems that the chain is running correctly, we can check it as follow:

```
bubble attach http://127.0.0.1:6789
Welcome to the Bubble JavaScript console!

instance: bubblenet/bubble/v0.7.3-unstable/linux-amd64/go1.17
at block: 26 (Wed, 15 Dec 51802 20:22:44 CST)
 datadir: /home/develop/bubble/data
 modules: admin:1.0 debug:1.0 miner:1.0 net:1.0 personal:1.0 bub:1.0 rpc:1.0 txgen:1.0 txpool:1.0 web3:1.0

> bub.blockNumber
29
```

For more information, please visit our [Docs](https://devdocs.bubble.network/docs/en/).

## Contributing to bubble

All of codes for bubble are open source and contributing are very welcome! Before beginning, please take a look at our contributing [guidelines](https://github.com/bubblenet/bubble/blob/develop/.github/CONTRIBUTING.md). You can also open an issue by clicking [here](https://github.com/bubblenet/bubble/issues/new/choose).

## Support
If you have any questions or suggestions please contact us at support@bubble.network.

## License
The bubble library (i.e. all code outside of the cmd directory) is licensed under the GNU Lesser General Public License v3.0, also included in our repository in the COPYING.LESSER file.

The bubble binaries (i.e. all code inside of the cmd directory) is licensed under the GNU General Public License v3.0, also included in our repository in the COPYING file.

