# Snowflakey

A tiny library to implement [Snowflake IDs](https://en.wikipedia.org/wiki/Snowflake_ID) in Golang.

I'm not the best Go programmer, so contributions welcome!

## Supported implementations
### Original (Twitter) snowflake

Reference material citation: Wikipedia article (linked above)

- 1 bit: reserved for sign bit (to fit in an `int64`)
- 41 bits: timestamp (millisecond precision, after an _arbitrary_, _recent_ epoch date, since UNIX epoch time won't fit here)
- 10 bits: machine ID (server identifier, in theory this means you can have up to 2^10 = 1024 servers generating snowflakes in parallel without possibility of collisions)
- 12 bits: machine-generated sequence number (for further collision avoidance at the millisecond-level)

RFC-style diagram:

```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|S|                       Timestamp                             |
+-+                 +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                   |    Machine ID     |   Machine Seq. Num.   |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

## Planned implementations
- Instagram snowflakes
- Mastodon snowflakes