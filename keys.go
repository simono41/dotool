package main

import "github.com/bendahl/uinput"

var runeChords = map[rune]Chord{
	'a': Chord{Key: uinput.KeyA},
	'b': Chord{Key: uinput.KeyB},
	'c': Chord{Key: uinput.KeyC},
	'd': Chord{Key: uinput.KeyD},
	'e': Chord{Key: uinput.KeyE},
	'f': Chord{Key: uinput.KeyF},
	'g': Chord{Key: uinput.KeyG},
	'h': Chord{Key: uinput.KeyH},
	'i': Chord{Key: uinput.KeyI},
	'j': Chord{Key: uinput.KeyJ},
	'k': Chord{Key: uinput.KeyK},
	'l': Chord{Key: uinput.KeyL},
	'm': Chord{Key: uinput.KeyM},
	'n': Chord{Key: uinput.KeyN},
	'o': Chord{Key: uinput.KeyO},
	'p': Chord{Key: uinput.KeyP},
	'q': Chord{Key: uinput.KeyQ},
	'r': Chord{Key: uinput.KeyR},
	's': Chord{Key: uinput.KeyS},
	't': Chord{Key: uinput.KeyT},
	'u': Chord{Key: uinput.KeyU},
	'v': Chord{Key: uinput.KeyV},
	'w': Chord{Key: uinput.KeyW},
	'x': Chord{Key: uinput.KeyX},
	'y': Chord{Key: uinput.KeyY},
	'z': Chord{Key: uinput.KeyZ},
	'1': Chord{Key: uinput.Key1},
	'!': Chord{Key: uinput.Key1, Shift: true},
	'2': Chord{Key: uinput.Key2},
	'@': Chord{Key: uinput.Key2, Shift: true},
	'3': Chord{Key: uinput.Key3},
	'#': Chord{Key: uinput.Key3, Shift: true},
	'4': Chord{Key: uinput.Key4},
	'$': Chord{Key: uinput.Key4, Shift: true},
	'5': Chord{Key: uinput.Key5},
	'%': Chord{Key: uinput.Key5, Shift: true},
	'6': Chord{Key: uinput.Key6},
	'^': Chord{Key: uinput.Key6, Shift: true},
	'7': Chord{Key: uinput.Key7},
	'&': Chord{Key: uinput.Key7, Shift: true},
	'8': Chord{Key: uinput.Key8},
	'*': Chord{Key: uinput.Key8, Shift: true},
	'9': Chord{Key: uinput.Key9},
	'(': Chord{Key: uinput.Key9, Shift: true},
	'0': Chord{Key: uinput.Key0},
	')': Chord{Key: uinput.Key0, Shift: true},
	'`': Chord{Key: uinput.KeyGrave},
	'~': Chord{Key: uinput.KeyGrave, Shift: true},
	'-': Chord{Key: uinput.KeyMinus},
	'_': Chord{Key: uinput.KeyMinus, Shift: true},
	'=': Chord{Key: uinput.KeyEqual},
	'+': Chord{Key: uinput.KeyEqual, Shift: true},
	'[': Chord{Key: uinput.KeyLeftbrace},
	'{': Chord{Key: uinput.KeyLeftbrace, Shift: true},
	'}': Chord{Key: uinput.KeyRightbrace},
	']': Chord{Key: uinput.KeyRightbrace, Shift: true},
	'\\': Chord{Key: uinput.KeyBackslash},
	'|': Chord{Key: uinput.KeyBackslash, Shift: true},
	';': Chord{Key: uinput.KeySemicolon},
	':': Chord{Key: uinput.KeySemicolon, Shift: true},
	'\'':Chord{Key: uinput.KeyApostrophe},
	'"': Chord{Key: uinput.KeyApostrophe, Shift: true},
	',': Chord{Key: uinput.KeyComma},
	'<': Chord{Key: uinput.KeyComma, Shift: true},
	'.': Chord{Key: uinput.KeyDot},
	'>': Chord{Key: uinput.KeyDot, Shift: true},
	'/': Chord{Key: uinput.KeySlash},
	'?': Chord{Key: uinput.KeySlash, Shift: true},
	' ': Chord{Key: uinput.KeySpace},
	'\t': Chord{Key: uinput.KeyTab},
}

