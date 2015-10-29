# Critic

> Everybody is a critic.  
> - random quote

This is me playing around with go.

I needed a cli programm which consumed messages off our kafka topic and writes them out if they are bigger then 900kb.

## Install

`go install github.com/velrok/critic`

## Usage

`critic --help`

## Why you should not use this

other shortcommings:

  - its fairly slow for none obvious reasons (from from very limited understanding of go)
  - I have used this script only once.
  
## Why you might want to use this

Its hopefully a fairly straint forward and simple command line tool to write out messages from a stream of messages to disk in order to debug them.