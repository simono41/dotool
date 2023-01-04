package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bendahl/uinput"
	"git.sr.ht/~geb/opt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func fatal(v ...interface{}) {
	fmt.Fprint(os.Stderr, "dotool: ")
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func warn(v ...interface{}) {
	fmt.Fprint(os.Stderr, "dotool WARNING: ")
	fmt.Fprintln(os.Stderr, v...)
}

func log(err error) {
	if err != nil {
		warn(err.Error())
	}
}

type Chord struct {
	Super bool
	Ctrl bool
	Alt bool
	Shift bool
	Key int
}

func parseChord(chord string) (Chord, error) {
	var c Chord
	keys := strings.Split(chord, "+")
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

	k := keys[len(keys)-1]
	if strings.HasPrefix(k, "k:") {
		code, err := strconv.Atoi(k[2:])
		if err != nil {
			return c, errors.New("invalid keycode: " + k)
		}
		c.Key = code
	} else if strings.HasPrefix(k, "x:") {
		if code, ok := xKeysShifted[strings.ToLower(k[2:])]; ok {
			if len(k[2:]) > 1 || unicode.IsUpper(rune(k[2])) {
				c.Shift = true
			}
			c.Key = code
		} else if code, ok := xKeysNormal[strings.ToLower(k[2:])]; ok {
			c.Key = code
		} else {
			return c, errors.New("unknown X11 key: " + k)
		}
	} else {
		code, ok := linuxKeys[strings.ToLower(k)]
		if len(k) == 1 && unicode.IsUpper(rune(k[0])) {
			c.Shift = true
		}
		if !ok {
			return c, errors.New("unknown key: " + k)
		}
		c.Key = code
	}
	return c, nil
}

func (c *Chord) KeyDown(kb uinput.Keyboard) {
	if c.Super {
		log(kb.KeyDown(uinput.KeyLeftmeta))
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


func usage() {
	fmt.Fprintln(os.Stderr, `dotool reads commands from stdin and simulates keyboard and pointer events.

The commands are:
	key CHORD...
	keydown CHORD...
	keyup CHORD...
	type TEXT
	click left/middle/right
	buttondown left/middle/right
	buttonup left/middle/right
	scroll NUMBER  (where NUMBER is the amount down/up if positive/negative)
	mouseto X Y  (where X and Y are percentages between 0.0 and 1.0)
	mousemove X Y  (where X and Y are the number of pixels to move)
	keydelay MILLISECONDS
	keyhold MILLISECONDS
	typedelay MILLISECONDS
	typehold MILLISECONDS

Example: echo "key h i shift+1" | dotool

dotool is installed with a udev rule to allow users in group input to run
it without root permissions. You can make it effective without rebooting by
running: sudo udevadm trigger

The keys are those used by Linux, but can also be specified using X11 names
prefixed with x: like x:exclam, as well as their Linux keycode like k:30.
They are case insensitive, except uppercase character keys also simulate shift.

The modifiers are: super, ctrl, alt and shift.

The daemon and client, dotoold and dotoolc, can used to keep a persistent
virtual device for a quicker initial response.

--list-keys
	Print the supported Linux keys and their keycodes.

--list-x-keys
	Print the supported X11 keys and their Linux keycodes.
`)
}

func cutCmd(s, cmd string) (string, bool) {
	if strings.HasPrefix(s, cmd + " ") || strings.HasPrefix(s, cmd + "\t") {
		return s[len(cmd)+1:], true
	}
	return "", false
}

func main() {
	{
		optset := opt.NewOptionSet()

		optset.FlagFunc("h", func() error {
			usage()
			os.Exit(0)
			panic("unreachable")
		})
		optset.Alias("h", "help")

		optset.BoolFunc("list-keys", func(bool) error {
			for i := 1; i < 249; i++ {
				for name, code := range linuxKeys {
					if code == i {
						fmt.Printf("%s %d\n", name, code)
						break
					}
				}
			}
			os.Exit(0)
			panic("unreachable")
		})

		optset.BoolFunc("list-x-keys", func(bool) error {
			for i := 1; i < 249; i++ {
				for name, code := range xKeysNormal {
					if code == i {
						fmt.Print(name)
						for name, code := range xKeysShifted {
							if code == i {
								if len(name) == 1 {
									name = strings.ToUpper(name)
								}
								fmt.Printf(" %s", name)
								break
							}
						}
						fmt.Printf(" %d\n", code)
						break
					}
				}
			}
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

	keydelay := time.Duration(0)
	keyhold := time.Duration(0)
	typedelay := time.Duration(0)
	typehold := time.Duration(0)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := strings.TrimLeftFunc(sc.Text(), unicode.IsSpace)
		if text == "" {
			continue
		}
		if s, ok := cutCmd(text, "key"); ok {
			for _, field := range strings.Fields(s) {
				if chord, err := parseChord(field); err == nil {
					chord.KeyDown(keyboard)
					time.Sleep(keyhold)
					chord.KeyUp(keyboard)
				} else {
					warn(err.Error())
				}
				time.Sleep(keydelay)
			}
		} else if s, ok := cutCmd(text, "keydown"); ok {
			for _, field := range strings.Fields(s) {
				if chord, err := parseChord(field); err == nil {
					chord.KeyDown(keyboard)
				} else {
					warn(err.Error())
				}
				time.Sleep(keydelay)
			}
		} else if s, ok := cutCmd(text, "keyup"); ok {
			for _, field := range strings.Fields(s) {
				if chord, err := parseChord(field); err == nil {
					chord.KeyUp(keyboard)
				} else {
					warn(err.Error())
				}
				time.Sleep(keydelay)
			}
		} else if s, ok := cutCmd(text, "keydelay"); ok {
			var d float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &d)
			if err == nil {
				keydelay = time.Duration(d)*time.Millisecond
			} else {
				warn("invalid delay: " + sc.Text())
			}
		} else if s, ok := cutCmd(text, "keyhold"); ok {
			var d float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &d)
			if err == nil {
				keyhold = time.Duration(d)*time.Millisecond
			} else {
				warn("invalid hold time: " + sc.Text())
			}
		} else if s, ok := cutCmd(text, "type"); ok {
			for _, r := range s {
				if chord, ok := runeChords[unicode.ToLower(r)]; ok {
					if unicode.IsUpper(r) {
						chord.Shift = true
					}
					chord.KeyDown(keyboard)
					time.Sleep(typehold)
					chord.KeyUp(keyboard)
				}
				time.Sleep(typedelay)
			}
		} else if s, ok := cutCmd(text, "typedelay"); ok {
			var d float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &d)
			if err == nil {
				typedelay = time.Duration(d)*time.Millisecond
			} else {
				warn("invalid delay: " + sc.Text())
			}
		} else if s, ok := cutCmd(text, "typehold"); ok {
			var d float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &d)
			if err == nil {
				typehold = time.Duration(d)*time.Millisecond
			} else {
				warn("invalid hold time: " + sc.Text())
			}
		} else if s, ok := cutCmd(text, "click"); ok {
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
		} else if s, ok := cutCmd(text, "buttondown"); ok {
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
		} else if s, ok := cutCmd(text, "buttonup"); ok {
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
		} else if s, ok := cutCmd(text, "scroll"); ok {
			var n float64
			_, err := fmt.Sscanf(s + "\n", "%f\n", &n)
			if err == nil {
				log(mouse.Wheel(false, int32(-n)))
			} else {
				warn("invalid scroll amount: " + sc.Text())
			}
		} else if s, ok := cutCmd(text, "mouseto"); ok {
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
		} else if s, ok := cutCmd(text, "mousemove"); ok {
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
