# go

Install Golang:
```bash
sudo apt update
sudo apt install software-properties-common

sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```

Set Go Path:
```bash
mkdir -p ~/go/{bin,pkg,src}
echo 'export GOPATH="$HOME/go"' >> ~/.bashrc
echo 'export PATH="$PATH:${GOPATH//://bin:}/bin"' >> ~/.bashrc
source ~/.bashrc
```

```bash
go run ./cmd/myapp
```

https://golangbyexample.com/all-data-types-in-golang-with-examples/

