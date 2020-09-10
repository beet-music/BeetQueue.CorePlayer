# Beetroot CorePlayer

**Beetroot CorePlayer** is the reference implementation of the Beetroot client. It's meant to represent the minimum implementation of such the client.

This is largely meant for testing and development use and I won't plan on implementing the following features:

* Manual queuing (the autoqueue options exists for the server)
* Graphical user interface
* Built-in audio player

## Using It

At the time of writing `core-player --help` has the following output:

```
Usage:
  core-player.exe [OPTIONS]

Application Options:
  /v, /verbose   Show verbose information
  /p, /player:   The program to use for playback (default: ffplay)
  /l, /library:  The location of the music library (default: .)
  /s, /server:   The location of the BeetrootCloud gateway (default:
                 cloud.beetroot.app)
  /e, /ext:

Help Options:
  /?             Show this help message
  /h, /help      Show this help message
```
