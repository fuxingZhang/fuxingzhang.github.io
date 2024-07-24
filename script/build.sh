HERE=$(cd -P -- $(dirname -- "$0") && pwd -P)
ROOT_DIR=$(cd $HERE/.. && pwd -P)

cd $ROOT_DIR

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/server main.go

docker build -f Dockerfile -t public:latest .
