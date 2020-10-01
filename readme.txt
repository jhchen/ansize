brew install go
mkdir /usr/local/lib/go
export GOPATH=/usr/local/lib/go
go get github.com/nfnt/resize
git clone git@github.com:jhchen/ansize.git
go build ansize.go
