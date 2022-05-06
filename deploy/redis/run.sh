#!/usr/bin/env bash
cd `dirname $0`

docker run -d \
      --rm \
      --name redis \
      -p 6379:6379 \
      --network test_net \
      redis


# opt=$1

# case $opt in
#   r* | run)
#     docker run -d \
#       --rm \
#       --name redis \
#       -p 6379:6379 \
#       --network bridge \
#       redis
#     ;;
#   s* | stop)
#     docker stop redis
#     ;;
#   c* | cli)
#     ./redis-cli
#     ;;
#   b* | benchmark)
#     ./redis-benchmark -q -n 100000
#     ;;
#   *)
#     echo "i don't understand; run with options: run,stop,cli,benchmark"
#     ;;
# esac