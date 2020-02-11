Warframe Acolyte Tracker
========================

Warframe Acolyte Tracker is a command-line application
for tracking the appearance of the Stalker's acolytes.

Usage
-----

Choose the executable for your operation system.

| OS               | executable      |
|------------------|-----------------|
| Windows (64 bit) | acolytes.exe    |
| Linux            | acolytes-linux  |
| MacOS            | acolytes-darwin |

### Examples

- `track` command without arguments will run the tracker **without notifications and sound**
> acolytes.exe track

- the tracker can send operating system notifications when the acolyte is discovered.
This option is **off** by default. **To enable** notifications, run with `-n` flag
> acolytes.exe track -n

- if notifications are enabled, a beep sound will play. **To disable sound** run with `-s` (silent) flag
> acolytes.exe track -ns

Enjoy!
