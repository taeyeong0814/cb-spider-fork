#!/bin/bash

if [ "$1" = "" ]; then
        echo
        echo -e 'usage: '$0' mock|aws|azure|gcp|alibaba|tencent|ibm|openstack|cloudit|ncp|nhncloud number'
        echo -e '\n\tex) '$0' aws'
        echo
        exit 0;
fi

source ../common/setup.env $1
source setup.env $1

echo -e "\n\n"
echo -e "###########################################################"
echo -e "# Try to suspend $1 VM"
echo -e "###########################################################"
echo -e "\n\n"


../common/3-1.vm-suspend.sh $1 

#### Check sync called
ret=`../common/2.vm-getstatus.sh $1 2>&1 | grep Suspended`

if [ "$ret" ];then
        echo -e "\n-------------------------------------------------------------- $0 $1 : pass"
else
        echo -e "\n-------------------------------------------------------------- $0 $1 : fail"
fi

echo -e "\n\n"
