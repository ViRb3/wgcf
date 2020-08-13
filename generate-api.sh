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

$API generate:oas --json
mv ".optic/generated/openapi.json" "openapi-spec.json"
$GO run "spec_format/main.go"
rm -rf "openapi"
$OPENAPI_GENERATOR generate -i "openapi-spec.json" -g go -o "openapi"
