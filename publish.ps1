cd .\internal
. .\build.ps1
cd ..

# docker run --rm -it -v ${PWD}:/workspace vsce package --pre-release
# docker run --rm -it -v ${PWD}:/workspace vsce publish --pre-release

docker run --rm -it -v ${PWD}:/workspace vsce package
docker run --rm -it -v ${PWD}:/workspace vsce publish
