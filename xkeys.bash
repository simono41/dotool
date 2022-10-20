#!/bin/bash

align() {
	sed '/\<KeyRo\>/ d
	/\<KeyKpjpcomma\>/ d
	/\<KeyMacro\>/ d
	/\<KeyYen\>/ d
	/\<KeySetup\>/ d
	/\<KeyDeletefile\>/ d
	/\<KeyClosecd\>/ d
	/\<KeyEjectclosecd\>/ d
	/\<KeyIso\>/ d
	/\<KeyMove\>/ d
	/\<KeyEdit\>/ d
	/\<KeyAlterase\>/ d
	/\<KeyUnknown\>/ d
	/\<KeyMicmute\>/ d'
}

echo 'var xKeysNormal = map[string]int{'
{
	paste -d ' ' <(xmodmap -pke | sed '1 d; s/.*= /"/; /.*=/ d; s/ .*/":/' | sed '/^"XF86Eject"/ { N; s/.*\n// }') \
		<(go doc uinput.keyesc | sed '/Key/ !d; s/^\s*/uinput./; s/ .*/,/' | align) |
		# Skip really non-matching section, we echo some of them below
		sed '/^"XF86Tools"/,/^"XF86AudioPreset"/ d' |
		# Remove duplicate keys
		sed '/^"XF86Mail":.*Email/ d; /^"Cancel":.*Stop/ d; /^"XF86Send":.*file/ d; /^"Print":.*Sysrq/ d'

	echo '"XF86WebCam": uinput.KeyCamera,'
	echo '"Print": uinput.KeyPrint,'
} | sed 's/^".*"/\L&/; s/^/\t/'
echo '}'

echo ''
echo 'var xKeysShifted = map[string]int{'
{
	paste -d ' ' <(xmodmap -pke | sed '1 d; s/.*= /"/; /.*=/ d; s/\S* /"/; s/ .*/":/' | sed '/^"XF86Eject"/ { N; s/.*\n// }') \
		<(go doc uinput.keyesc | sed '/Key/ !d; s/^\s*/uinput./; s/ .*/,/' | align) | sed '/^"NoSymbol"/ d; /^\S*_[LR]"/ d' |
		# Remove duplicate keys
		sed '/^"KP_Decimal":.*Kpcomma/ d; /\<Key102Nd\>/ d'
} | sed 's/^".*"/\L&/; s/^/\t/'
echo '}'
