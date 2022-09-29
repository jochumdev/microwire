# microwire

**WIP - This is work in Progress - DO NOT USE**

microwire intends to extends go-micro with wire generated code, its intended to merge microwire into go-micro.dev/v5.

## Goals

- Backward compatiblity is not required at all places but a nice to have
- Replace [util/cmd](https://github.com/go-micro/go-micro/tree/master/util/cmd)
- Choose what features you want from go-micro:
  - You can select if you want Auth/Events/"name it here"
  - If you don't want CLI opts you can exclude them
  - No default plugins, that makes the core even slimer
- Choose your Prefix for CLI opts not only "micro"
- No more Globals, everything lives in the DI
- wire for the end users will allow them wire theier app together
- Easy to make a micro Service and a monolith with the same codebase

## Known bugs

- When you execute an App with "--help" or "--version" it does not exit after that.

## Authors

- René Jochum - rene@jochum.dev
- Davincible - Lot's of help

## License

Go Micro is Apache 2.0 licensed.