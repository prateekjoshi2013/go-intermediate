https://go101.org/apps-and-libs/gotv.html

### Installation

- Run to install GoTV.

```sh
    go install go101.org/gotv@latest
```


- **A 1.17+ toolchain version is needed to finish the installation.**

### Usage

- Run gotv without any arguments to show help messages.

#### sync the local Go git repository with remote

```sh
    gotv fetch-versions
```

#### list all versions seen locally

```sh
    gotv list-versions
```

#### build and cache some toolchain versions

```sh
    gotv cache-version ToolchainVersion [ToolchainVersion ...]
```

#### uncache some toolchain versions to save disk space

```sh
    gotv uncache-version ToolchainVersion [ToolchainVersion ...]
```

#### pin a specified toolchain version at a stable path

```sh
    gotv pin-version ToolchainVersion
```

#### unpin the current pinned toolchain version

```sh
    gotv unpin-version
```

#### set the default toolchain version (since v0.2.1)

```sh
    gotv default-version ToolchainVersion
```

#### check the default toolchain version (since v0.2.1)

```sh
    gotv default-version
```

### Pin a toolchain version

- We can use the gotv pin-version command to pin a specific toolchain version to a stable path. 

- After adding the stable path to the PATH environment veriable, we can use the official go command directly. 
  
  - And after doing these, the toolchain versions installed through ways other than GoTV may be safely uninstalled.

- It is recommanded to pin a 1.17+ version for bootstrap purpose now. The following example shows how to pin Go toolchain version 1.17.13:

```sh
    $ gotv pin-version 1.17.
    [Run]: cp -r $HOME/.cache/gotv/tag_go1.17.13 $HOME/.cache/gotv/pinned-toolchain

    Please put the following shown pinned toolchain path in
    your PATH environment variable to use go commands directly:

        /home/username/.cache/gotv/pinned-toolchain/bin
```

- After the prompted path is added to the PATH environment variable, open a new terminal window:

```sh
    $ go version
    go version go1.17.13 linux/amd64
```

The command ```gotv pin-version .!``` will upgrade the pinned toolchain to the latest release version (which may be a beta or rc version).

