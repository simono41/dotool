package main

import (
	"git.sr.ht/~geb/dotool/xkb"
	"github.com/bendahl/uinput"
)

var LinuxKeys = map[string]Chord{
	// Linux Only
	"micmute": Chord{Key: uinput.KeyMicmute},
	"unknown": Chord{Key: uinput.KeyUnknown},
	"alterase": Chord{Key: uinput.KeyAlterase},
	"question": Chord{Key: uinput.KeyQuestion},
	"hp": Chord{Key: uinput.KeyHp},
	"bassboost": Chord{Key: uinput.KeyBassboost},
	"play": Chord{Key: uinput.KeyPlay},
	"close": Chord{Key: uinput.KeyClose},
	"edit": Chord{Key: uinput.KeyEdit},
	"move": Chord{Key: uinput.KeyMove},
	"iso": Chord{Key: uinput.KeyIso},
	"closecd": Chord{Key: uinput.KeyClosecd},
	"deletefile": Chord{Key: uinput.KeyDeletefile},
	"setup": Chord{Key: uinput.KeySetup},
	"yen": Chord{Key: uinput.KeyYen},
	"macro": Chord{Key: uinput.KeyMacro},
	"kpjpcomma": Chord{Key: uinput.KeyKpjpcomma},
	"ro": Chord{Key: uinput.KeyRo},
	"zenkakuhankaku": Chord{Key: uinput.KeyZenkakuhankaku},

	// Potentially overwritten by initKeys()
	"esc": Chord{Key: uinput.KeyEsc},
	"1": Chord{Key: uinput.Key1},
	"2": Chord{Key: uinput.Key2},
	"3": Chord{Key: uinput.Key3},
	"4": Chord{Key: uinput.Key4},
	"5": Chord{Key: uinput.Key5},
	"6": Chord{Key: uinput.Key6},
	"7": Chord{Key: uinput.Key7},
	"8": Chord{Key: uinput.Key8},
	"9": Chord{Key: uinput.Key9},
	"0": Chord{Key: uinput.Key0},
	"minus": Chord{Key: uinput.KeyMinus},
	"equal": Chord{Key: uinput.KeyEqual},
	"backspace": Chord{Key: uinput.KeyBackspace},
	"tab": Chord{Key: uinput.KeyTab},
	"q": Chord{Key: uinput.KeyQ},
	"w": Chord{Key: uinput.KeyW},
	"e": Chord{Key: uinput.KeyE},
	"r": Chord{Key: uinput.KeyR},
	"t": Chord{Key: uinput.KeyT},
	"y": Chord{Key: uinput.KeyY},
	"u": Chord{Key: uinput.KeyU},
	"i": Chord{Key: uinput.KeyI},
	"o": Chord{Key: uinput.KeyO},
	"p": Chord{Key: uinput.KeyP},
	"leftbrace": Chord{Key: uinput.KeyLeftbrace},
	"rightbrace": Chord{Key: uinput.KeyRightbrace},
	"enter": Chord{Key: uinput.KeyEnter},
	"leftctrl": Chord{Key: uinput.KeyLeftctrl},
	"a": Chord{Key: uinput.KeyA},
	"s": Chord{Key: uinput.KeyS},
	"d": Chord{Key: uinput.KeyD},
	"f": Chord{Key: uinput.KeyF},
	"g": Chord{Key: uinput.KeyG},
	"h": Chord{Key: uinput.KeyH},
	"j": Chord{Key: uinput.KeyJ},
	"k": Chord{Key: uinput.KeyK},
	"l": Chord{Key: uinput.KeyL},
	"semicolon": Chord{Key: uinput.KeySemicolon},
	"apostrophe": Chord{Key: uinput.KeyApostrophe},
	"grave": Chord{Key: uinput.KeyGrave},
	"leftshift": Chord{Key: uinput.KeyLeftshift},
	"backslash": Chord{Key: uinput.KeyBackslash},
	"z": Chord{Key: uinput.KeyZ},
	"x": Chord{Key: uinput.KeyX},
	"c": Chord{Key: uinput.KeyC},
	"v": Chord{Key: uinput.KeyV},
	"b": Chord{Key: uinput.KeyB},
	"n": Chord{Key: uinput.KeyN},
	"m": Chord{Key: uinput.KeyM},
	"comma": Chord{Key: uinput.KeyComma},
	"dot": Chord{Key: uinput.KeyDot},
	"slash": Chord{Key: uinput.KeySlash},
	"rightshift": Chord{Key: uinput.KeyRightshift},
	"kpasterisk": Chord{Key: uinput.KeyKpasterisk},
	"leftalt": Chord{Key: uinput.KeyLeftalt},
	"space": Chord{Key: uinput.KeySpace},
	"capslock": Chord{Key: uinput.KeyCapslock},
	"f1": Chord{Key: uinput.KeyF1},
	"f2": Chord{Key: uinput.KeyF2},
	"f3": Chord{Key: uinput.KeyF3},
	"f4": Chord{Key: uinput.KeyF4},
	"f5": Chord{Key: uinput.KeyF5},
	"f6": Chord{Key: uinput.KeyF6},
	"f7": Chord{Key: uinput.KeyF7},
	"f8": Chord{Key: uinput.KeyF8},
	"f9": Chord{Key: uinput.KeyF9},
	"f10": Chord{Key: uinput.KeyF10},
	"numlock": Chord{Key: uinput.KeyNumlock},
	"scrolllock": Chord{Key: uinput.KeyScrolllock},
	"kp7": Chord{Key: uinput.KeyKp7},
	"kp8": Chord{Key: uinput.KeyKp8},
	"kp9": Chord{Key: uinput.KeyKp9},
	"kpminus": Chord{Key: uinput.KeyKpminus},
	"kp4": Chord{Key: uinput.KeyKp4},
	"kp5": Chord{Key: uinput.KeyKp5},
	"kp6": Chord{Key: uinput.KeyKp6},
	"kpplus": Chord{Key: uinput.KeyKpplus},
	"kp1": Chord{Key: uinput.KeyKp1},
	"kp2": Chord{Key: uinput.KeyKp2},
	"kp3": Chord{Key: uinput.KeyKp3},
	"kp0": Chord{Key: uinput.KeyKp0},
	"kpdot": Chord{Key: uinput.KeyKpdot},
	"102nd": Chord{Key: uinput.Key102Nd},
	"f11": Chord{Key: uinput.KeyF11},
	"f12": Chord{Key: uinput.KeyF12},
	"katakana": Chord{Key: uinput.KeyKatakana},
	"hiragana": Chord{Key: uinput.KeyHiragana},
	"henkan": Chord{Key: uinput.KeyHenkan},
	"katakanahiragana": Chord{Key: uinput.KeyKatakanahiragana},
	"muhenkan": Chord{Key: uinput.KeyMuhenkan},
	"kpenter": Chord{Key: uinput.KeyKpenter},
	"rightctrl": Chord{Key: uinput.KeyRightctrl},
	"kpslash": Chord{Key: uinput.KeyKpslash},
	"sysrq": Chord{Key: uinput.KeySysrq},
	"rightalt": Chord{Key: uinput.KeyRightalt},
	"linefeed": Chord{Key: uinput.KeyLinefeed},
	"home": Chord{Key: uinput.KeyHome},
	"up": Chord{Key: uinput.KeyUp},
	"pageup": Chord{Key: uinput.KeyPageup},
	"left": Chord{Key: uinput.KeyLeft},
	"right": Chord{Key: uinput.KeyRight},
	"end": Chord{Key: uinput.KeyEnd},
	"down": Chord{Key: uinput.KeyDown},
	"pagedown": Chord{Key: uinput.KeyPagedown},
	"insert": Chord{Key: uinput.KeyInsert},
	"delete": Chord{Key: uinput.KeyDelete},
	"mute": Chord{Key: uinput.KeyMute},
	"volumedown": Chord{Key: uinput.KeyVolumedown},
	"volumeup": Chord{Key: uinput.KeyVolumeup},
	"power": Chord{Key: uinput.KeyPower},
	"kpequal": Chord{Key: uinput.KeyKpequal},
	"kpplusminus": Chord{Key: uinput.KeyKpplusminus},
	"pause": Chord{Key: uinput.KeyPause},
	"scale": Chord{Key: uinput.KeyScale},
	"kpcomma": Chord{Key: uinput.KeyKpcomma},
	"hangeul": Chord{Key: uinput.KeyHangeul},
	"hanja": Chord{Key: uinput.KeyHanja},
	"leftmeta": Chord{Key: uinput.KeyLeftmeta},
	"rightmeta": Chord{Key: uinput.KeyRightmeta},
	"compose": Chord{Key: uinput.KeyCompose},
	"stop": Chord{Key: uinput.KeyStop},
	"again": Chord{Key: uinput.KeyAgain},
	"props": Chord{Key: uinput.KeyProps},
	"undo": Chord{Key: uinput.KeyUndo},
	"front": Chord{Key: uinput.KeyFront},
	"copy": Chord{Key: uinput.KeyCopy},
	"open": Chord{Key: uinput.KeyOpen},
	"paste": Chord{Key: uinput.KeyPaste},
	"find": Chord{Key: uinput.KeyFind},
	"cut": Chord{Key: uinput.KeyCut},
	"help": Chord{Key: uinput.KeyHelp},
	"menu": Chord{Key: uinput.KeyMenu},
	"calc": Chord{Key: uinput.KeyCalc},
	"sleep": Chord{Key: uinput.KeySleep},
	"wakeup": Chord{Key: uinput.KeyWakeup},
	"file": Chord{Key: uinput.KeyFile},
	"sendfile": Chord{Key: uinput.KeySendfile},
	"xfer": Chord{Key: uinput.KeyXfer},
	"prog1": Chord{Key: uinput.KeyProg1},
	"prog2": Chord{Key: uinput.KeyProg2},
	"www": Chord{Key: uinput.KeyWww},
	"msdos": Chord{Key: uinput.KeyMsdos},
	"coffee": Chord{Key: uinput.KeyCoffee},
	"direction": Chord{Key: uinput.KeyDirection},
	"cyclewindows": Chord{Key: uinput.KeyCyclewindows},
	"mail": Chord{Key: uinput.KeyMail},
	"bookmarks": Chord{Key: uinput.KeyBookmarks},
	"computer": Chord{Key: uinput.KeyComputer},
	"back": Chord{Key: uinput.KeyBack},
	"forward": Chord{Key: uinput.KeyForward},
	"ejectcd": Chord{Key: uinput.KeyEjectcd},
	"ejectclosecd": Chord{Key: uinput.KeyEjectclosecd},
	"nextsong": Chord{Key: uinput.KeyNextsong},
	"playpause": Chord{Key: uinput.KeyPlaypause},
	"previoussong": Chord{Key: uinput.KeyPrevioussong},
	"stopcd": Chord{Key: uinput.KeyStopcd},
	"record": Chord{Key: uinput.KeyRecord},
	"rewind": Chord{Key: uinput.KeyRewind},
	"phone": Chord{Key: uinput.KeyPhone},
	"config": Chord{Key: uinput.KeyConfig},
	"homepage": Chord{Key: uinput.KeyHomepage},
	"refresh": Chord{Key: uinput.KeyRefresh},
	"exit": Chord{Key: uinput.KeyExit},
	"scrollup": Chord{Key: uinput.KeyScrollup},
	"scrolldown": Chord{Key: uinput.KeyScrolldown},
	"kpleftparen": Chord{Key: uinput.KeyKpleftparen},
	"kprightparen": Chord{Key: uinput.KeyKprightparen},
	"new": Chord{Key: uinput.KeyNew},
	"redo": Chord{Key: uinput.KeyRedo},
	"f13": Chord{Key: uinput.KeyF13},
	"f14": Chord{Key: uinput.KeyF14},
	"f15": Chord{Key: uinput.KeyF15},
	"f16": Chord{Key: uinput.KeyF16},
	"f17": Chord{Key: uinput.KeyF17},
	"f18": Chord{Key: uinput.KeyF18},
	"f19": Chord{Key: uinput.KeyF19},
	"f20": Chord{Key: uinput.KeyF20},
	"f21": Chord{Key: uinput.KeyF21},
	"f22": Chord{Key: uinput.KeyF22},
	"f23": Chord{Key: uinput.KeyF23},
	"f24": Chord{Key: uinput.KeyF24},
	"playcd": Chord{Key: uinput.KeyPlaycd},
	"pausecd": Chord{Key: uinput.KeyPausecd},
	"prog3": Chord{Key: uinput.KeyProg3},
	"prog4": Chord{Key: uinput.KeyProg4},
	"dashboard": Chord{Key: uinput.KeyDashboard},
	"suspend": Chord{Key: uinput.KeySuspend},
	"fastforward": Chord{Key: uinput.KeyFastforward},
	"print": Chord{Key: uinput.KeyPrint},
	"camera": Chord{Key: uinput.KeyCamera},
	"sound": Chord{Key: uinput.KeySound},
	"email": Chord{Key: uinput.KeyEmail},
	"chat": Chord{Key: uinput.KeyChat},
	"search": Chord{Key: uinput.KeySearch},
	"connect": Chord{Key: uinput.KeyConnect},
	"finance": Chord{Key: uinput.KeyFinance},
	"sport": Chord{Key: uinput.KeySport},
	"shop": Chord{Key: uinput.KeyShop},
	"cancel": Chord{Key: uinput.KeyCancel},
	"brightnessdown": Chord{Key: uinput.KeyBrightnessdown},
	"brightnessup": Chord{Key: uinput.KeyBrightnessup},
	"media": Chord{Key: uinput.KeyMedia},
	"switchvideomode": Chord{Key: uinput.KeySwitchvideomode},
	"kbdillumtoggle": Chord{Key: uinput.KeyKbdillumtoggle},
	"kbdillumdown": Chord{Key: uinput.KeyKbdillumdown},
	"kbdillumup": Chord{Key: uinput.KeyKbdillumup},
	"send": Chord{Key: uinput.KeySend},
	"reply": Chord{Key: uinput.KeyReply},
	"forwardmail": Chord{Key: uinput.KeyForwardmail},
	"save": Chord{Key: uinput.KeySave},
	"documents": Chord{Key: uinput.KeyDocuments},
	"battery": Chord{Key: uinput.KeyBattery},
	"bluetooth": Chord{Key: uinput.KeyBluetooth},
	"wlan": Chord{Key: uinput.KeyWlan},
	"uwb": Chord{Key: uinput.KeyUwb},
	"videonext": Chord{Key: uinput.KeyVideoNext},
	"videoprev": Chord{Key: uinput.KeyVideoPrev},
	"brightnesscycle": Chord{Key: uinput.KeyBrightnessCycle},
	"brightnesszero": Chord{Key: uinput.KeyBrightnessZero},
	"displayoff": Chord{Key: uinput.KeyDisplayOff},
	"wimax": Chord{Key: uinput.KeyWimax},
	"rfkill": Chord{Key: uinput.KeyRfkill},
}

