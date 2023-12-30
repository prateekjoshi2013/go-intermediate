- Go still expects there to be a single workspace for third-party Go tools
installed via **go install** . 
- By default, this workspace is located in ***$HOME/go***,
  
  - source code for these tools stored in ***$HOME/go/src*** 
  
  - the compiled binaries in ***$HOME/go/bin***. 

-  You can use this default or specify a different workspace by setting the **$GOPATH** environment variable.

- Whether or not you use the default location, it’s a good idea to **explicitly define GOPATH** and to put the $GOPATH/bin directory in your executable path. 
  
  - Explicitly defining GOPATH makes it clear where your Go workspace is located and adding $GOPATH/bin to your executable path makes it easier to run third-party tools installed via go install

- If you are on a Unix-like system using bash, add the following lines to your .profile. 

- If you are using zsh, add these lines to .zshrc instead

    ```sh
      export GOPATH=$HOME/go
      export PATH=$PATH:$GOPATH/bin
    ```