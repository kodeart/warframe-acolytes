Warframe Acolytes Tracker
=========================

[![Build Status](https://travis-ci.org/kodeart/warframe-acolytes.svg?branch=master)](https://travis-ci.org/kodeart/warframe-acolytes)

Warframe Acolytes Tracker is a command-line application
for tracking the appearance of the Stalker's acolytes.

**This app DOES NOT use Warframe game files**.
You are perfectly safe.

![Warframe Acolytes Tracker](./images/screenshot.png)

Enjoy!

Usage
-----

[Download the ZIP file][1] for your operating system, unzip it anywhere you like, run the binary.

| OS               | executable   | zip file             |
|------------------|--------------|----------------------|
| Windows (64 bit) | acolytes.exe | acolytes-windows.zip |
| Linux            | acolytes     | acolytes-linux.zip   |
| macOS            | acolytes     | acolytes-darwin.zip  |

```shell script
cd path/to/warframe/acolytes/folder
```
```shell script
# Windows
acolytes.exe track -n

# Linux
$ ./acolytes track -n

# MacOS
$ ./acolytes track -n
```

### Examples

- `track` command without arguments will run the tracker **without notifications and sound**
> acolytes.exe track
---
- the tracker can send operating system notifications when the acolyte is discovered.
This option is **off** by default. **To enable notifications**, run with `-n` flag
> acolytes.exe track -n
---
- if notifications are enabled, a beep sound will play. **To disable sound** run with `-s` (silent) flag
> acolytes.exe track -ns

### World check frequency

You can use the `--refresh` or `-r` flag to set the seconds for
world-state check, if you feel that 30s default is too frequent. Ex:
```shell script
acolytes.exe track -n -r 60
``` 
will scan the world state every 60 seconds, You cannot set it below 30.

### Build from source

- checkout the source code
- run `./build.sh`
- the binaries are in the `./releases` folder

### License

[![Software license](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](LICENSE)


[1]: https://github.com/kodeart/warframe-acolytes/releases/latest
