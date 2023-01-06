# dotool

dotool reads commands from stdin and simulates keyboard and mouse events.
It works everywhere on Linux, including in X11, Wayland and TTYs.

It takes about half a second to register the virtual device, but it can be
kept using the daemon.

## Install From Source

With go (>=1.19) run `sudo ./install.sh`.

## Usage

dotool will require root permissions unless you are in group input.
See `dotool --help`.

This greets the world:
`echo 'type Sup, Lads!' | dotool`

This screams for three seconds:
`{ echo keydown A; sleep 3; echo key H shift+1; } | dotool`

This drags the mouse:
`printf %s\\n 'buttondown left' 'mousemove 0 100' 'buttonup left' | dotool`

The daemon and client, `dotoold` and `dotoolc`, can used to keep a persistent
virtual device for a quicker initial response:
```
dotoold &
echo 'type super' | dotoolc
echo 'type speedy' | dotoolc
```

## Numen, Chat and Contact

dotool was written for [Numen Voice Control](https://numenvoice.com)
and you're very welcome to join the Matrix chat at
[#numen:matrix.org](https://matrix.to/#/#numen:matrix.org).

You can also send questions, thoughts or patches by composing an email to
[~geb/public-inbox@lists.sr.ht](https://lists.sr.ht/~geb/public-inbox).

## Support Me

[Thank you!](https://liberapay.com/geb)

## License

GPLv3 only, see [LICENSE](./LICENSE).

Copyright (c) 2022 John Gebbie
