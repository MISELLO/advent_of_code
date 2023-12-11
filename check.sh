#!/bin/bash

# This script shows style suggestions on your go code.

clear
echo "*** goimports -d ***"
goimports -d .
echo "*** golint ***"
golint .