var linuxXSyms = map[string]uint32{
	"esc": 0xff1b,
	"1": 0x31,
	"2": 0x32,
	"3": 0x33,
	"4": 0x34,
	"5": 0x35,
	"6": 0x36,
	"7": 0x37,
	"8": 0x38,
	"9": 0x39,
	"0": 0x30,
	"minus": 0x2d,
	"equal": 0x3d,
	"backspace": 0xff08,
	"tab": 0xff09,
	"q": 0x71,
	"w": 0x77,
	"e": 0x65,
	"r": 0x72,
	"t": 0x74,
	"y": 0x79,
	"u": 0x75,
	"i": 0x69,
	"o": 0x6f,
	"p": 0x70,
	"leftbrace": 0x5b,
	"rightbrace": 0x5d,
	"enter": 0xff0d,
	"leftctrl": 0xffe3,
	"a": 0x61,
	"s": 0x73,
	"d": 0x64,
	"f": 0x66,
	"g": 0x67,
	"h": 0x68,
	"j": 0x6a,
	"k": 0x6b,
	"l": 0x6c,
	"semicolon": 0x3b,
	"apostrophe": 0x27,
	"grave": 0x60,
	"leftshift": 0xffe1,
	"backslash": 0x5c,
	"z": 0x7a,
	"x": 0x78,
	"c": 0x63,
	"v": 0x76,
	"b": 0x62,
	"n": 0x6e,
	"m": 0x6d,
	"comma": 0x2c,
	"dot": 0x2e,
	"slash": 0x2f,
	"rightshift": 0xffe2,
	"kpasterisk": 0xffaa,
	"leftalt": 0xffe9,
	"space": 0x20,
	"capslock": 0xffe5,
	"f1": 0xffbe,
	"f2": 0xffbf,
	"f3": 0xffc0,
	"f4": 0xffc1,
	"f5": 0xffc2,
	"f6": 0xffc3,
	"f7": 0xffc4,
	"f8": 0xffc5,
	"f9": 0xffc6,
	"f10": 0xffc7,
	"numlock": 0xff7f,
	"scrolllock": 0xff14,
	"kp7": 0xff95,
	"kp8": 0xff97,
	"kp9": 0xff9a,
	"kpminus": 0xffad,
	"kp4": 0xff96,
	"kp5": 0xff9d,
	"kp6": 0xff98,
	"kpplus": 0xffab,
	"kp1": 0xff9c,
	"kp2": 0xff99,
	"kp3": 0xff9b,
	"kp0": 0xff9e,
	"kpdot": 0xff9f,
	"102nd": 0x3c,
	"f11": 0xffc8,
	"f12": 0xffc9,
	"katakana": 0xff26,
	"hiragana": 0xff25,
	"henkan": 0xff23,
	"katakanahiragana": 0xff27,
	"muhenkan": 0xff22,
	"kpenter": 0xff8d,
	"rightctrl": 0xffe4,
	"kpslash": 0xffaf,
	"sysrq": 0xff61,
	"rightalt": 0xffea,
	"linefeed": 0xff0a,
	"home": 0xff50,
	"up": 0xff52,
	"pageup": 0xff55,
	"left": 0xff51,
	"right": 0xff53,
	"end": 0xff57,
	"down": 0xff54,
	"pagedown": 0xff56,
	"insert": 0xff63,
	"delete": 0xffff,
	"mute": 0x1008ff12,
	"volumedown": 0x1008ff11,
	"volumeup": 0x1008ff13,
	"power": 0x1008ff2a,
	"kpequal": 0xffbd,
	"kpplusminus": 0xb1,
	"pause": 0xff13,
	"scale": 0x1008ff4a,
	"kpcomma": 0xffae,
	"hangeul": 0xff31,
	"hanja": 0xff34,
	"leftmeta": 0xffeb,
	"rightmeta": 0xffec,
	"compose": 0xff67,
	"stop": 0xff69,
	"again": 0xff66,
	"props": 0x1005ff70,
	"undo": 0xff65,
	"front": 0x1005ff71,
	"copy": 0x1008ff57,
	"open": 0x1008ff6b,
	"paste": 0x1008ff6d,
	"find": 0xff68,
	"cut": 0x1008ff58,
	"help": 0xff6a,
	"menu": 0x1008ff65,
	"calc": 0x1008ff1d,
	"sleep": 0x1008ff2f,
	"wakeup": 0x1008ff2b,
	"file": 0x1008ff5d,
	"sendfile": 0x1008ff7b,
	"xfer": 0x1008ff8a,
	"prog1": 0x1008ff41,
	"prog2": 0x1008ff42,
	"www": 0x1008ff2e,
	"msdos": 0x1008ff5a,
	"coffee": 0x1008ff2d,
	"direction": 0x1008ff74,
	"cyclewindows": 0x1008ff7f,
	"mail": 0x1008ff19,
	"bookmarks": 0x1008ff30,
	"computer": 0x1008ff33,
	"back": 0x1008ff26,
	"forward": 0x1008ff27,
	"ejectcd": 0x1008ff2c,
	"ejectclosecd": 0x1008ff2c,
	"nextsong": 0x1008ff17,
	"playpause": 0x1008ff14,
	"previoussong": 0x1008ff16,
	"stopcd": 0x1008ff15,
	"record": 0x1008ff1c,
	"rewind": 0x1008ff3e,
	"phone": 0x1008ff6e,
	"config": 0x1008ff81,
	"homepage": 0x1008ff18,
	"refresh": 0x1008ff73,
	"exit": 0x1008ff56,
	"scrollup": 0x1008ff78,
	"scrolldown": 0x1008ff79,
	"kpleftparen": 0x28,
	"kprightparen": 0x29,
	"new": 0x1008ff68,
	"redo": 0xff66,
	"f13": 0xffca,
	"f14": 0xffcb,
	"f15": 0xffcc,
	"f16": 0xffcd,
	"f17": 0xffce,
	"f18": 0xffcf,
	"f19": 0xffd0,
	"f20": 0xffd1,
	"f21": 0xffd2,
	"f22": 0xffd3,
	"f23": 0xffd4,
	"f24": 0xffd5,
	"playcd": 0x1008ff14,
	"pausecd": 0x1008ff31,
	"prog3": 0x1008ff43,
	"prog4": 0x1008ff44,
	"dashboard": 0x1008ff4b,
	"suspend": 0x1008ffa7,
	"fastforward": 0x1008ff97,
	"print": 0xff61,
	"camera": 0x1008ff8f,
	"sound": 0x1008ffb6,
	"email": 0x1008ff19,
	"chat": 0x1008ff8e,
	"search": 0x1008ff1b,
	"connect": 0x1008ff5f,
	"finance": 0x1008ff3c,
	"sport": 0x1008ff5e,
	"shop": 0x1008ff36,
	"cancel": 0xff69,
	"brightnessdown": 0x1008ff03,
	"brightnessup": 0x1008ff02,
	"media": 0x1008ff32,
	"switchvideomode": 0x1008ff59,
	"kbdillumtoggle": 0x1008ff04,
	"kbdillumdown": 0x1008ff06,
	"kbdillumup": 0x1008ff05,
	"send": 0x1008ff7b,
	"reply": 0x1008ff72,
	"forwardmail": 0x1008ff90,
	"save": 0x1008ff77,
	"documents": 0x1008ff5b,
	"battery": 0x1008ff93,
	"bluetooth": 0x1008ff94,
	"wlan": 0x1008ff95,
	"uwb": 0x1008ff96,
	"videonext": 0x1008fe22,
	"videoprev": 0x1008fe23,
	"brightnesscycle": 0x1008ff07,
	"brightnesszero": 0x100810f4,
	"displayoff": 0x100810f5,
	"wimax": 0x1008ffb4,
	"rfkill": 0x1008ffb5,
}

