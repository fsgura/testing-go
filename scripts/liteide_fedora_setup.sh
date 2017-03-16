#!/bin/bash
#
# Script to easily download and install the LiteIDE development environment,
# downloading the given release from github source repository
#
# Default values for global variables
: ${DOWNLOADS_FOLDER:="${HOME}/Downloads"
: ${LITEIDE_GITHUB:="https://github.com/visualfc/liteide/releases/download"}
: ${LITEIDE_RELEASE:="x31/liteidex31.linux64-qta4.tar.bz2"}
: ${DESTINATION_FOLDER:="/opt"}
: ${DESKTOP_SETUP:="true"}
#
DESKTOP_FILE_TEMPLATE='[Desktop Entry]
Type=Application
Name=LiteIDE
Comment=IDE for editing and building projects written in the Go programming language
Exec=/opt/liteide/bin/liteide
Icon=/opt/liteide/share/liteide/welcome/images/liteide128.xpm
Terminal=false
StartupNotify=false
Categories=IDE;Development;
'
#
PREREQS=( "wget:curl" "tar" )
#
function check_prereq()
{

}

function get_extract_liteide()
{
  cd ${DOWNLOADS_FOLDER}
  wget ${LITEIDE_GITHUB}/${LITEIDE_RELEASE}
  local liteide_filename=${LITEIDE_RELEASE##/*}
  tar -xjf ${liteide_filename} -C ${DESTINATION_FOLDER}
}

function desktop_file_setup()
{
  local desktop_folder="${HOME}/Desktop"
  local desktop_source_file="${DESTINATION_FOLDER}/liteide/liteide.desktop"
  local desktop_file=${desktop_folder}/liteide.desktop

  cp ${desktop_source_file}{,.old}

  [[ -d ${desktop_folder} ]] && {
    cat <<EOF >${desktop_source_file}
${DESKTOP_FILE_TEMPLATE}
EOF
    desktop-file-validate ${desktop_source_file}
    [[ $? == 0 ]] && {
      desktop-file-install ${desktop_source_file}
      cp ${desktop_source_file} ${desktop_file}
    } || {
      printf '%s\n' "PLEASE CHECK : Error while validating ${desktop_source_file}"
    }
  } || {
    printf '%s\n' "PLEASE CHECK : ${desktop_folder} has not been found in your path."
  }
}

function main()
{
  check_prereq
  get_extract_liteide
  desktop_file_setup
}
#
main
#
