# go module 프로그래밍 관련 설명

https://jusths.tistory.com/107

# import 여러가지 방법
import "fmt"
or
import f "fmt"
or
import . "fmt"
or
import "website.com/Owner/blog/app/models"
type Category models.Category

https://stackoverflow.com/questions/39491435/how-to-import-structs-from-another-package-in-go?rq=1

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