var XKeys = map[string]Chord{}

func newChord(keymap *xkb.Keymap, mask, code uint32) Chord{
	altGrMask := uint32(1) << keymap.ModGetIndex("Mod5")
	ctrlMask := uint32(1) << keymap.ModGetIndex(xkb.ModNameCtrl)
	altMask := uint32(1) << keymap.ModGetIndex(xkb.ModNameAlt)
	shiftMask := uint32(1) << keymap.ModGetIndex(xkb.ModNameShift)
	return Chord{
		false,
		(mask & altGrMask) != 0,
		(mask & ctrlMask) != 0,
		(mask & altMask) != 0,
		(mask & shiftMask) != 0,
		int(code) - 8,
	}
}

func initKeys(keymap *xkb.Keymap) {
	maxCode := keymap.MaxKeycode()
	for code := keymap.MinKeycode(); code <= maxCode && code < 256; code++ {
		numLevels := keymap.NumLevelsForKey(code, 0)
		for level := uint32(0); level < numLevels; level++ {
			for _, sym := range keymap.KeyGetSymsByLevel(code, 0, level) {
				chord := newChord(keymap, keymap.KeyGetMod(code, 0, level), code)
				for name, s := range linuxXSyms {
					if s == sym {
						LinuxKeys[name] = chord
					}
				}
				XKeys[xkb.KeysymGetName(sym)] = chord
			}
		}
	}
}

func getChord(keymap *xkb.Keymap, keysym uint32) Chord {
	maxCode := keymap.MaxKeycode()
	for code := keymap.MinKeycode(); code <= maxCode && code < 256; code++ {
		numLevels := keymap.NumLevelsForKey(code, 0)
		for level := uint32(0); level < numLevels; level++ {
			for _, sym := range keymap.KeyGetSymsByLevel(code, 0, level) {
				if sym == keysym {
					return newChord(keymap, keymap.KeyGetMod(code, 0, level), code)
				}
			}
		}
	}
	return Chord{}
}
