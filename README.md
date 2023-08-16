# dotool

dotool reads actions from stdin and simulates keyboard/mouse input using
Linux's uinput module. It works system-wide, including in X11, Wayland
and TTYs.

## Install From Source

With `go`, `libxkbcommon-dev` and `scdoc` installed, run:

    ./build.sh && sudo ./build.sh install

And to trigger the udev rule, run:

    sudo udevadm control --reload && sudo udevadm trigger

## Usage

See the [manpage](doc/dotool.1.scd).

## Numen and Contact

dotool was written for [Numen](https://numenvoice.org), which has a
[chat on Matrix](https://matrix.to/#/#numen:matrix.org) you're welcome to join.

You can also send questions or patches by composing an email to
[~geb/public-inbox@lists.sr.ht](https://lists.sr.ht/~geb/public-inbox).

## Support My Work 👀

[Thank you!](https://liberapay.com/geb)

## License

GPLv3 only, see [LICENSE](./LICENSE).

Copyright (c) 2022-2023 John Gebbie
