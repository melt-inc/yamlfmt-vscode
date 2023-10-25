$env:DOCKER_BUILDKIT = 1
docker build --tag vsce "https://github.com/microsoft/vscode-vsce.git#main"
