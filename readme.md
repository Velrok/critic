# Critic

> Everybody is a critic.  
> - random quote

This is me playing around with go.

I needed a cli programm which consumes messages from a kafka topic and writes them out if they are bigger then 900kb.

## Install

`go install github.com/velrok/critic`

## Usage

```
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

## Why you should *not* use this

  - its fairly slow
  - I have used this script only once.
  
## Why you might want to use this

It's hopefully a fairly straint forward command line tool to write out big kafka messages for debugging.
