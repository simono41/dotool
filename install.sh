#!/bin/sh
# ./install.sh [DESTDIR] [BINDIR]
: "${DOTOOL_VERSION=$(git describe --long --abbrev=12 --tags --dirty 2>/dev/null || echo 1.2)}"
go build -ldflags "-X main.Version=$DOTOOL_VERSION" || exit
mkdir -p "$1/${2:-usr/local/bin}" || exit
cp -v dotool "$1/${2:-usr/local/bin}" || exit
mkdir -p "$1/etc/udev/rules.d" || exit
cp -v 80-dotool.rules "$1/etc/udev/rules.d" || exit
./_install.sh "$1" "$2" || exit

# Make the new/updated udev rule effective
udevadm control --reload
udevadm trigger