// generated with:
// go doc uinput.keyesc | sed '/Key/ !d; s/Key\(\S*\).*/"\L\1\E": uinput.Key\1,/'
var linuxKeys = map[string]int{
	"esc": uinput.KeyEsc,
	"1": uinput.Key1,
	"2": uinput.Key2,
	"3": uinput.Key3,
	"4": uinput.Key4,
	"5": uinput.Key5,
	"6": uinput.Key6,
	"7": uinput.Key7,
	"8": uinput.Key8,
	"9": uinput.Key9,
	"0": uinput.Key0,
	"minus": uinput.KeyMinus,
	"equal": uinput.KeyEqual,
	"backspace": uinput.KeyBackspace,
	"tab": uinput.KeyTab,
	"q": uinput.KeyQ,
	"w": uinput.KeyW,
	"e": uinput.KeyE,
	"r": uinput.KeyR,
	"t": uinput.KeyT,
	"y": uinput.KeyY,
	"u": uinput.KeyU,
	"i": uinput.KeyI,
	"o": uinput.KeyO,
	"p": uinput.KeyP,
	"leftbrace": uinput.KeyLeftbrace,
	"rightbrace": uinput.KeyRightbrace,
	"enter": uinput.KeyEnter,
	"leftctrl": uinput.KeyLeftctrl,
	"a": uinput.KeyA,
	"s": uinput.KeyS,
	"d": uinput.KeyD,
	"f": uinput.KeyF,
	"g": uinput.KeyG,
	"h": uinput.KeyH,
	"j": uinput.KeyJ,
	"k": uinput.KeyK,
	"l": uinput.KeyL,
	"semicolon": uinput.KeySemicolon,
	"apostrophe": uinput.KeyApostrophe,
	"grave": uinput.KeyGrave,
	"leftshift": uinput.KeyLeftshift,
	"backslash": uinput.KeyBackslash,
	"z": uinput.KeyZ,
	"x": uinput.KeyX,
	"c": uinput.KeyC,
	"v": uinput.KeyV,
	"b": uinput.KeyB,
	"n": uinput.KeyN,
	"m": uinput.KeyM,
	"comma": uinput.KeyComma,
	"dot": uinput.KeyDot,
	"slash": uinput.KeySlash,
	"rightshift": uinput.KeyRightshift,
	"kpasterisk": uinput.KeyKpasterisk,
	"leftalt": uinput.KeyLeftalt,
	"space": uinput.KeySpace,
	"capslock": uinput.KeyCapslock,
	"f1": uinput.KeyF1,
	"f2": uinput.KeyF2,
	"f3": uinput.KeyF3,
	"f4": uinput.KeyF4,
	"f5": uinput.KeyF5,
	"f6": uinput.KeyF6,
	"f7": uinput.KeyF7,
	"f8": uinput.KeyF8,
	"f9": uinput.KeyF9,
	"f10": uinput.KeyF10,
	"numlock": uinput.KeyNumlock,
	"scrolllock": uinput.KeyScrolllock,
	"kp7": uinput.KeyKp7,
	"kp8": uinput.KeyKp8,
	"kp9": uinput.KeyKp9,
	"kpminus": uinput.KeyKpminus,
	"kp4": uinput.KeyKp4,
	"kp5": uinput.KeyKp5,
	"kp6": uinput.KeyKp6,
	"kpplus": uinput.KeyKpplus,
	"kp1": uinput.KeyKp1,
	"kp2": uinput.KeyKp2,
	"kp3": uinput.KeyKp3,
	"kp0": uinput.KeyKp0,
	"kpdot": uinput.KeyKpdot,
	"zenkakuhankaku": uinput.KeyZenkakuhankaku,
	"102nd": uinput.Key102Nd,
	"f11": uinput.KeyF11,
	"f12": uinput.KeyF12,
	"ro": uinput.KeyRo,
	"katakana": uinput.KeyKatakana,
	"hiragana": uinput.KeyHiragana,
	"henkan": uinput.KeyHenkan,
	"katakanahiragana": uinput.KeyKatakanahiragana,
	"muhenkan": uinput.KeyMuhenkan,
	"kpjpcomma": uinput.KeyKpjpcomma,
	"kpenter": uinput.KeyKpenter,
	"rightctrl": uinput.KeyRightctrl,
	"kpslash": uinput.KeyKpslash,
	"sysrq": uinput.KeySysrq,
	"rightalt": uinput.KeyRightalt,
	"linefeed": uinput.KeyLinefeed,
	"home": uinput.KeyHome,
	"up": uinput.KeyUp,
	"pageup": uinput.KeyPageup,
	"left": uinput.KeyLeft,
	"right": uinput.KeyRight,
	"end": uinput.KeyEnd,
	"down": uinput.KeyDown,
	"pagedown": uinput.KeyPagedown,
	"insert": uinput.KeyInsert,
	"delete": uinput.KeyDelete,
	"macro": uinput.KeyMacro,
	"mute": uinput.KeyMute,
	"volumedown": uinput.KeyVolumedown,
	"volumeup": uinput.KeyVolumeup,
	"power": uinput.KeyPower,
	"kpequal": uinput.KeyKpequal,
	"kpplusminus": uinput.KeyKpplusminus,
	"pause": uinput.KeyPause,
	"scale": uinput.KeyScale,
	"kpcomma": uinput.KeyKpcomma,
	"hangeul": uinput.KeyHangeul,
	"hanja": uinput.KeyHanja,
	"yen": uinput.KeyYen,
	"leftmeta": uinput.KeyLeftmeta,
	"rightmeta": uinput.KeyRightmeta,
	"compose": uinput.KeyCompose,
	"stop": uinput.KeyStop,
	"again": uinput.KeyAgain,
	"props": uinput.KeyProps,
	"undo": uinput.KeyUndo,
	"front": uinput.KeyFront,
	"copy": uinput.KeyCopy,
	"open": uinput.KeyOpen,
	"paste": uinput.KeyPaste,
	"find": uinput.KeyFind,
	"cut": uinput.KeyCut,
	"help": uinput.KeyHelp,
	"menu": uinput.KeyMenu,
	"calc": uinput.KeyCalc,
	"setup": uinput.KeySetup,
	"sleep": uinput.KeySleep,
	"wakeup": uinput.KeyWakeup,
	"file": uinput.KeyFile,
	"sendfile": uinput.KeySendfile,
	"deletefile": uinput.KeyDeletefile,
	"xfer": uinput.KeyXfer,
	"prog1": uinput.KeyProg1,
	"prog2": uinput.KeyProg2,
	"www": uinput.KeyWww,
	"msdos": uinput.KeyMsdos,
	"coffee": uinput.KeyCoffee,
	"direction": uinput.KeyDirection,
	"cyclewindows": uinput.KeyCyclewindows,
	"mail": uinput.KeyMail,
	"bookmarks": uinput.KeyBookmarks,
	"computer": uinput.KeyComputer,
	"back": uinput.KeyBack,
	"forward": uinput.KeyForward,
	"closecd": uinput.KeyClosecd,
	"ejectcd": uinput.KeyEjectcd,
	"ejectclosecd": uinput.KeyEjectclosecd,
	"nextsong": uinput.KeyNextsong,
	"playpause": uinput.KeyPlaypause,
	"previoussong": uinput.KeyPrevioussong,
	"stopcd": uinput.KeyStopcd,
	"record": uinput.KeyRecord,
	"rewind": uinput.KeyRewind,
	"phone": uinput.KeyPhone,
	"iso": uinput.KeyIso,
	"config": uinput.KeyConfig,
	"homepage": uinput.KeyHomepage,
	"refresh": uinput.KeyRefresh,
	"exit": uinput.KeyExit,
	"move": uinput.KeyMove,
	"edit": uinput.KeyEdit,
	"scrollup": uinput.KeyScrollup,
	"scrolldown": uinput.KeyScrolldown,
	"kpleftparen": uinput.KeyKpleftparen,
	"kprightparen": uinput.KeyKprightparen,
	"new": uinput.KeyNew,
	"redo": uinput.KeyRedo,
	"f13": uinput.KeyF13,
	"f14": uinput.KeyF14,
	"f15": uinput.KeyF15,
	"f16": uinput.KeyF16,
	"f17": uinput.KeyF17,
	"f18": uinput.KeyF18,
	"f19": uinput.KeyF19,
	"f20": uinput.KeyF20,
	"f21": uinput.KeyF21,
	"f22": uinput.KeyF22,
	"f23": uinput.KeyF23,
	"f24": uinput.KeyF24,
	"playcd": uinput.KeyPlaycd,
	"pausecd": uinput.KeyPausecd,
	"prog3": uinput.KeyProg3,
	"prog4": uinput.KeyProg4,
	"dashboard": uinput.KeyDashboard,
	"suspend": uinput.KeySuspend,
	"close": uinput.KeyClose,
	"play": uinput.KeyPlay,
	"fastforward": uinput.KeyFastforward,
	"bassboost": uinput.KeyBassboost,
	"print": uinput.KeyPrint,
	"hp": uinput.KeyHp,
	"camera": uinput.KeyCamera,
	"sound": uinput.KeySound,
	"question": uinput.KeyQuestion,
	"email": uinput.KeyEmail,
	"chat": uinput.KeyChat,
	"search": uinput.KeySearch,
	"connect": uinput.KeyConnect,
	"finance": uinput.KeyFinance,
	"sport": uinput.KeySport,
	"shop": uinput.KeyShop,
	"alterase": uinput.KeyAlterase,
	"cancel": uinput.KeyCancel,
	"brightnessdown": uinput.KeyBrightnessdown,
	"brightnessup": uinput.KeyBrightnessup,
	"media": uinput.KeyMedia,
	"switchvideomode": uinput.KeySwitchvideomode,
	"kbdillumtoggle": uinput.KeyKbdillumtoggle,
	"kbdillumdown": uinput.KeyKbdillumdown,
	"kbdillumup": uinput.KeyKbdillumup,
	"send": uinput.KeySend,
	"reply": uinput.KeyReply,
	"forwardmail": uinput.KeyForwardmail,
	"save": uinput.KeySave,
	"documents": uinput.KeyDocuments,
	"battery": uinput.KeyBattery,
	"bluetooth": uinput.KeyBluetooth,
	"wlan": uinput.KeyWlan,
	"uwb": uinput.KeyUwb,
	"unknown": uinput.KeyUnknown,
	"videonext": uinput.KeyVideoNext,
	"videoprev": uinput.KeyVideoPrev,
	"brightnesscycle": uinput.KeyBrightnessCycle,
	"brightnesszero": uinput.KeyBrightnessZero,
	"displayoff": uinput.KeyDisplayOff,
	"wimax": uinput.KeyWimax,
	"rfkill": uinput.KeyRfkill,
	"micmute": uinput.KeyMicmute,
}

