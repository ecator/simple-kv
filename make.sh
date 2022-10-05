outdir="release"
targetWin="${outdir}/simple-kv.exe"
targetLinux="${outdir}/simple-kv"

export GOARCH="amd64"
version=$(git describe --tags --abbrev=0)

[ -d "$outdir" ] && rm -rf "$outdir"

export GOOS="windows"
go build -ldflags="-w -s -X 'main.DEBUG=0' -X 'simple-kv/cmd.VERSION=$version'" -o $targetWin

export GOOS="linux"
go build -ldflags="-w -s -X 'main.DEBUG=0' -X 'simple-kv/cmd.VERSION=$version'" -o $targetLinux
