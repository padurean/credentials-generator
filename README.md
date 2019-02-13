# CLI to generate user, password and ACL credentials

## Usage

`./credentials-generator <nb-credentials> <username-prefix> <counter-start>`

Example:

`./credentials-generator 25 combox- 26`

Output will be saved in a _credentials.json_ file

## Pre-built binaries

Prebuilt binaries are already available for MacOS (right in this folder),
Linux and Windows (see the corresponding subfolders)

### Build it yourself

- On your platform: `go build`
- For Linux: `env GOOS=linux GOARCH=amd64 go build -o linux-amd64/credentials-generator`
- For Windows: `env GOOS=windows GOARCH=386 go build -o windows/credentials-generator.exe`