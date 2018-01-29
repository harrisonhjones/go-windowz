# windowz

A command line utility (and library) to set the z order of the foreground window after a delay, for Windows, written in Go/Golang.

## Install

1. `go get harrisonhjones.com/go-windowz`
1. Add windowz.exe to your PATH

## Usage

`windowz --help`

## Examples

Set the foreground window to topmost in 5 seconds:

`windowz`

To "undo" set it to `nottopmost` (and optionally wait just 3 seconds):

`windowz -d 3s -z nottopmost`