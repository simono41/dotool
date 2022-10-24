#!/bin/sh
go build && cp -v dotool dotoolc dotoold /usr/local/bin
cp -v 80-dotool.rules /etc/udev/rules.d
udevadm trigger
