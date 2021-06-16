# Up

## The simple terminal package that allows you to check if you can connect to the internet.

# About

This is a simple package that just tests the ability to connect to [Example.com](http://example.com). If it doesn't is tells you that your internet is down, if it works it will say it works.

# Installation

## Windows

Installing is very simple, if your using windows, just download the required .exe file and run it from the command prompt. 

Another way to run this script on windows, is an installation of WSL. Follow linux steps if you will install using a WSL Terminal

## Linux

Installing on linux is as simple as it can be.

````
curl -s https://api.github.com/repos/elijahcruz12/http-up/releases/latest \
| grep "up-linux-amd64" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget -qi - 

chmod +x up-linux-amd64

sudo mv up-linux-amd64 /usr/local/bin/up
````

From here you just need to run `up` from your terminal, and you'll be able to test your internet.

# Usage

## Windows

````
up-windows-amd64.exe

Testing your internet connection
Internet is working
````

## Linux

````
user:group $ up

Testing your internet connection
Internet is working
````

This is currently the only this it does, however we will be adding the ability to just give `up` or `down` when calling the script, which will come in a later version.