#!/bin/sh
# ./install.sh [DESTDIR] [BINDIR]
version="$(git describe --long --abbrev=12 --tags --dirty 2>/dev/null || echo 1.0)"
go build -ldflags "-X main.Version=$version" || exit
mkdir -p "$1/${2:-usr/local/bin}" || exit
cp -v dotool dotoolc dotoold "$1/${2:-usr/local/bin}" || exit
mkdir -p "$1/etc/udev/rules.d" || exit
cp -v 80-dotool.rules "$1/etc/udev/rules.d" || exit
udevadm trigger
