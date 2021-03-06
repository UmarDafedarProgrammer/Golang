# Online execution

https://play.golang.org/
https://www.onlinegdb.com/online_go_compiler
https://onecompiler.com/go
https://goplay.space/

# IDEs

- vim: vim-go plugin provides Go programming language support
- Visual Studio Code: Go extension provides support for the Go programming language
- GoLand: GoLand is distributed either as a standalone IDE or as a plugin for IntelliJ IDEA Ultimate
- Atom: Go-Plus is an Atom package that provides enhanced Go support
- LiteIDE

## Ubuntu/Debian/Kali Linux:

    1) To Install Go
        wget https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz
        sudo tar -C /usr/local -xvf go1.12.4.linux-amd64.tar.gz

        Open the profile
            vim ~/.profile

        Add below Go Paths in that profile file
            export GOPATH=$HOME/go
            export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

        After saving the changes, run below command to load these changes:
            source ~/.profile

        To check go version
            go version

        Finally, clean the downloaded binary
            rm go1.12.4.linux-amd64.tar.gz


    2) To update Go
        go version
        # go version go1.8.1 linux/amd64

        sudo rm -r /usr/local/go/

        wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
        # ... Saving to: ‘go1.11.linux-amd64.tar.gz’

        sudo tar -C /usr/local -xzf go1.11.linux-amd64.tar.gz

        go version
        # go version go1.11 linux/amd64

        rm go1.11.linux-amd64.tar.gz

## Windows/MAC Installation Procedure

    - Go to Download Link (https://golang.org/dl) and download the corresponding executable

## Environment variables Setup

    - With latest versios of Go, GOPATH setup is not needed.
    - Developers do not have to care about GOPATH anymore.
    - It has a default value (depending on OS)
    - GOPATH is still used by the Go toolchain internally, for caching downloaded modules and compilation artifacts.
