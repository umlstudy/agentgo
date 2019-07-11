#/bin/bash
# cat run.sh | sed s/^M//g | bash -
git pull
go build
kill `ps -ef | grep './gateway &' | awk '{print $2}'`
./gateway &