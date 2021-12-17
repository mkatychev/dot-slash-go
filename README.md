# dot-slash-go

An extensible, friendly framework for project [go scripts](https://www.thoughtworks.com/insights/blog/praise-go-script-part-i).

![dot-slash-go Example](https://user-images.githubusercontent.com/11223234/104383755-3cadad80-5528-11eb-9c81-529dc70c36c7.png)

## Features

* Simple installation in existing projects
* Easily extend with new commands
* Autogenerated usage instructions
* Git module vendoring

## Installation

In your project root:

```sh
curl -fsSL "https://raw.githubusercontent.com/mkatychev/dot-slash-go/HEAD/.go/core/vendor" | bash -s -- init
```

## Adding Commands
dot-slash-go script commands are just a stock-standard script with a filename that matches the command name.
These scripts are contained within your `.go` folder, or within nested folders there if you want
to create a tree-based command structure.

For example, the script `.go/test/hello` would be available through `./go test hello`. Any arguments
passed after the command will be curried through to the script, making it trivial to pass values and
options around as needed.

### Contextual Help
The dot-slash-go script provides tools which enable your users to easily discover how to use your 
command line without needing to read your docs (a travesty, we know). To make this possible, 
you'll want to add two extra files for each command.

The first, `[command].usage` should define the arguments list that your command expects to receive,
something like `NAME [MIDDLE_NAMES...] SURNAME`. This file is entirely optional, leaving it out will
have go present the command as if it didn't accept arguments.

The second, `[command].help` is used to describe the arguments that your command accepts, as well as
provide a bit of additional context around how it works, when you should use it etc.

In addition to providing help for commands, you may also provide it for directories to explain what
their sub-commands are intended to achieve. To do this, simply add a `.help` file to the directory.


## `./go core`

`./go core vendor` and `./go core update` uses `git subtree` to vendor git repositories including dot-slash-go itself, the vendored git directories are defined in `.go/.git-vendor`:

```
Core dot-slash-go functionality commands

Usage:
  ./go core update
  ./go core vendor (add|list|remove|update) [options]

Commands:
  update     runs `./go core vendor update dot-slash-go` from root directory
  vendor     adds, lists, removes, or updates a vendored git repository
```

`./go core update` will pull changes from the upstream dot-slash-go repo defined in `.go/.git-vendor`

## Credits
This is based on these projects:
* [Bash CLI](https://github.com/SierraSoftworks/bash-cli)
* [git-vendor](https://github.com/brettlangdon/git-vendor)

Some changes include
* Consolidation of main logic to a single file
* Removal of install/uninstall commands
* Storage of commands and metadata in .go rather than app

## Frequently Asked Questions

1. **Can I use dot-slash-go to run things which aren't bash scripts?**
   Absolutely, dot-slash-go simply executes files - it doesn't care whether they're written in Bash, Ruby,
   Python or Go - if you can execute the file then you can use it with dot-slash-go.

1. **How can I define common variables, or utility functions?**
   Place a dot-prefixed executable shell script in your `.go` folder such as `.globals`. Then, execute 
   that script from your commands. For example:
   
   ```bash
   #!/usr/bin/env bash
   
   set -e
   . ".go/.globals"
   
   # Call a function defined in globals
   init
   # Do command-specific work
   ```
