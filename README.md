# TESTING-GO

##### *A repository just to start a set of tests, using a base default go installation on Linux* #####
> <p><b>*Step-by-step tutorial, starting from the go environment setup, going trough the first line of code, diving into packages and ending up with complex examples, commented and including unit-testing, to let the reader introduce him/herself to the awesome go programming language ecosystem.*</b></p>


|Current environment|                |
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

## STEP THREE ##
#### Choose your IDE/Editor and install it ####

There are several options to develop go code in a fast and productive way; depending on the programmer attitude and habits, a go program can be written in your favorite editor or also using an IDE. The following list includes some of the most common options.

* **Sublime**
  + Link to the main site:
    - http://www.sublimetext.com/
  + A plugin from DisposaBoy, available on github with instructions to install it and examples
      - https://github.com/DisposaBoy/GoSublime
  + An interesting post from Mark Wolfe on Sublime text editor and go integration
    - http://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development/


* **vi / vim**
  + Link to the main historical site for this editor is:
    - http://www.vim.org/download.php
  + Installing vim and integrating with go, by Victor Farazdagi
    - http://farazdagi.com/blog/2015/vim-as-golang-ide/


* **Atom**
  + Link to Atom main site:
    - https://atom.io/
  + A plugin for Atom integration with the go language by Joe Fitzgerald:
    - https://github.com/joefitzgerald/go-plus


* **LiteIDE**
  + Link to sourceforge LiteIDE site:
    - http://sourceforge.net/projects/liteide/files/
  + Link to LiteIDE github:
    - http://
  + In APPENDIX B of this README document, there's a fast and brief HOW-TO to install LiteIDE in Fedora, having it up and running with a useful desktop icon and rapid access.


* **Gogland preview by JetBrains**
  + Link to JetBrains Gogland status and project:
    - https://www.jetbrains.com/go/

      *NOTE*: This IDE is still in preview, but at the current state it looks like a very good choice, especially for a developer already familiar with the JetBrains products family


* **Emacs / Xemacs**
  + Link to the historical emacs site on GNU org:
    - http://www.gnu.org/software/emacs
  + Link to Xemacs site:
    - http://www.xemacs.org/
  + The configuration for emacs used by Guillaume Charmes, on github:
    - https://github.com/creack/dotfiles


<b>*Last considerations*</b>: Also Eclipse, NetBeans and JetBrains Idea have some additional plugins to integrate with the go development, although none of them has the full functionality to use debuggers and profiling tools. Feel free to test those integrations, especially if your preference is to develop using one of those IDE tools, but better to start integrating with more advanced tools, like the ones in the previous list.

## CHAPTER ONE ##
#### Writing our first go programs ####

**Introduction**
Before start coding, we have to let our impatience wait a little bit more, because the go folder structures when developing is a point that has to be considered to avoid issues later on, especially when writing libraries, packages and additional modular components that will be required in our projects. Also the way of importing and using third-party libraries and modules requires the knowledge of this folder structure.

Let's start underlining the git repository that will be hosting our code. It's not mandatory to let our code be public, but generally the next guidelines are self-explanatory. Spend some minutes to learn the next points, and all of these rules will be a fundamental basis for your go development in the next future.

For our case-study we will use a github account as example:
* let's suppose you have an account named "mygithubuser" (but better if you change it with something existing on the git server domain of your choice)
*
  + our reference folder will be github.com/mygithubuser
  + *first step is to create the base folder*:
  + `mkdir -p $GOPATH/src/github.com/mygithubuser`

The previously created folder will be our base path for all of our go works, we can consider it like an initialization of our universal go local source folder.

Now we'll going to write our **first_program.go**

The source code is available in the folder first_program of this repo.
To recall our example folder, we will create our local program basepath issuing:

`mkdir -p $GOPATH/src/github.com/mygithubuser/first_program`



## APPENDIX A ##
#### Useful links ####

The following list has been written to give the programmer a fast reference for the basic go official resources available on the internet.

* http://golang.org/
  + The go programming language official site; there're important resources there such as the current state of the compiler, many tutorials from basic to advanced level and other useful informations.


* https://github.com/golang/go/wiki
  + The github go wiki. Any additional comment is not required, just check it !


* https://plus.google.com/+golang/posts
  + A Google+ profile dedicated to go.


* http://blog.golang.org/
  + A blog talking about go, full of interesting stats and considerations, with also references to the basics regarding the installation process and the first steps with go development.


* http://www.youtube.com/user/gocoding
  + for those who prefere to start using go watching some videos, there they can find 'em out !


* http://dave.cheney.net/
  + Dave Cheney is a go guru, talking in many worldwide events and conferences; his site includes many considerations on good practices programming in go. Could be a good reading to improve the hidden practice when programming in go.


* http://gophervids.appspot.com/
  + A collection of videos, conferences and webinars on go, a nice way to spend some time watching at the current state of integration with many of the most used platform, such as AWS, Google cloud and more.

Some videos a new-to-go programmer should watch at, they're both related to go concurrency patterns:


* http://www.youtube.com/watch?v=QDDwwePbDtw
  + Advanced Go concurrency patterns


* http://www.youtube.com/watch?v=f6kdp27TYZs
  + Go concurrency patterns


  ## APPENDIX B ##
  #### LiteIDE on Fedora installation guidelines ####

  To install LiteIDE on Fedora, the fastest and easiest way is to checkout to binaries, set the desktop link and start using the integrated development environment. Compiling LiteIDE from source is not a goal of the current docuemntation, although it's well documented reading the instructions on the author's site and github repository.

  Download the latest LiteIDE release from:
    - https://github.com/visualfc/liteide/releases/tag/x31

  From your download location, extract the downloaded file to a folder of your choice, should be in /opt main path. A working example for the current x31 release of LiteIDE can be found in the scripts folder of this repo; just run **liteide_fedora_setup.sh** and hit the desktop icon to test your final installation.
