#!/bin/bash

if [ $# -eq 0 ]
then
   echo "Please, provide a file to check."
   exit
fi

echo "*** goimports -d ***"
goimports -d $1
echo "*** golint ***"
golint $1
