A CLI-tool to manage directories with labels
============================================

How to use
----------

After setup, you can just `s mylabel` to save the current directory using
mylabel. When you navigate on a different directory and want to return, type
`l mylabel`. There is additional functionality for namespaces, list & delete.
The directories are stored as symlinks inside '~/.local/share/bkt'.

Setup
-----

1) Install from source via `go build` or use the precompiled binary in bin
(compiled for amd64).
2) include `alias.sh` into your local shell - (tested with zsh, but other shells
should work too)
