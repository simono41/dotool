# dotool

dotool reads commands from stdin and simulates keyboard and mouse events.
It works everywhere on Linux, including in X11, Wayland and TTYs.

It takes about half a second to register the virtual device, but it can be kept using the daemon.

## Install

Run `go build` with go (>=1.19) and copy `dotool`, `dotoold` and `dotoolc` into your PATH.

## Usage

dotool will usually require root permissions unless you add a udev rule for uinput.
You could run these commands to add a rule for yourself and make it effective:
```
echo KERNEL==\"uinput\", GROUP=\"$USER\", MODE:=\"0660\" | sudo tee /etc/udev/rules.d/99-dotool-$USER.rules
sudo udevadm trigger
```

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
