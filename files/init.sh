#!/usr/bin/env bash

/etc/init.d/mysql start
mysql -u root < files/db.sql

