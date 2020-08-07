#!/bin/bash

# is running under WSL?
if [ "$(uname -r | grep "Microsoft")" ]
then
    API="cmd.exe /c api"
    OPENAPI_GENERATOR="cmd.exe /c openapi-generator"
    GO="go.exe"
else
    API="api"
    OPENAPI_GENERATOR="openapi-generator"
    GO="go"
fi
OAS_FILE="openapi-spec.json"

$API generate:oas --json
mv .optic/generated/openapi.json $OAS_FILE
$GO run spec_format/main.go
rm -rf openapi
$OPENAPI_GENERATOR generate -i $OAS_FILE -g go -o openapi
