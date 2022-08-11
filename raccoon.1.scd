raccoon(1) [""["raccoon-1.0.0"]]

# NAME

raccoon - Minecraft rcon client.

# SYNOPSIS

raccoon [SERVER] [COMMAND]...

# DESCRIPTION

Raccoon has two usage modes. An interactive TUI interface and a CLI interface
for firing off a single command at a time. Both modes read a list of configured
servers in /etc/raccoon/config.toml or as a per user config in your XDG config
directory (probably $HOME/.config/raccoon/config.toml). The format is very
simple:

```
[servers.local]
	addr = "localhost:25575"
	pass = "hunter123"
[servers.survival]
	addr = "survival.example.com:25575"
	pass = "alsohunter123"
[servers.creative]
	addr = "creative.example.com:25575"
	pass = "alsohunter123"
```

The first argument when running raccoon should be the name of the configured
server followed by the command to run. If no commands are given the raccoon
will enter TUI mode.

# EXAMPLE
	raccoon survival list

# BUGS

Raccoons are omnivorous and while they prefer plant-based foods, such as seeds,
nuts, and fruit they will also eat insects and other bugs.

Minecraft itself has a known bug where, when executing rcon commands, the
response is all in one string with no newline characters, even when the command
yields several lines of text when executed on the console:
https://bugs.mojang.com/browse/MC-7569

Fortunately, a friend of mine wrote a tiny server mod for fabric which fixes
this issue:
https://git.sr.ht/~hnaguski/rconn

# AUTHORS

"Maintained" by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/raccoon/
