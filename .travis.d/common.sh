#!/usr/bin/env bash
#
# Common constants, functions and helpers.
set -o nounset
set -o errexit -o errtrace

readonly turn_color_off="\033[0m"
readonly blue="\033[0;34m"
readonly green="\033[1;32m"
readonly yellow="\033[0;33m"
readonly red="\033[1;31m"

function info
{
    local message=$1
    echo -e "${blue}[INFO] ${message}${turn_color_off}"
}

function warning
{
    local message=$1
    echo -e "${yellow}[WARN] ${message}${turn_color_off}"
}

function success
{
    local message=$1
    echo -e "${green}[SUCCESS] ${message}${turn_color_off}"
}

function error
{
    local message=$1
    echo -e "${red}[SUCCESS] ${message}${turn_color_off}"
}
