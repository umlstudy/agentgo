# go module 프로그래밍 관련 설명

https://jusths.tistory.com/107

# vscode test debug 

https://stackoverflow.com/questions/43092364/debugging-go-tests-in-visual-studio-code
{
    "name": "Tests",
    "type": "go",
    "request": "launch",
    "mode": "test",
    "remotePath": "",
    "port": 2346,
    "host": "127.0.0.1",
    "program": "${workspaceRoot}",
    "env": {},
    "args": [
        "main_test.go"
        ],
    "showLog": true
},
{
    "name": "Launch file",
    "type": "go",
    "request": "launch",
    "mode": "debug",
    "program": "${file}"
},