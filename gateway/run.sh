#/bin/bash
git pull
go build
kill `ps -ef | grep './gateway &' | awk '{print $2}'`
./gateway &