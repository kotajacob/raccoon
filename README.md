# `raccoon`
Minecraft rcon client. Can be used in TUI or CLI modes.

![raccoon](raccoon.png)

# build and install
```
make
sudo make install
```

# usage
First create a config file in either `/etc/raccoon/config.toml` or in your
user's XDG config directory (probably `$HOME/.config/raccoon/config.toml`):
```
addr = "localhost:25575"
pass = "hunter123"
```

It's probably a good idea to make that file read-only by it's owner and perhaps
an administrator group. Once the config has been created run `raccoon list` or
any other command or no commands at all to enter an interactive TUI prompt.

# Author
Written and maintained by Dakota Walsh.
Up-to-date sources can be found at https://git.sr.ht/~kota/raccoon/

# License
GNU GPL version 3 or later, see LICENSE.
