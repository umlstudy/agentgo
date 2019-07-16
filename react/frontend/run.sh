#/bin/bash
# sed 's/\r//g' run.sh | bash -
git pull
npm run build
kill `ps -ef | grep 'serve -s build' | grep -v grep | awk '{print $2}'`
serve -s build &