#!/bin/bash
source $(dirname $0)/scripts/env.sh
FORK=elderberry
DATA_AVAILABILITY_MODE=$1
if [ -z $DATA_AVAILABILITY_MODE ]; then
    echo "Missing DATA_AVAILABILITY_MODE: ['rollup', 'cdk-validium']"
    exit 1
fi
BASE_FOLDER=$(dirname $0)

docker images -q cdk:latest > /dev/null
if [ $? -ne 0 ] ; then
    echo "Building cdk:latest"
    pushd $BASE_FOLDER/..
    make build-docker
    popd
else
    echo "docker cdk:latest already exists"
fi

$BASE_FOLDER/scripts/kurtosis_prepare_params_yml.sh "$KURTOSIS_FOLDER" "elderberry" "cdk-validium"
[ $? -ne 0 ] && echo "Error preparing params.yml" && exit 1

kurtosis clean --all
kurtosis run --enclave cdk-v1 --args-file $DEST_KURTOSIS_PARAMS_YML --image-download always $KURTOSIS_FOLDER
#[ $? -ne 0 ] && echo "Error running kurtosis" && exit 1
echo "Waiting 10 minutes to get some verified batch...."
scripts/batch_verification_monitor.sh 0 600