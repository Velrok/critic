# Critic

> Everybody is a critic.  
> - random quote

This is me playin around with go.

I needed a cli programm which consumed messages off our kafka topic and writes them out if they are bigger then 900kb.

This is very much not fit for any purpose expect one exect one stated above ;) .

## Why you should not use this

This has A LOT of hard codeded stuff:

  - files are written as .edn
  - files are considered write worthy if they are bigger then 900kb that number is fixed in code
  - its not configureable at all

other shortcommings:

  - its fairly slow for none obvious reasons (from from very limited understanding of go)