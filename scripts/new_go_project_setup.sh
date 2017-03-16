#!/bin/bash
#
# To parse the arguments the following arrays are declared
declare input_options
declare -r E_OPTERR=85
declare -r script_name=${0##*/}
declare -r short_opts="f:,m::,s:t:vVh"
declare -r long_opts="force-go-path-creation:,project-main-folder::,project-sub-folders:,touch-default-files:,verbose,version,help"

[[ -z ${GOPATH} ]] && {
  printf '%s\n' "ERROR: Something is wrong with your current configuration:"
  printf '%s\n' "Your configuration is currently missing the GOPATH variable"
  printf '%s\n' "Please setup your computer before continuing !"
  exit 1
}
