# Go Vote Queen

A console script to vote the selected target page automatically

# Installation

Run the script

    go get github.com/zeuxisoo/go-vote-queen

    $GOPATH/bin/go-vote-queen -h
    $GOPATH/bin/go-vote-queen -k <API_KEY> -a <API_AREA>

# Development

Create the project directory

    mkdir ~/Document/go-vote-queen

Set the GOPATH

    export GOPATH=~/Document/go-vote-queen

Clone the repo to source directory

    cd ~/Document/go-vote-queen
    mkdir -p src/github.com/zeuxisoo && cd $_
    git clone https://github.com/zeuxisoo/go-vote-queen.git go-vote-queen

Install the dependent packages

    cd go-vote-queen
    go get