// generated by ./xkeys.bash
var xKeysNormal = map[string]int{
	"escape": uinput.KeyEsc,
	"1": uinput.Key1,
	"2": uinput.Key2,
	"3": uinput.Key3,
	"4": uinput.Key4,
	"5": uinput.Key5,
	"6": uinput.Key6,
	"7": uinput.Key7,
	"8": uinput.Key8,
	"9": uinput.Key9,
	"0": uinput.Key0,
	"minus": uinput.KeyMinus,
	"equal": uinput.KeyEqual,
	"backspace": uinput.KeyBackspace,
	"tab": uinput.KeyTab,
	"q": uinput.KeyQ,
	"w": uinput.KeyW,
	"e": uinput.KeyE,
	"r": uinput.KeyR,
	"t": uinput.KeyT,
	"y": uinput.KeyY,
	"u": uinput.KeyU,
	"i": uinput.KeyI,
	"o": uinput.KeyO,
	"p": uinput.KeyP,
	"bracketleft": uinput.KeyLeftbrace,
	"bracketright": uinput.KeyRightbrace,
	"return": uinput.KeyEnter,
	"control_l": uinput.KeyLeftctrl,
	"a": uinput.KeyA,
	"s": uinput.KeyS,
	"d": uinput.KeyD,
	"f": uinput.KeyF,
	"g": uinput.KeyG,
	"h": uinput.KeyH,
	"j": uinput.KeyJ,
	"k": uinput.KeyK,
	"l": uinput.KeyL,
	"semicolon": uinput.KeySemicolon,
	"apostrophe": uinput.KeyApostrophe,
	"grave": uinput.KeyGrave,
	"shift_l": uinput.KeyLeftshift,
	"backslash": uinput.KeyBackslash,
	"z": uinput.KeyZ,
	"x": uinput.KeyX,
	"c": uinput.KeyC,
	"v": uinput.KeyV,
	"b": uinput.KeyB,
	"n": uinput.KeyN,
	"m": uinput.KeyM,
	"comma": uinput.KeyComma,
	"period": uinput.KeyDot,
	"slash": uinput.KeySlash,
	"shift_r": uinput.KeyRightshift,
	"kp_multiply": uinput.KeyKpasterisk,
	"alt_l": uinput.KeyLeftalt,
	"space": uinput.KeySpace,
	"caps_lock": uinput.KeyCapslock,
	"f1": uinput.KeyF1,
	"f2": uinput.KeyF2,
	"f3": uinput.KeyF3,
	"f4": uinput.KeyF4,
	"f5": uinput.KeyF5,
	"f6": uinput.KeyF6,
	"f7": uinput.KeyF7,
	"f8": uinput.KeyF8,
	"f9": uinput.KeyF9,
	"f10": uinput.KeyF10,
	"num_lock": uinput.KeyNumlock,
	"scroll_lock": uinput.KeyScrolllock,
	"kp_home": uinput.KeyKp7,
	"kp_up": uinput.KeyKp8,
	"kp_prior": uinput.KeyKp9,
	"kp_subtract": uinput.KeyKpminus,
	"kp_left": uinput.KeyKp4,
	"kp_begin": uinput.KeyKp5,
	"kp_right": uinput.KeyKp6,
	"kp_add": uinput.KeyKpplus,
	"kp_end": uinput.KeyKp1,
	"kp_down": uinput.KeyKp2,
	"kp_next": uinput.KeyKp3,
	"kp_insert": uinput.KeyKp0,
	"kp_delete": uinput.KeyKpdot,
	"iso_level3_shift": uinput.KeyZenkakuhankaku,
	"f11": uinput.KeyF11,
	"f12": uinput.KeyF12,
	"katakana": uinput.KeyKatakana,
	"hiragana": uinput.KeyHiragana,
	"henkan_mode": uinput.KeyHenkan,
	"hiragana_katakana": uinput.KeyKatakanahiragana,
	"muhenkan": uinput.KeyMuhenkan,
	"kp_enter": uinput.KeyKpenter,
	"control_r": uinput.KeyRightctrl,
	"kp_divide": uinput.KeyKpslash,
	"alt_r": uinput.KeyRightalt,
	"linefeed": uinput.KeyLinefeed,
	"home": uinput.KeyHome,
	"up": uinput.KeyUp,
	"prior": uinput.KeyPageup,
	"left": uinput.KeyLeft,
	"right": uinput.KeyRight,
	"end": uinput.KeyEnd,
	"down": uinput.KeyDown,
	"next": uinput.KeyPagedown,
	"insert": uinput.KeyInsert,
	"delete": uinput.KeyDelete,
	"xf86audiomute": uinput.KeyMute,
	"xf86audiolowervolume": uinput.KeyVolumedown,
	"xf86audioraisevolume": uinput.KeyVolumeup,
	"xf86poweroff": uinput.KeyPower,
	"kp_equal": uinput.KeyKpequal,
	"plusminus": uinput.KeyKpplusminus,
	"pause": uinput.KeyPause,
	"xf86launcha": uinput.KeyScale,
	"kp_decimal": uinput.KeyKpcomma,
	"hangul": uinput.KeyHangeul,
	"hangul_hanja": uinput.KeyHanja,
	"super_l": uinput.KeyLeftmeta,
	"super_r": uinput.KeyRightmeta,
	"menu": uinput.KeyCompose,
	"redo": uinput.KeyAgain,
	"sunprops": uinput.KeyProps,
	"undo": uinput.KeyUndo,
	"sunfront": uinput.KeyFront,
	"xf86copy": uinput.KeyCopy,
	"xf86open": uinput.KeyOpen,
	"xf86paste": uinput.KeyPaste,
	"find": uinput.KeyFind,
	"xf86cut": uinput.KeyCut,
	"help": uinput.KeyHelp,
	"xf86menukb": uinput.KeyMenu,
	"xf86calculator": uinput.KeyCalc,
	"xf86sleep": uinput.KeySleep,
	"xf86wakeup": uinput.KeyWakeup,
	"xf86explorer": uinput.KeyFile,
	"xf86xfer": uinput.KeyXfer,
	"xf86launch1": uinput.KeyProg1,
	"xf86launch2": uinput.KeyProg2,
	"xf86www": uinput.KeyWww,
	"xf86dos": uinput.KeyMsdos,
	"xf86screensaver": uinput.KeyCoffee,
	"xf86rotatewindows": uinput.KeyDirection,
	"xf86taskpane": uinput.KeyCyclewindows,
	"xf86mail": uinput.KeyMail,
	"xf86favorites": uinput.KeyBookmarks,
	"xf86mycomputer": uinput.KeyComputer,
	"xf86back": uinput.KeyBack,
	"xf86forward": uinput.KeyForward,
	"xf86eject": uinput.KeyEjectcd,
	"xf86audionext": uinput.KeyNextsong,
	"xf86audioplay": uinput.KeyPlaypause,
	"xf86audioprev": uinput.KeyPrevioussong,
	"xf86audiostop": uinput.KeyStopcd,
	"xf86audiorecord": uinput.KeyRecord,
	"xf86audiorewind": uinput.KeyRewind,
	"xf86phone": uinput.KeyPhone,
	"xf86messenger": uinput.KeyChat,
	"xf86search": uinput.KeySearch,
	"xf86go": uinput.KeyConnect,
	"xf86finance": uinput.KeyFinance,
	"xf86game": uinput.KeySport,
	"xf86shop": uinput.KeyShop,
	"cancel": uinput.KeyCancel,
	"xf86monbrightnessdown": uinput.KeyBrightnessdown,
	"xf86monbrightnessup": uinput.KeyBrightnessup,
	"xf86audiomedia": uinput.KeyMedia,
	"xf86display": uinput.KeySwitchvideomode,
	"xf86kbdlightonoff": uinput.KeyKbdillumtoggle,
	"xf86kbdbrightnessdown": uinput.KeyKbdillumdown,
	"xf86kbdbrightnessup": uinput.KeyKbdillumup,
	"xf86send": uinput.KeySend,
	"xf86reply": uinput.KeyReply,
	"xf86mailforward": uinput.KeyForwardmail,
	"xf86save": uinput.KeySave,
	"xf86documents": uinput.KeyDocuments,
	"xf86battery": uinput.KeyBattery,
	"xf86bluetooth": uinput.KeyBluetooth,
	"xf86wlan": uinput.KeyWlan,
	"xf86uwb": uinput.KeyUwb,
	"xf86next_vmode": uinput.KeyVideoNext,
	"xf86prev_vmode": uinput.KeyVideoPrev,
	"xf86monbrightnesscycle": uinput.KeyBrightnessCycle,
	"xf86brightnessauto": uinput.KeyBrightnessZero,
	"xf86displayoff": uinput.KeyDisplayOff,
	"xf86wwan": uinput.KeyWimax,
	"xf86rfkill": uinput.KeyRfkill,
	"xf86webcam": uinput.KeyCamera,
	"print": uinput.KeyPrint,
}

