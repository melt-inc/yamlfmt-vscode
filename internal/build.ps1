$env:GOOS = 'js'
$env:GOARCH = 'wasm'

go build -o="yamlfmt.wasm"
