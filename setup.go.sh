# install go for ubuntu
curl -LO https://go.dev/dl/go1.25.0.linux-amd64.tar.gz

# unzip it at global location
sudo tar -C /usr/local -xzf go1.25.0.linux-amd64.tar.gz

# register the path in current terminal session
export PATH=$PATH:/usr/local/go/bin

# copy/paste the path in profile: export PATH=$PATH:/usr/local/go/bin
nano ~/.bashrc 
source ~/.bashrc 
go version


# check package install path
go list -f '{{.Target}}'

# install cgo runtime for race condition check flag: go run -race <myfile.go>
sudo apt update
sudo apt install build-essential
