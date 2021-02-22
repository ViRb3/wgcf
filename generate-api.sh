#!/bin/bash

# is running under WSL?
if [ "$(uname -r | grep -i "microsoft")" ]
then
    API="cmd.exe /c api"
    OPENAPI_GENERATOR="cmd.exe /c openapi-generator-cli"
    GO="go.exe"
else
    API="api"
    OPENAPI_GENERATOR="openapi-generator-cli"
    GO="go"
fi

$API generate:oas --json
mv ".optic/generated/openapi.json" "openapi-spec.json"
$GO run "spec_format/main.go"
rm -rf "openapi"

$OPENAPI_GENERATOR generate -i "openapi-spec.json" -g go -o "openapi"
