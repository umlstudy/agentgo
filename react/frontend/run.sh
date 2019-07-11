#/bin/bash
# cat run.sh | sed s/^M//g | bash -
git pull
npm run build
kill `ps -ef | grep 'serve -s build &' | awk '{print $2}'`
serve -s build &