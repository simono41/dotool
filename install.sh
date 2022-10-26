#!/bin/sh
go build && cp -v dotool dotoolc dotoold /usr/local/bin || exit
mkdir -p /etc/udev/rules.d && cp -v 80-dotool.rules /etc/udev/rules.d || exit
udevadm trigger
