
cd `dirname $0`

abigen --abi=src/abi.json --pkg=bcsns --out=bcsns.go


# curl -L https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-amd64-1.10.16-20356e57.tar.gz > geth.tgz