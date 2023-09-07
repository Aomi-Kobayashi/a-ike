set GOOS=windows
set GOARCH=386
set FUNC=test

go build -o %FUNC%.exe


pause

%FUNC%.exe

pause