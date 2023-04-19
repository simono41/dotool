# dotool

dotool reads commands from stdin and simulates keyboard and pointer events.
It works everywhere on Linux, including in X11, Wayland and TTYs.

## Install From Source

With `go` and `libxkbcommon-dev` installed, run:

    sudo ./install.sh

## Permission

dotool requires permission to `/dev/uinput` to create the virtual input
devices, and a udev rule grants this to users in group input.

You could try:

    echo type hello | dotool

and if need be, you can run:

    sudo groupadd -f input
    sudo usermod -a -G input $USER

and re-login and trigger the udev rule or just reboot.

## Usage

See `dotool --help`, but this greets the world:

    echo 'type Sup, Lads!' | dotool

and this screams for three seconds:

    { echo keydown A; sleep 3; echo key H shift+1; } | dotool

There is an initial delay registering the virtual devices, but you can
keep writing commands to the same instance or use the daemon and client,
`dotoold` and `dotoolc`.

    dotoold &
    echo type super | dotoolc
    echo type speedy | dotoolc

## Keyboard Layouts

dotool will produce gobbledygook if your environment has assigned it a
different keyboard layout than it's simulating keycodes for.  You can
match them up with the environment variables `DOTOOL_XKB_LAYOUT` and
`DOTOOL_XKB_VARIANT`.

    echo type azerty | DOTOOL_XKB_LAYOUT=fr dotool

## Numen and Contact

dotool was written for [Numen](https://numenvoice.org), which has a
[chat on Matrix](https://matrix.to/#/#numen:matrix.org) you're welcome to join.

You can also send questions, thoughts or patches by composing an email to
[~geb/public-inbox@lists.sr.ht](https://lists.sr.ht/~geb/public-inbox).

## Support Me

[Thank you!](https://liberapay.com/geb)

## License

GPLv3 only, see [LICENSE](./LICENSE).

Copyright (c) 2022-2023 John Gebbie
