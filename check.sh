#!/bin/bash

# This script shows style suggestions on your go code.

clear
date +***\ %d-%m-%Y\ %H:%M:%S\ ***
echo "*** goimports -d ***"
goimports -d .
echo "*** golint ***"
golint .
