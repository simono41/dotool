#!/bin/sh
# ./install.sh [DESTDIR] [BINDIR]
: "${DOTOOL_VERSION=$(git describe --long --abbrev=12 --tags --dirty 2>/dev/null || echo 1.1)}"
go build -ldflags "-X main.Version=$DOTOOL_VERSION" || exit
mkdir -p "$1/${2:-usr/local/bin}" || exit
cp -v dotool dotoolc dotoold "$1/${2:-usr/local/bin}" || exit
mkdir -p "$1/etc/udev/rules.d" || exit
cp -v 80-dotool.rules "$1/etc/udev/rules.d" || exit
udevadm trigger
