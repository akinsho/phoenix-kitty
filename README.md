# Kitty phoenix
This is a golang take on the [kitty-save-session](https://github.com/dflock/kitty-save-session)
python script.

If you've been using that for a while there's probably no need to use this instead as it does
the same thing.

## Requirements

- `go`

## Usage
For now you need to build the binary yourself using `go build`, this should output a `phoenix-kitty` binary
you can then run that using `./phoenix-kitty` which should output a kitty session file based on your current
session called `session.conf` by default.

You can pass the program some arguments to change its behaviour
* `-vim` will output a vim modeline for the session file
* `-source` will read the kitty state from a json file e.g. one previously output from kitty
* `-filename` will allow you to change the name of the file that is output
