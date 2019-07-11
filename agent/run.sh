#/bin/bash
# cat run.sh | sed s/^M//g | bash -
cp agentSettings.json agentSettings.json_ 
git pull
go build
kill `ps -ef | grep './agent -host mac.sejong.asia &' | awk '{print $2}'`
./agent -host mac.sejong.asia &
cp agentSettings.json_ agentSettings.json