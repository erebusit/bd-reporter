#!/usr/bin/env bash

if [ "$CIRCLECI" != "true" ]; then
  echo "CIRCLECI environment variable not found, or not set to 'true'"
  exit 0
fi

for i in "$@"; do
  case $i in
  -t=* | --token=*)
    T="${i#*=}"
    shift
    ;;
  -v=* | --version=*)
    V="${i#*=}"
    shift
    ;;
  -e=* | --environment=*)
    E="${i#*=}"
    shift
    ;;
  -p=* | --project=*)
    P="${i#*=}"
    shift
    ;;
  -s=* | --status=*)
    S="${i#*=}"
    shift
    ;;
  esac
done

TOKEN="${T:=$BD_TOKEN}"
VERSION="${V:=$CIRCLE_BUILD_NUM}"
ENVIRONMENT="${E:=production}"
PROJECT="${P:=$CIRCLE_PROJECT_REPONAME}"
STATUS="${S:=success}"

PAYLOAD="{\"project\":\"$PROJECT\", \"version\": \"$VERSION\", \"status\": \"$STATUS\", \"environment\": \"$ENVIRONMENT\"}"

curl -X POST -H "Authorization: ApiKey $BD_TOKEN" https://deployments.eu.quickci.io/deployments --data "$PAYLOAD"
