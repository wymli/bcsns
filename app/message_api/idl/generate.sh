#!/bin/env bash
cd `dirname $0`
goctl api go -api ./*.api -dir ../  -style=go_zero