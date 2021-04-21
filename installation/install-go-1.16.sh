#!/bin/bash

curl -O https://storage.googleapis.com/golang/go1.16.linux-amd64.tar.gz
tar -xvf go1.16.linux-amd64.tar.gz
sudo mv go /usr/local
rm go1.16.linux-amd64.tar.gz
echo -e "export GOPATH=$HOME/work\nexport PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.bashrc