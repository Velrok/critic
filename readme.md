# Critic

> Everybody is a critic.  
> - random quote

This is me playing around with go.

I needed a cli programm which consumed messages off our kafka topic and writes them out if they are bigger then 900kb.

## Install

`go install github.com/velrok/critic`

## Usage

```bash
usage: critic [<flags>]

Flags:
  --help               Show help (also see --help-long and --help-man).
  -f, --suffix="json"  The suffix to be used when writing the files.
  -s, --size=900       If the message received is bigger then <size> bytes write the message.
  -o, --output-dir="/tmp"
                       Messages are written into this directory.
  --message-separator="\n"
                       A character that is the separator between messages.
```

## Why you should not use this

other shortcommings:

  - its fairly slow for none obvious reasons (from from very limited understanding of go)
  - I have used this script only once.
  
## Why you might want to use this

Its hopefully a fairly straint forward and simple command line tool to write out messages from a stream of messages to disk in order to debug them.