#!/bin/bash

git clone https://github.com/listnt/webview.git

sudo apt install libgtk-4-dev libwebkitgtk-6.0-dev
sudo apt install swig
sudo apt install ninja-build

cd webview

sudo make build

cd ../

go build -o out

chmod +x out

./out