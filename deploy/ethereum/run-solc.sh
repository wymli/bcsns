# 安装不了
# sudo add-apt-repository ppa:ethereum/ethereum
# sudo apt-get update
# sudo apt-get install solc


docker run \
  --rm \
  -it \
  --name solc \
  -v ~/bcsns/blockchain/src:/blockchain/src \
  ethereum/solc:stable solc --abi /blockchain/src/contract.sol