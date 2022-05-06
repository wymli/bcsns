#!/bin/env bash

cd `dirname $0`

# docker pull ethereum/client-go:latest

# 搭建运行私有链
# docker run \
#     --rm \
#     -itd \
#     --name geth \
#     -v /etc/localtime:/etc/localtime \
#     -v /etc/timezone:/etc/timezone   \
#     -v ~/bcsns/deploy/ethereum/private_chain_data:/blockchain \
#     --entrypoint /bin/sh \
#     ethereum/client-go:alltools-stable

# geth --http --http.corsdomain="https://remix.ethereum.org" --http.api web3,eth,debug,personal,net --http.addr 0.0.0.0 --vmdebug --datadir /blockchain --dev console

# geth init --datadir /blockchain /blockchain/genesis.json
# geth --identity "bcsns-node" --datadir /blockchain --networkid 88 --http --http.addr 0.0.0.0 --http.port 8545 --port 30303 --http.api "eth,net,web3" --nodiscover --allow-insecure-unlock console

# js web3 操作指南

# eth.getBalance("0x0000000000000000000000000000000000000001")
# personal.unlockAccount(eth.accounts[0])  解锁账户,后续才可以转账

# amount = web3.toWei(2,'ether')
# eth.sendTransaction({from:eth.accounts[0], to:eth.accounts[1], value:amount})

# eth.blockNumber