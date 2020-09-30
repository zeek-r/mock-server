#!/bin/bash
while getopts ":s:" o; do
    case "${o}" in
    s)
        STAGE=${OPTARG}
        ;;
    esac
done

shift $((OPTIND - 1))

set -x
set -e

echo "Deploying to AWS region: $REGION"
sls deploy -v --region "$REGION" --stage "$STAGE" --user apollo

# bash send-to-datafridge.bash -r "$REGION"
