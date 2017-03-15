#!/bin/bash
#
# Bash profile script to set the go development environment
# This file should be copied to /etc/profile.d and granted with execution
# permission.
#
GOPATH="/usr/local/git/go"  # The folder where sources will be developed
GOROOT="/usr/lib/golang"    # The folder contining the go compiler
GOBIN="${GOPATH}/bin"       # The folder where go binaries will be created
GOEXE="${GOBIN}"            # Windows alias for the GOBIN variable
GORACE=""                   # More on this variable in README.md
#
# Check for prerequisites in the current filesystem
function directory_check()
{
  #
  # Checking for GOROOT; if no directory is present the environment should
  # be reconfigured or this script deactivated
  [[ ! -d ${GOROOT} ]] && {
    echo "<ERROR> No folder ${GOROOT} in local filesystem.                  "
    echo "        The go environment cannot be setup, please deactivate this"
    echo "        profile script or setup/install go in your system.        "
  }
  #
  # If GOPATH doesn't exist, it will be created
  [[ ! -d ${GOPATH} ]] && mkdir ${GOPATH}
  #
  # If GOBIN doesn't exist, it will be created
  [[ ! -d ${GOBIN} ]] && mkdir ${GOBIN}
  #
}
#
# Export the go environent variables to the current operating system profile
function environment_setup()
{
  #
  # system PATH is updated to contain the value of GOBIN folder
  PATH=${PATH}:${GOBIN}
  #
  # values are exported to the operating system environment
  export GOPATH GOROOT GOBIN GOEXE GORACE PATH
}
#
# main funcion
function main()
{
  directory_check
  environment_setup
}
#
main
#
