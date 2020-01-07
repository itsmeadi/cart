#!/usr/bin/env bash

/etc/init.d/mysql start

go build -o app .
./app