package main

import (
	"bufio"
	"errors"
	"fmt"
	"git.sr.ht/~geb/dotool/xkb"
	"git.sr.ht/~geb/opt"
	"github.com/bendahl/uinput"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var Version string

func usage() {
	fmt.Println(`dotool reads commands from stdin and simulates keyboard and pointer events.

The commands are:
    key CHORD...
    keydown CHORD...
    keyup CHORD...
    type TEXT
    click left/middle/right
    buttondown left/middle/right
    buttonup left/middle/right
    wheel AMOUNT  (a positive AMOUNT is up, a negative is down)
    hwheel AMOUNT  (a positive AMOUNT is right, a negative is left)
    mouseto X Y  (where X and Y are percentages between 0.0 and 1.0)
    mousemove X Y  (where X and Y are the number of pixels to move)
    keydelay MILLISECONDS  (default: 2)
    keyhold MILLISECONDS  (default: 8)
    typedelay MILLISECONDS  (default: 2)
    typehold MILLISECONDS  (default: 8)


dotool is installed with a udev rule to allow users in group input to run
it without root permissions.

You can add yourself to group input by running:

    sudo groupadd -f input
    sudo usermod -a -G input $USER

It's foolproof to reboot to make the rule effective.


Keys can be specified by Linux names, XKB names prefixed with x:, or Linux
keycodes prefixed with k:.  The Linux names are case-insensitive, except
uppercase character keys also simulate shift.

The modifiers are: super, ctrl, alt and shift.

    echo key shift+1 x:exclam shift+k:2 | dotool


There is an initial delay registering the virtual devices, but you can keep
writing commands to the same instance or use the daemon and client, dotoold
and dotoolc.

    { echo keydown A; sleep 3; echo key H shift+1; } | dotool

    dotoold &
    echo type super | dotoolc
    echo type speedy | dotoolc


dotool will type gobbledygook if your environment has assigned it a different
keyboard layout than it's simulating keycodes for.  You can match them up
with the environment variables DOTOOL_XKB_LAYOUT and DOTOOL_XKB_VARIANT.

    echo type azerty | DOTOOL_XKB_LAYOUT=fr dotool


--list-keys
    Print the possible Linux keys and exit.

--list-x-keys
    Print the possible XKB keys and exit.

--version
    Print the version and exit.`)
}

func fatal(a ...any) {
	fmt.Fprintln(os.Stderr, "dotool:", fmt.Sprint(a...))
	os.Exit(1)
}

func warn(a ...any) {
	fmt.Fprintln(os.Stderr, "dotool: WARNING:", fmt.Sprint(a...))
}

func log(err error) {
	if err != nil {
		warn(err.Error())
	}
}

type Chord struct {
	Super, AltGr, Ctrl, Alt, Shift bool
	Key int
}

func parseChord(keymap *xkb.Keymap, chord string) (Chord, error) {
	var c Chord
	keys := strings.Split(chord, "+")

	k := keys[len(keys)-1]
	if strings.HasPrefix(k, "k:") {
		code, err := strconv.Atoi(k[2:])
		if err != nil {
			return c, errors.New("invalid keycode: " + k[2:])
		}
		c.Key = code
	} else if strings.HasPrefix(k, "x:") {
		var ok bool
		c, ok = XKeys[k[2:]]
		if !ok {
			return c, errors.New("impossible XKB key for layout: " + k[2:])
		}
	} else {
		var ok bool
		c, ok = LinuxKeys[strings.ToLower(k)]
		if !ok {
			return c, errors.New("impossible key for layout: " + k)
		}
		if len(k) == 1 && unicode.IsUpper(rune(k[0])) {
			c.Shift = true
		}
	}

	for i := 0; i < len(keys) - 1; i++ {
		switch strings.ToLower(keys[i]) {
		case "super":
			c.Super = true
		case "ctrl", "control":
			c.Ctrl = true
		case "alt":
			c.Alt = true
		case "shift":
			c.Shift = true
		default:
			return c, errors.New("unknown modifier: " + keys[i])
		}
	}

	return c, nil
}

func (c *Chord) KeyDown(kb uinput.Keyboard) {
	if c.Super {
		log(kb.KeyDown(uinput.KeyLeftmeta))
	}
	if c.AltGr {
		log(kb.KeyDown(84))
	}
	if c.Ctrl {
		log(kb.KeyDown(uinput.KeyLeftctrl))
	}
	if c.Alt {
		log(kb.KeyDown(uinput.KeyLeftalt))
	}
	if c.Shift {
		log(kb.KeyDown(uinput.KeyLeftshift))
	}
	log(kb.KeyDown(c.Key))
}

func (c *Chord) KeyUp(kb uinput.Keyboard) {
	if c.Super {
		log(kb.KeyUp(uinput.KeyLeftmeta))
	}
	if c.AltGr {
		log(kb.KeyUp(84))
	}
	if c.Ctrl {
		log(kb.KeyUp(uinput.KeyLeftctrl))
	}
	if c.Alt {
		log(kb.KeyUp(uinput.KeyLeftalt))
	}
	if c.Shift {
		log(kb.KeyUp(uinput.KeyLeftshift))
	}
	log(kb.KeyUp(c.Key))
}


func listKeys(keymap *xkb.Keymap, keys map[string]Chord) {
	for code := 1; code < 256; code++ {
		for name, chord := range keys {
			if chord.Key == code && (chord == Chord{Key: code}) {
				fmt.Println(name, code)
			}
		}
		for name, chord := range keys {
			if chord.Key == code && (chord != Chord{Key: code}) {
				fmt.Println(name, code)
			}
		}
	}
}

func cutWord(s, word string) (string, bool) {
	if s == word {
		return "", true
	}
	if strings.HasPrefix(s, word + " ") || strings.HasPrefix(s, word + "\t") {
		return s[len(word)+1:], true
	}
	return "", false
}

func main() {
	var keymap *xkb.Keymap
	{
		layout := os.Getenv("DOTOOL_XKB_LAYOUT")
		if layout == "" {
			layout = os.Getenv("XKB_DEFAULT_LAYOUT")
		}
		variant := os.Getenv("DOTOOL_XKB_VARIANT")
		if variant == "" {
			variant = os.Getenv("XKB_DEFAULT_VARIANT")
		}
		if variant != "" && layout == "" {
			// Otherwise xkbcommon just ignores the variant.
			fatal("you need to set $DOTOOL_XKB_LAYOUT or $XKB_DEFAULT_LAYOUT if the variant is set")
		}
		names := xkb.RuleNames{Layout: layout, Variant: variant}

		ctx := xkb.ContextNew(xkb.ContextNoFlags)
		defer ctx.Unref()
		keymap = ctx.KeymapNewFromNames(&names, xkb.KeymapCompileNoFlags)
		if keymap == nil {
			fatal("failed to compile keymap")
		}
		defer keymap.Unref()

		initKeys(keymap)
	}

	{
		optset := opt.NewOptionSet()

		optset.FlagFunc("h", func() error {
			usage()
			os.Exit(0)
			panic("unreachable")
		})
		optset.Alias("h", "help")

		optset.FlagFunc("list-keys", func() error {
			listKeys(keymap, LinuxKeys)
			os.Exit(0)
			panic("unreachable")
		})

		optset.FlagFunc("list-x-keys", func() error {
			listKeys(keymap, XKeys)
			os.Exit(0)
			panic("unreachable")
		})

		optset.FlagFunc("version", func() error {
			fmt.Println(Version)
			os.Exit(0)
			panic("unreachable")
		})

		err := optset.Parse(true, os.Args[1:])
		if err != nil {
			fatal(err.Error())
		}
		if len(optset.Args()) > 0 {
			fatal("there should be no arguments, commands are read from stdin")
		}
	}

	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("dotool keyboard"))
	if err != nil {
		fatal(err.Error())
	}
	defer keyboard.Close()

	const precision = 10000
	touchpad, err := uinput.CreateTouchPad("/dev/uinput", []byte("dotool absolute pointer"), 0, precision, 0, precision)
	if err != nil {
		fatal(err.Error())
	}
	defer touchpad.Close()

	mouse, err := uinput.CreateMouse("/dev/uinput", []byte("dotool relative pointer"))
	if err != nil {
		fatal(err.Error())
	}
	defer mouse.Close()

	keydelay := time.Duration(2)*time.Millisecond
	keyhold := time.Duration(8)*time.Millisecond
	typedelay := time.Duration(2)*time.Millisecond
	typehold := time.Duration(8)*time.Millisecond

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := strings.TrimLeftFunc(sc.Text(), unicode.IsSpace)
		if text == "" {
			continue
		}
		if s, ok := cutWord(text, "key"); ok {
			for _, field := range strings.Fields(s) {
				if chord, err := parseChord(keymap, field); err == nil {
					chord.KeyDown(keyboard)
					time.Sleep(keyhold)
					chord.KeyUp(keyboard)
				} else {
					warn(err.Error())
				}
				time.Sleep(keydelay)
			}
		} else if s, ok := cutWord(text, "keydown"); ok {
			for _, field := range strings.Fields(s) {
				if chord, err := parseChord(keymap, field); err == nil {
					chord.KeyDown(keyboard)
				} else {
					warn(err.Error())
				}
				time.Sleep(keydelay)
			}
		} else if s, ok := cutWord(text, "keyup"); ok {
			for _, field := range strings.Fields(s) {
				if chord, err := parseChord(keymap, field); err == nil {
					chord.KeyUp(keyboard)
				} else {
					warn(err.Error())
				}
				time.Sleep(keydelay)
			}
		} else if s, ok := cutWord(text, "keydelay"); ok {
			var d float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &d)
			if err == nil {
				keydelay = time.Duration(d)*time.Millisecond
			} else {
				warn("invalid delay: " + sc.Text())
			}
		} else if s, ok := cutWord(text, "keyhold"); ok {
			var d float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &d)
			if err == nil {
				keyhold = time.Duration(d)*time.Millisecond
			} else {
				warn("invalid hold time: " + sc.Text())
			}
		} else if s, ok := cutWord(text, "type"); ok {
			for _, r := range s {
				if sym := xkb.Utf32ToKeysym(uint32(r)); sym == 0 {
					warn("invalid character: " + string(r))
				} else {
					chord := getChord(keymap, sym)
					if chord.Key == 0 {
						warn("impossible character for layout: " + string(r))
						time.Sleep(typehold)
					} else {
						chord.KeyDown(keyboard)
						time.Sleep(typehold)
						chord.KeyUp(keyboard)
					}
				}
				time.Sleep(typedelay)
			}
		} else if s, ok := cutWord(text, "typedelay"); ok {
			var d float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &d)
			if err == nil {
				typedelay = time.Duration(d)*time.Millisecond
			} else {
				warn("invalid delay: " + sc.Text())
			}
		} else if s, ok := cutWord(text, "typehold"); ok {
			var d float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &d)
			if err == nil {
				typehold = time.Duration(d)*time.Millisecond
			} else {
				warn("invalid hold time: " + sc.Text())
			}
		} else if s, ok := cutWord(text, "click"); ok {
			for _, button := range strings.Fields(s) {
				switch button {
				case "left", "1":
					log(mouse.LeftClick())
				case "middle", "2":
					log(mouse.MiddleClick())
				case "right", "3":
					log(mouse.RightClick())
				default:
					warn("unknown button: " + button)
				}
			}
		} else if s, ok := cutWord(text, "buttondown"); ok {
			for _, button := range strings.Fields(s) {
				switch button {
				case "left", "1":
					log(mouse.LeftPress())
				case "middle", "2":
					log(mouse.MiddlePress())
				case "right", "3":
					log(mouse.RightPress())
				default:
					warn("unknown button: " + button)
				}
			}
		} else if s, ok := cutWord(text, "buttonup"); ok {
			for _, button := range strings.Fields(s) {
				switch button {
				case "left", "1":
					log(mouse.LeftRelease())
				case "middle", "2":
					log(mouse.MiddleRelease())
				case "right", "3":
					log(mouse.RightRelease())
				default:
					warn("unknown button: " + button)
				}
			}
		} else if s, ok := cutWord(text, "wheel"); ok {
			var n float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &n)
			if err == nil {
				log(mouse.Wheel(false, int32(n)))
			} else {
				warn("invalid wheel amount: " + sc.Text())
			}
		} else if s, ok := cutWord(text, "hwheel"); ok {
			var n float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &n)
			if err == nil {
				log(mouse.Wheel(true, int32(n)))
			} else {
				warn("invalid hwheel amount: " + sc.Text())
			}
		} else if s, ok := cutWord(text, "scroll"); ok {
			var n float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &n)
			if err == nil {
				log(mouse.Wheel(false, int32(-n)))
			} else {
				warn("invalid scroll amount: " + sc.Text())
			}
		} else if s, ok := cutWord(text, "mouseto"); ok {
			var x, y float64
			_, err := fmt.Sscanf(s + "\n", "%f %f\n", &x, &y)
			if err == nil {
				if x < 0.0 || x > 1.0 || y < 0.0 || y > 1.0 {
					warn("clamping percentages between 0.0 and 1.0: " + sc.Text())
					x = math.Max(0, math.Min(x, 1.0))
					y = math.Min(0, math.Min(y, 1.0))
				}
				x, y := int32(x * float64(precision)), int32(y * float64(precision))
				err := touchpad.MoveTo(x, y)
				if err != nil {
					warn(err.Error())
				}
			} else {
				warn("invalid coordinate: " + sc.Text())
			}
		} else if s, ok := cutWord(text, "mousemove"); ok {
			var x, y float64
			_, err := fmt.Sscanf(s + "\n", "%f %f\n", &x, &y)
			if err == nil {
				err := mouse.Move(int32(x), int32(y))
				if err != nil {
					warn(err.Error())
				}
			} else {
				warn("invalid coordinate: " + sc.Text())
			}
		} else {
			warn("unknown command: " + sc.Text())
		}
	}
	if sc.Err() != nil {
		fatal(sc.Err().Error())
	}
}