var xKeysShifted = map[string]int{
	"exclam": uinput.Key1,
	"at": uinput.Key2,
	"numbersign": uinput.Key3,
	"dollar": uinput.Key4,
	"percent": uinput.Key5,
	"asciicircum": uinput.Key6,
	"ampersand": uinput.Key7,
	"asterisk": uinput.Key8,
	"parenleft": uinput.Key9,
	"parenright": uinput.Key0,
	"underscore": uinput.KeyMinus,
	"plus": uinput.KeyEqual,
	"iso_left_tab": uinput.KeyTab,
	"q": uinput.KeyQ,
	"w": uinput.KeyW,
	"e": uinput.KeyE,
	"r": uinput.KeyR,
	"t": uinput.KeyT,
	"y": uinput.KeyY,
	"u": uinput.KeyU,
	"i": uinput.KeyI,
	"o": uinput.KeyO,
	"p": uinput.KeyP,
	"braceleft": uinput.KeyLeftbrace,
	"braceright": uinput.KeyRightbrace,
	"a": uinput.KeyA,
	"s": uinput.KeyS,
	"d": uinput.KeyD,
	"f": uinput.KeyF,
	"g": uinput.KeyG,
	"h": uinput.KeyH,
	"j": uinput.KeyJ,
	"k": uinput.KeyK,
	"l": uinput.KeyL,
	"colon": uinput.KeySemicolon,
	"quotedbl": uinput.KeyApostrophe,
	"asciitilde": uinput.KeyGrave,
	"bar": uinput.KeyBackslash,
	"z": uinput.KeyZ,
	"x": uinput.KeyX,
	"c": uinput.KeyC,
	"v": uinput.KeyV,
	"b": uinput.KeyB,
	"n": uinput.KeyN,
	"m": uinput.KeyM,
	"less": uinput.KeyComma,
	"greater": uinput.KeyDot,
	"question": uinput.KeySlash,
	"kp_7": uinput.KeyKp7,
	"kp_8": uinput.KeyKp8,
	"kp_9": uinput.KeyKp9,
	"kp_4": uinput.KeyKp4,
	"kp_5": uinput.KeyKp5,
	"kp_6": uinput.KeyKp6,
	"kp_1": uinput.KeyKp1,
	"kp_2": uinput.KeyKp2,
	"kp_3": uinput.KeyKp3,
	"kp_0": uinput.KeyKp0,
	"sys_req": uinput.KeySysrq,
	"break": uinput.KeyPause,
	"xf86audiopause": uinput.KeyPrevioussong,
}
