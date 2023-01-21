# Changelog

Notable changes to dotool will be documented in this file.

## [1.1](https://git.sr.ht/~geb/dotool/refs/1.1)

### Added

- Added --version.
- Added keyhold and typehold.

### Changed

- The install script is now suitable for packaging and sets --version.
- There is now a default keydelay, keyhold, typedelay and typehold.
- Delays now come after keypresses not before.
- Now aborts if there are any command-line arguments.

### Fixed

- Added a udev rule so it shouldn't type guff if you're using a non-us
keyboard layout on at least X11.
- Fixed x:backspace that was simulating shift.
