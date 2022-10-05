
$outdir = "release"
$targetWin="${outdir}\simple-kv.exe"
$targetLinux="${outdir}\simple-kv"


$version = $(git describe --tags --abbrev=0)

$env:GOARCH="amd64"

if (Test-Path $outdir){
    Remove-Item -Recurse $outdir
}

$env:GOOS="windows"
go build -ldflags="-w -s -X 'main.DEBUG=0' -X 'simple-kv/cmd.VERSION=$version'" -o $targetWin

$env:GOOS="linux"
go build -ldflags="-w -s -X 'main.DEBUG=0' -X 'simple-kv/cmd.VERSION=$version'" -o $targetLinux