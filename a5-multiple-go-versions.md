- **GOROOT** is an environment variable in the Go programming language that specifies the root directory of the Go installation. This variable is used by the Go runtime and tools to locate the standard library, compiler, and other essential components of the Go distribution.

One option is to install a secondary Go environment. For example, if you are currently running version 1.15.2 and wanted to try out version 1.15.6, you would use the following commands:

```sh
    go get golang.org/dl/go.1.15.6
    go1.15.6 download
```

- You can then use the command go1.15.6 instead of the go command to see if version 1.15.6 works for your programs:

```sh
    go1.15.6 build
```  

- Once you have validated that your code works, you can delete the secondary environment

  - first find its GOROOT, 

```sh
    go1.15.6 env GOROOT

    /Users/gobook/sdk/go1.15.6
```

  - Deleting it, and then deleting its binary from your $GOPATH/bin directory. Here’s how to do that on Mac OS, Linux, and BSD:

```sh
    rm -rf $(go1.15.6 env GOROOT)
    rm $(go env GOPATH)/bin/go1.15.6
```

### Updating go and go tools

- Mac and Windows users have the easiest path. Those who installed with brew or
chocolatey can use those tools to update. 
  
  - Those who used the installers on https://golang.org/dl can download the latest installer,which removes the old version when it installs the new one.

- Linux and BSD users need to download the latest version, move the old version to a backup directory, expand the new version, and then delete the old version:

```sh
    mv /usr/local/go /usr/local/old-go
    tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
    rm -rf /usr/local/old-go
```