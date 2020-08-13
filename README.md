# wgcf
> wgcf is a cross-platform CLI for [Cloudflare Warp](https://1.1.1.1/) that allows you to create and 
> manage accounts, assign license keys, and generate WireGuard profiles.

![](https://img.shields.io/drone/build/ViRb3/wgcf)
![](https://img.shields.io/github/issues/ViRb3/wgcf)
![](https://img.shields.io/github/downloads/ViRb3/wgcf/total)
![](https://img.shields.io/github/languages/code-size/ViRb3/wgcf)

## Features
- Register new account
- Change license key to use existing Warp+ subscription
- Generate WireGuard profile
- Check account status
- Print trace information to debug Warp/Warp+ status

## Download
You can find pre-compiled binaries on the [releases page](https://github.com/ViRb3/wgcf/releases).

## Usage
Run `wgcf` in a terminal without any arguments to display the help screen. All commands and parameters are documented.

### Register new account
Run the following command in a terminal:
```bash
wgcf register
```
The new account will be saved under `wgcf-account.toml`

### Generate WireGuard profile
[Register an account](#register-new-account) if you haven't already. Then, run:
```bash
wgcf generate
```
The WireGuard profile will be saved under `wgcf-profile.conf`

### Change license key
If you have an existing Warp+ subscription, for an example on your phone, you can bind the account generated by this tool to your phone's account, sharing its Warp+ status. Please note that there is a current limit of maximum 5 linked devices active at a time. 

> :warning: This device's private key will be changed!

First, get your Warp+ account license key. To view it on Android:
1. Open the `1.1.1.1` app
2. Click on the hamburger menu button on the top-right corner
3. Navigate to: `Account` > `Key`

Edit `wgcf-account.toml` directly with the new license key and run:
```bash
wgcf update
```
or, using an environment variable:
```bash
WGCF_LICENSE_KEY="123412341234" wgcf update
```
The license will be applied, and a new private key will be created for your account. [Regenerate](#generate-wireguard-profile) your WireGuard configuration and use the new private key.


### Check device status
Run the following command in a terminal:
```bash
wgcf status
```

### Verify Warp/Warp+ works
Connect to the WireGuard profile [generated](#generate-wireguard-profile) by this tool, then run:
```bash
wgcf trace
```
If you look at the last line, it should say `warp=on` or `warp=plus`, depending on whether you have Warp or Warp+ respectively.

## Development
### Sub-packages
- [api_tests](api_tests/main.go) - Tests for API documentation generation
- [spec_format](spec_format/main.go) - OpenAPI3 specification formatter to post-process the spec generated by Optic
### API
This project uses [Optic](https://github.com/opticdev/optic) to automatically generate API documentation using the tests defined in [api_tests](api_tests/main.go). These tests cover all endpoints used by wgcf. The documentation is exported as an OpenAPI3 [specification](openapi-spec.json), which is then used with [openapi-generator](https://openapi-generator.tech/) to generate the Go client API code under [wgcf/openapi](openapi/client.go).

To update the API documentation, [install Optic](https://github.com/opticdev/optic/releases/latest), then run:
```bash
api start
```
Resolve and save all the differences in the Web UI.

To regenerate the Go client API code, [install openapi-generator](https://openapi-generator.tech/docs/installation), then run:
```bash
bash generate-api.sh
```
This script supports both Linux and WSL.

## Notice of Non-Affiliation and Disclaimer
We are not affiliated, associated, authorized, endorsed by, or in any way officially connected with Cloudflare, or any of its subsidiaries or its affiliates. The official Cloudflare website can be found at https://www.cloudflare.com/.

The names Cloudflare Warp and Cloudflare as well as related names, marks, emblems and images are registered trademarks of their respective owners.