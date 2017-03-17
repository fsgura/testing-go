<style>.hlt {background-color:blue;color:yellow; }</style>
<h1><span class="hlt">TESTING-GO</span></h1>

##### *A repository just to start a set of tests, using a base default go installation on Linux*

*Step-by-step tutorial, starting from the go environment setup, going trough the first line of code, diving into packages and ending up with complex examples, commented and including unit-testing, to let the reader introduce him/herself to the awesome go programming language ecosystem.*

### ==TABLE OF CONTENTS==

- #### [TESTING-GO](#testing-go)
	- **[Introduction](#introduction)**
    	- [Current Ecosystem Setup](#current-ecosystem-setup)
		- [NOTE](#note)
		- [STEP ONE](#step-one)
			- *[Install Go and setup the OS environment](#install-go-and-setup-the-os-environment)*
			- *[Note on the GORACE environment variable](#note-on-the-gorace-environment-variable)*
		- [STEP TWO](#step-two)
			- *[Install additional packages](#install-additional-packages)*
		- [STEP THREE](#step-three)
			- *[Choose your IDE/Editor and install it](#choose-your-ide-editor-and-install-it)*
	- **[CHAPTER ONE - first steps in go](#chapter-one)**
		- [Writing our first go program](#writing-our-first-go-program)
			- *[Introduction](#Introduction)*
			- *[Setting path and code](#setting-path-and-code)*
			- *[Building and installing our first program](#building-adn-installing-our-first-program)*
		- [Writing our first go module package](#writing-our-first-go-module-package)
			- *[Our first library](#our-first-library)*
			- *[Using our library](#using-our-library)*
			- *[Unit Testing](#and-now-unit-testing-why-not)*
		- [After stage one, you know now](#you-ve-learned)
	- **[CHAPTER TWO - language basics](#chapter-two)**
		- [After stage two, you know now](#you-ve-learned)
	- **[CHAPTER THREE - advanced techniques and examples](#chapter-three)**
		- [After stage three, you know now](#you-ve-learned)
	- **[CHAPTER FOUR - our-first-restful-api-gateway](#chapter-four)**
		- [After stage four, you know now](#you-ve-learned)
	- **[CHAPTER FIVE - Adding the Swagger 2.0 UI to our RESTful API gateway](#chapter-five)**
		- [After stage five, you know now](#you-ve-learned)
	- **[CHAPTER SIX - Interacting with docker](#chapter-six)**
		- [After stage six, you know now](#you-ve-learned)
	- **[CHAPTER SEVEN - Interacting with kafka](#chapter-seven)**
		- [After stage seven, you know now](#you-ve-learned)
	- **[CHAPTER EIGHT - Interacting with zookeeper](#chapter-eight)**
		- [After stage eight, you know now](#you-ve-learned)
	- **[CHAPTER NINE - Authenticating to Auth0, from basic to advanced](#chapter-nine)**
		- [After stage nine, you know now](#you-ve-learned)
	- **[CHAPTER TEN - Interacting with Cassandra](#chapter-ten)**
		- [After stage ten, you know now](#you-ve-learned)
	- **[CHAPTER ELEVEN - Interacting with Elasticsearch](#chapter-eleven)**
		- [After stage eleven, you know now](#you-ve-learned)
    - **[APPENDIX A](#appendix-a)**
    - **[APPENDIX B](#appendix-b)**

***

### ==Introduction==

#### ==*Current ecosystem setup*==

|Current environment|                |
| :-------------    | :------------- |
|**Testing OS**     | *Linux*        |
|**Distribution**   | *Fedora 25*    |
|**Work-space**     | *Go 1.7*       |
|**CPU Arch**       | *x86_64*       |

installed issuing

```bash
dnf -y install golang golang-bin golang-docs golang-misc golang-shared golang-src golang-tests
```
from the standard fedora repo.

**Additional and useful info**: the package including go for gcc should be installed, since later on it will be needed from other packages and tools to correctly integrate the code when c/c++ libraries are bundled in a go project. In most cases the bundled go compiler will produce libraries without needing other additional products, but the performance of the compiled artifactory can be lower if gcc compiler and related environment variables optimization are not used when compiling under Unix/Linux Linux.

In Fedora

```bash
dnf -y install gcc-go
```

#### **NOTE**

Fedora 25 provides a dnf set of packages for a ready to use go work-space.
By the way, when installing the bundle, the environment variables are not fully configured.
When starting with go and the related work-space setting, it's quite hard to find a good documentation to follow a step by step process.

## ==STEP ONE==

#### Install Go and setup the OS environment

An initial look at the packages file-system structure can give and idea of how the go provided tool-set has been structured.

**Main path:**

|`/usr/lib/golang`|->   |*(Main folder of the go development tools)*                          |
|:---------------:|:---:|:--------------------------------------------------------------------|
|-->              |`bin`|*(where the executable files `go` and `gofmt` are installed)*        |
|-->              |`pkg`|*(where the packages and other third party components are installed)*|
|-->              |`src`|*(where the source code resides)*                                    |

After installing the go package suite in Fedora, the following configuration can be verified calling the command

```bash
go env
```

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

The folder where all go sources will be written, developed and tested will be `/usr/local/git/go` to easily maintain the relationship with local and remote repositories.

The following table resumes the new updated environment variables values:

|        |                   |
|--------|:-----------------:|
|`GOPATH`|`/usr/local/git/go`|
|`GOROOT`|`/usr/lib/golang`  |
|`GOBIN` |`${GOPATH}/bin`    |
|`GOEXE` |`${GOBIN}`         |
|`GORACE`|*empty*            |

The `go_profile.sh` script in the folder `scripts` contains the lines necessary to update the environment settings and should be added to `/etc/profile.d`, then executed issuing `source /etc/profile.d/go_profile.sh` during the first session; after the given setup, it will be automatically loaded on every next boot by the OS startup.

#### Note on the GORACE environment variable
Full documentation at this link:
+ https://golang.org/doc/articles/race_detector.html

An extract from the previous article is a brief summary of the GORACE configuration options:
The GORACE environment variable sets race detector options. The format is:

```bash
GORACE="option1=val1 option2=val2"
```

The options are:

|option           |default value|description|
|:----------------|:-----------:|:------------------------------------------------------------------|
|log_path         |stderr       |The race detector writes its report to a file named log_path.pid. The  special names stdout and stderr cause reports to be written to standard output and standard error, respectively.|
|exitcode         |66           |The exit status to use when exiting after a detected race.         |
|strip_path_prefix|"" *empty*   |Strip this prefix from all reported file paths, to make reports more concise.|
|history_size     |1            |The per-goroutine memory access history is 32K \* 2 \*\*history_size elements. Increasing this value can avoid a "failed to restore the stack" error in reports, at the cost of increased memory usage.|
|halt_on_error    |0            |Controls whether the program exits after reporting first data race.|

*Example:*
```bash
GORACE="log_path=/tmp/race/report strip_path_prefix=/my/go/sources/" go test -race
```

## ==STEP TWO==

#### Install additional packages

Fedora provides additional packages that are useful to complete our first setup of go.
They can be installed using dnf or also downloaded using the go get command:

* **goimports** is useful when saving code to fill in import paths

using dnf:
```bash
dnf install golang-googlecode-tools-goimports
```
using go get:
```bash
go get golang.org/x/tools/cmd/goimports
```

* **gorename** provides identifier rename support in editors

using dnf:
```bash
dnf install golang-googlecode-tools-gorename
```
using go get:
```bash
go get golang.org/x/tools/cmd/gorename
```

* **oracle** helps with code navigation and search

using dnf:
```bash
dnf install golang-googlecode-tools-oracle
```
using go get:
```bash
go get golang.org/x/tools/cmd/oracle
```

* **golint** should be run after every build to check your code

using dnf:
```bash
dnf install golint
```
using go get:
```bash
go get github.com/golang/lint/golint
```

* **gocode** is used by many editors to provide intellisense

using dnf:
* *<b>currently not available</b>*

using go get:
```bash
go get github.com/nsf/gocode
```

## ==STEP THREE==
#### Choose your IDE/Editor and install it

There are several options to develop go code in a fast and productive way; depending on the programmer attitude and habits, a go program can be written in your favorite editor or also using an IDE. The following list includes some of the most common options.

* **Sublime**
  + Link to the main site:
    - http://www.sublimetext.com
  + A plugin from DisposaBoy, available on github with instructions to install it and examples
    - https://github.com/DisposaBoy/GoSublime
  + An interesting post from Mark Wolfe on Sublime text editor and go integration
    - http://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development


* **vi / vim**
  + Link to the main historical site for this editor is:
    - http://www.vim.org/download.php
  + Installing vim and integrating with go, by Victor Farazdagi
    - http://farazdagi.com/blog/2015/vim-as-golang-ide


* **Atom**
  + Link to Atom main site:
    - https://atom.io
  + A plugin for Atom integration with the go language by Joe Fitzgerald:
    - https://github.com/joefitzgerald/go-plus


* **LiteIDE**
  + Link to sourceforge LiteIDE site:
    - http://sourceforge.net/projects/liteide/files
  + Link to LiteIDE github:
    - https://github.com/visualfc/liteide
<br />*In the [APPENDIX B](#appendix_b) of this README document, there's a fast and brief HOW-TO to install LiteIDE in Fedora, having it up and running with a useful desktop icon and rapid access.*


* **Gogland preview by JetBrains**
  + Link to JetBrains Gogland status and project:
    - https://www.jetbrains.com/go

      *NOTE*: This IDE is still in preview, but at the current state it looks like a very good choice, especially for a developer already familiar with the JetBrains products family


* **Emacs / Xemacs**
  + Link to the historical emacs site on GNU org:
    - http://www.gnu.org/software/emacs
  + Link to Xemacs site:
    - http://www.xemacs.org/
  + The configuration for emacs used by Guillaume Charmes, on github:
    - https://github.com/creack/dotfiles


<b>*Last considerations*</b>: Also Eclipse, NetBeans and JetBrains Idea have some additional plugins to integrate with the go development, although none of them has the full functionality to use debuggers and profiling tools. Feel free to test those integrations, especially if your preference is to develop using one of those IDE tools, but better to start integrating with more advanced tools, like the ones in the previous list.

## ==CHAPTER ONE==

### Writing our first go program

#### **Introduction**

Before start coding, we have to let our impatience wait a little bit more, because the go folder structures when developing is a point that has to be considered to avoid issues later on, especially when writing libraries, packages and additional modular components that will be required in our projects. Also the way of importing and using third-party libraries and modules requires the knowledge of this folder structure.

Let's start underlining the git repository that will be hosting our code. It's not mandatory to let our code be public, but generally the next guidelines are self-explanatory. Spend some minutes to learn the next points, and all of these rules will be a fundamental basis for your go development in the next future.


#### **Setting path and code**

For our case-study we will use a github account as example:

* let's suppose you have an account named "mygithubuser" (but better if you change it with something existing on the git server domain of your choice)

  + *our reference folder will be github.com/mygithubuser*
  + *first step is creating the base folder*:

```bash
mkdir -p ${GOPATH}/src/github.com/${mygithubuser}
```

The previously created folder will be our base path for all of our go works, we can consider it like an initialization of our universal go local source folder.

Now we'll going to write our **first_program.go**

The source code is available in the folder first_program of this repo.
To recall our example folder, we will create our local program base-path issuing:

```bash
mkdir -p ${GOPATH}/src/github.com/${mygithubuser}/first_program
cd !$
```

#### **Building and installing our first program**

After creating the folder and changing our directory to that path, let's put the go program inside it.
Now the go code can be built and installed.
To check if the program build correctly, it's enough to run

```bash
go build first_program.go
```

This will produce our go binary in the current folder if everything compiles correctly and the reason is that we specified our first_program.go as a main package. In the next example, when writing packages, we'll see the build command only check the code and build it without producing executables.
We can remove the first_program binary from our current local folder and proceed to install it.

```bash
rm -f first_program
```

The command

```bash
go install first_program.go
```

will compile our binary executable file in `${GOPATH}/bin` and since this folder should have been added to the local path following the configuration steps, just calling it from the command line will execute our program.

### **Writing our first go module package**

#### **Our first library**

It's now time to write a package containing some functions to use it in our programs.
For this scope, we'll create another folder in our main path.

```bash
mkdir -p ${GOPATH}/src/github.com/${mygithubuser}/string_utils
cd !$
```

As usual, the file containing the source code for this example can be found in the string_utils folder of this project, so let's copy/edit reverse_string.go in the previously created directory.
Let's check our package compiles without errors.

```bash
go build -a -v -x reverse_string.go
```

If all operations perform successfully, we can now integrate our library containing the function *Reverse* from a regular program.

**Using our library**

Let's use the code inside the folder ``using_string_utils``, specifically the program `revert_string_program.go` building and running it. Checking the source, you can note that we're using our previously created library and the build process successfully links all the reference in one binary, saved in the `${GOPATH}/bin` folder. In details, let's issue the following command:

```bash
go build -a -v -x revert_string_program.go
```

This should successfully build our program, and the executable will be produced in the same folder.
We can now remove the revert_string_program binary from our current local folder and proceed to install it.

```bash
rm -f revert_string_program
```

To install it, let's run

```bash
go install revert_string_program.go
```

This install process will produce a binary in our current ${GOBIN} path, but the most interesting and good surprise is that if we check our pkg library path, now the compiler installed our static library automatically !
Check it running

```bash
ls ${GOPATH}/pkg/gccgo_linux_amd64/github.com/fsgura/testing-go
```

If you're not compiling your go artifacts having gcc-go installed, check the next folder instead

```bash
ls ${GOPATH}/pkg/linux_amd64/github.com/fsgura/testing-go
```

*Note*: if you changed some folders in your current test, remember to update the previous references.

**...and now ... Unit Testing? Why not ???**

A brief example of unit testing, using the lightweight go bundled unit testing set of functions can be applied to our previous example. Let's see how !

First of all, the most common practice to write unit tests for go programs and packages is to do that inside the same folder of the sources to be compiled. Everyone can do that trying to find a different way, but the references could be something messy to set. We will follow the "go-way" for this example.

Let's check the file in the current project folder string_utils called *reverse_string_test.go*, containing the unit tests for our package. After reading the comments and analyzing the code, let's see the test results, typing the following command in our command line:

```bash
go test reverse_string_test.go
```

The output will require no additional comments if everything is configured fine.
Something importantis the reference to the color ansi package, which is not present in the default installation of the go development environent; when adding a package, dependencies should be checked and eventually satisfied by downloading the referenced libraries.

For our color library, the command to run is

```bash
go get github.com/mgutz/ansi
```

In the test file string_utils/reverse_string_test.go, a brief usage of the previosly installed color library has been used to obtain a colored output, when running the tests inside it in the terminal console window.

*NOTE*: the unit test files follow a strict notation, which is needed by the go compiler to find them.

*Every test has to be named like the program/library to test and its file-name has to end with a **_test.go**

**If you reached this point you're at the end of your first go experience chapter !!!**

**On stage one, you know now**
+ how to write your first program and compile it
+ how to write your first library and compile it
+ how to write a unit test

#### *that's a big first step !!!*
___
#### **CONGRATULATIONS and ...**

<center>**;-) GOing to see you again in the second chapter, don't miss that !!!**</center>

***

## ==CHAPTER TWO==

### Language Basics

#### **Introduction**

***

## ==CHAPTER THREE==

### Advanced techniques and examples

#### **Introduction**

***

## ==CHAPTER FOUR==

### Our first RESTful API gateway

#### **Introduction**

It's time to put our hands working on the so common and known REST. A main difference between Go
and other languages to build this kind of service is ... that if the developer strictly follows an ordered and well sorted guideline related to folder structure, packages keeping them as small as possible and comments when the project is a distributed one everything in Go is really more simple, more effective and faster, due to its binary executables.
As already mentioned when writing about the Language Basics and Advanced Techniques, the good practice using variables and pointers, slices and structs, and all the best ways to correctly define functions and resources consumption will produce a top production-performance program for each case.

**File-system setup**

Let's go working on our local disk now ! Setting up correctly the project skeleton, as already marked in the whole current document, will give us the correct modular development and deployment of our final artifact !

```bash
mkdir -p ${GOPATH}/src/github.com/${${mygithubuser}}/gorilla_based_rest_gateway
cd !$
```

As usual for this tutorial, in the bash folder of the current project you will find a script to automate the whole folder structure creation. Feel free to customize the folder names with preferred ones of your choice, but don't forget to check the code when building and executing your components.
A required dependency for our RESTful API gateway is the **gorilla package**
If you don't have it already in your go packages folder, it's time to go for it !

```bash
go get github.com/gorilla/mux
```

***

## ==CHAPTER FIVE==

### Adding the Swagger 2.0 UI to our RESTful API gateway

#### **Introduction**

***

## ==CHAPTER SIX==

### Interacting with docker

#### **Introduction**

***

## ==CHAPTER SEVEN==

### Interacting with kafka

#### **Introduction**

***

## ==CHAPTER EIGHT==

### Interacting with zookeeper

### **Introduction**

***

## ==CHAPTER NINE==

### Autheticating to Auth0, from basic to advanced

#### **Introduction**

***

## ==CHAPTER TEN==

### Interacting with Cassandra

#### **Introduction**

***

## ==CHAPTER ELEVEN==

### Interacting with Elasticsearch

#### **Introduction**

***

## ==APPENDIX A==

### Useful links

The following list has been written to give the programmer a fast reference for the basic go official resources available on the internet.

* http://golang.org/
  + The go programming language official site; there're important resources there such as the current state of the compiler, many tutorials from basic to advanced level and other useful informations.


* https://github.com/golang/go/wiki
  + The github go wiki. Any additional comment is not required, just check it !


* https://plus.google.com/+golang/posts
  + A Google+ profile dedicated to go.


* http://blog.golang.org/
  + A blog talking about go, full of interesting stats and considerations, with also references to the basics regarding the installation process and the first steps with go development.

* https://github.com/avelino/awesome-go
  + A github document site, with a long list of go resources, tutorials and all you need to check on the internet, constantly updated. Very good summary of the available third party works using go.

* http://www.youtube.com/user/gocoding
  + for those who prefere to start using go watching some videos, there they can find 'em out !


* http://dave.cheney.net/
  + Dave Cheney is a go guru, talking in many worldwide events and conferences; his site includes many considerations on good practices programming in go. Could be a good reading to improve the hidden practice when programming in go.


* http://gophervids.appspot.com/
  + A collection of videos, conferences and webinars on go, a nice way to spend some time watching at the current state of integration with many of the most used platform, such as AWS, Google cloud and more.

Some videos a new-to-go programmer should watch at, they're both related to go concurrency patterns:


* http://www.youtube.com/watch?v=QDDwwePbDtw
  + Advanced Go concurrency patterns


* [http://www.youtube.com/watch?v=f6kdp27TYZs](http://www.youtube.com/watch?v=f6kdp27TYZs)
  + Go concurrency patterns

***

## ==APPENDIX B==

### LiteIDE on Fedora installation guidelines

  To install LiteIDE on Fedora, the fastest and easiest way is to checkout to binaries, set the desktop link and start using the integrated development environment. Compiling LiteIDE from source is not a goal of the current docuemntation, although it's well documented reading the instructions on the author's site and github repository.

  Download the latest LiteIDE release from:
  * [https://github.com/visualfc/liteide/releases/tag/x31](https://github.com/visualfc/liteide/releases/tag/x31)

  From your download location, extract the downloaded file to a folder of your choice, should be in /opt main path. A working example for the current x31 release of LiteIDE can be found in the scripts folder of this repo; just run **liteide_fedora_setup.sh** and hit the desktop icon to test your final installation.

***

## ==Changelog==

***

## ==License==

***

## ==Authors and contributors==

Project initially started by:
* Fabrizio Sgura  <fsgura@psl.com.co>
