# TESTING-GO

##### *A repository just to start a set of tests, using a base default go installation on Linux* #####
> <p><b>*Step-by-step tutorial, starting from the go environment setup, going trough the first line of code, diving into packages and ending up with complex examples, commented and including unit-testing, to let the reader introduce him/herself to the awesome go programming language ecosystem.*</b></p>


|Current environment|
| :-------------    | :------------- |
|**Testing OS**     | *Linux*        |
|**Distribution**   | *Fedora 25*    |
|**Work-space**     | *Go 1.7*       |
|**CPU Arch**       | *x86_64*       |

installed using


 `dnf -y install golang golang-bin golang-docs golang-misc golang-shared golang-src golang-tests`


from the standard fedora repo

Additional and useful, the package including go for gcc should be installed, since later on it will be needed from other packages and tools to correctly integrate the code when c/c++ libraries are bundled in a go project.

In Fedora

`dnf -y install gcc-go`

#### <u>NOTE</u>

Fedora 25 provides a dnf set of packages for a ready to use go work-space.
By the way, when installing the bundle, the environment variables are not fully configured.
When starting with go and the related work-space setting, it's quite hard to find a good documentation to follow a step by step process.

## STEP ONE ##
#### Install Go and setup the OS environment ####

An initial look at the packages file-system structure can give and idea of how the go provided tool-set has been structured.

**Main path:**

|`/usr/lib/golang`|->   |*(Main folder of the go development tools)*                          |
|:---------------:|:---:|:--------------------------------------------------------------------|
|-->              |`bin`|*(where the executable files `go` and `gofmt` are installed)*        |
|-->              |`pkg`|*(where the packages and other third party components are installed)*|
|-->              |`src`|*(where the source code resides)*                                    |

After installing the go package suite in Fedora, the following configuration can be verified calling the command

> `go env`

|*ENVIRONMENT VARIABLE*  |*VALUE SET*                                   |
|-----------------------:|:---------------------------------------------|
|**GOARCH**              |__amd64__                                     |
|**GOBIN**               |*empty - not set*                             |
|**GOEXE**               |*empty - not set*                             |
|**GOHOSTARCH**          |__amd64__                                     |
|**GOOS**                |__linux__                                     |
|**GOHOSTOS**            |__linux__                                     |
|**GOPATH**              |*empty - not set*                             |
|**GORACE**              |*empty - not set*                             |
|**GOROOT**              |__/usr__                                      |
|**GOTOOLDIR**           |__/usr/libexec/gcc/x86_64-redhat-linux/6.3.1__|
|**GO15VENDOREXPERIMENT**|__1__                                         |
|**CC**                  |__/usr/bin/gcc__                              |
|**GOGCCFLAGS**          |__-fPIC -m64 -pthread -fmessage-length=0__    |
|**CXX**                 |__/usr/bin/g++__                              |
|**CGO_ENABLED**         |__1__                                         |

These default values require some updates to correctly set the environment configuration; in details, the following variables will be reconfigured:

The environment variable `PATH` by default in every Unix OS settings includes the `/bin` folder, which in Fedora contains a reference to both `go` and `gofmt` executable programs. Effectively both programs are part of an `alternatives` configuration, set in case of multiple versions coexisting in the same workstation.

The folder where all go sources will be written, developed and tested will be

`/usr/local/git/go`

to easily maintain the relationship with local and remote repositories.

The following table will resume the new updated environment variables values:

|        |                   |
|--------|-------------------|
|`GOPATH`|`/usr/local/git/go`|
|`GOROOT`|`/usr/lib/golang`  |
|`GOBIN` |${GOPATH}/bin      |
|`GOEXE` |${GOBIN}           |
|`GORACE`|                   |

The `go_profile.sh` script in the folder `scripts` contains the lines necessary to update the environment settings and should be added to `/etc/profile.d` and executed issuing `source /etc/profile.d/go_profile.sh` during the first session, while it will be automatically loaded on every next boot.

#### Note on the GORACE environment variable ####
Full documentation at this link:
+ https://golang.org/doc/articles/race_detector.html

An extract from the previous article is a brief summary of the GORACE configuration options:
The GORACE environment variable sets race detector options. The format is:

> `GORACE="option1=val1 option2=val2"`

The options are:

|option           |default value|description|
|:----------------|:-----------:|:------------------------------------------------------------------|
|log_path         |stderr       |The race detector writes its report to a file named log_path.pid. The  special names stdout and stderr cause reports to be written to standard output and standard error, respectively.|
|exitcode         |66           |The exit status to use when exiting after a detected race.         |
|strip_path_prefix|"" *empty*   |Strip this prefix from all reported file paths, to make reports more concise.|
|history_size     |1            |The per-goroutine memory access history is 32K * 2**history_size elements. Increasing this value can avoid a "failed to restore the stack" error in reports, at the cost of increased memory usage.|
|halt_on_error    |0            |Controls whether the program exits after reporting first data race.|

> *Example:*
> > **`GORACE="log_path=/tmp/race/report strip_path_prefix=/my/go/sources/" go test -race`**


## STEP TWO ##
#### Install additional packages ####

Fedora provides additional packages that are useful to complete our first setup of go.
They can be installed using dnf or also downloaded using the go get command

* **goimports** is useful when saving code to fill in import paths
> using dnf:
> * `dnf install golang-googlecode-tools-goimports`
>
> using go get:
> * `go get golang.org/x/tools/cmd/goimports`

* **gorename** provides identifier rename support in editors
> using dnf:
> * `dnf install golang-googlecode-tools-gorename`
>
> using go get:
> * `go get golang.org/x/tools/cmd/gorename`

* **oracle** helps with code navigation and search
> using dnf:
> * `dnf install golang-googlecode-tools-oracle`
>
> using go get:
> * `go get golang.org/x/tools/cmd/oracle`

* **golint** should be run after every build to check your code
> using dnf:
> * `dnf install golint`
>
> using go get:
> * `go get github.com/golang/lint/golint`

* **gocode** is used by many editors to provide intellisense
> using dnf:
> * *<b>currently not available</b>*
>
> using go get:
> * `go get github.com/nsf/gocode`
