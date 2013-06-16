# Ansize

Converts an image to binary ANSI art like so:

[![stage-pikachu.png](https://d23f6h5jpj26xu.cloudfront.net/jason_24594515608662_small.png)](http://img.svbtle.com/jason_24594515608662.png)

Check out the examples folder for some image samples and their corresponding output. Ex.

    cat examples/pikachu.ansi

## Installation

1. Install go
2. Install github.com/nfnt/resize
3. Build ansize

On a Mac the commands are

    brew install go
    go get github.com/nfnt/resize
    go build ansize.go

## Usage

    ./ansize <image> <output> [width]
    