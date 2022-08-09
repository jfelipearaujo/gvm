#!/bin/bash

VER="1_0_4"

echo "Building version: $VER"

# Windows
./win_x64.sh "$VER"
./win_x86.sh "$VER"

# Linux
#./linux_x64.sh "$VER"
#./linux_x86.sh "$VER"

# OSX
#./darwin_x64.sh "$VER"
#./darwin_x86.sh "$VER"

echo "Done"