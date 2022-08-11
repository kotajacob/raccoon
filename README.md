# `raccoon`
Minecraft rcon client. Can be used in TUI or CLI modes.

![raccoon](raccoon.png)

# build and install
```
make
sudo make install
```

# Usage
First create a config file in either `/etc/raccoon/config.toml` or in your
user's XDG config directory (probably `$HOME/.config/raccoon/config.toml`):
```
[servers.local]
	addr = "localhost:25575"
	pass = "hunter123"
```
You can list as many servers as you'd like following this simple pattern. It's
probably a good idea to make that file read-only by it's owner and perhaps an
administrator group.

Once the config has been created run `raccoon local list` where local is the
name of your configured server and list is the minecraft command to run.
If you give only a server name without a command raccoon will enter an
interactive TUI prompt.

Minecraft itself has a known bug where, when executing rcon commands, the
response is all in one string with no newline characters, even when the command
yields several lines of text when executed on the console:
https://bugs.mojang.com/browse/MC-7569

Fortunately, a friend of mine wrote a tiny server mod for fabric which fixes
this issue:
https://git.sr.ht/~hnaguski/rconn

# Author
Written and maintained by Dakota Walsh.
Up-to-date sources can be found at https://git.sr.ht/~kota/raccoon/

# License
GNU GPL version 3 or later, see LICENSE.
