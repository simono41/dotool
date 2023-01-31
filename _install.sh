#!/bin/sh
# ./_install.sh [DESTDIR] [BINDIR]
mkdir -p "$1/${2:-usr/local/bin}" || exit
cp -v dotoolc dotoold "$1/${2:-usr/local/bin}" || exit
mkdir -p "$1/usr/share/X11/xorg.conf.d" || exit
cp -v 50-dotool.conf "$1/usr/share/X11/xorg.conf.d" || exit
mkdir -p "$1/etc/sway/config.d" || exit
cp -v dotool.sway "$1/etc/sway/config.d/dotool" || exit
