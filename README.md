# Ansize

Converts an image to binary ANSI art like so:

[![stage-pikachu.png](https://d23f6h5jpj26xu.cloudfront.net/jason_24594515608662_small.png)](http://img.svbtle.com/jason_24594515608662.png)

Check out the examples folder for some image samples and their corresponding output. Ex.

    cat examples/pikachu.ansi

I optimized for images with dark backgrounds and used 0's and 1's for the character set but it's pretty easy to customize the code to convert to your liking! The basic strategy of conversion is very simple:

1. Shrink image to desired size
2. For each pixel, find the corresponding color in ANSI's limited color palette
3. Set the foreground to that color
4. Print a random 0 or 1

## Installation

1. Install go
2. Set your $GOPATH
3. Install github.com/nfnt/
4. Build ansize

On a Mac with Homebrew the commands are

    brew install go
    mkdir /usr/local/lib/go
    export GOPATH=/usr/local/lib/go
    go get github.com/nfnt/resize
    go build ansize.go

On ArchLinux, you can simply install the package [`ansize-git`](https://aur.archlinux.org/packages/ansize-git/).

## Usage

    ./ansize <image> <output> [width]
