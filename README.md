# dotool

dotool reads commands from stdin and simulates keyboard and mouse events.
It works everywhere on Linux, including in X11, Wayland and TTYs.

It takes about half a second to register the virtual device, but it can be kept using the daemon.

## Install From Source

With go (>=1.19) run `sudo ./install.sh`.

## Usage

dotool will require root permissions unless you are in group input.

This greets the world:
`echo 'type Sup, Lads!' | dotool`

This screams for three seconds:
`{ echo keydown A; sleep 3; echo key H shift+1; } | dotool`

This drags the mouse:
`printf %s\\n 'buttondown left' 'mousemove 0 100' 'buttonup left' | dotool`

The daemon and client, `dotoold` and `dotoolc`, can used to keep a persistent virtual device for a quicker initial response:
```
dotoold &
echo 'type super' | dotoolc
echo 'type speedy' | dotoolc
```

## Contact

You can ask a question or send a patch by composing an email to [~geb/public-inbox@lists.sr.ht](https://lists.sr.ht/~geb/public-inbox).
