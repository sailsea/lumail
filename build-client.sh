#!/usr/bin/env bash

dir=lumail-cli

#windows32
mkdir -p $dir/lumail-windows-386
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o lumail.exe
sleep 1
zip $dir/lumail-windows-386.zip lumail.exe
sleep 1
mv lumail.exe $dir/lumail-windows-386/lumail.exe

#windows64
mkdir -p $dir/lumail-windows-amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o lumail.exe
sleep 1
zip $dir/lumail-windows-amd64.zip lumail.exe
sleep 1
mv lumail.exe $dir/lumail-windows-amd64/lumail.exe

#linux32
mkdir -p $dir/lumail-linux-386
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o lumail
sleep 1
zip $dir/lumail-linux-386.zip lumail
sleep 1
mv lumail $dir/lumail-linux-386/lumail

#linux64
mkdir -p $dir/lumail-linux-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lumail
sleep 1
zip $dir/lumail-linux-amd64.zip lumail
sleep 1
mv lumail $dir/lumail-linux-amd64/lumail

#mac32
mkdir -p $dir/lumail-darwin-386
CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o lumail
sleep 1
zip $dir/lumail-darwin-386.zip lumail
sleep 1
mv lumail $dir/lumail-darwin-386/lumail

#mac64
mkdir -p $dir/lumail-darwin-amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o lumail
sleep 1
zip $dir/lumail-darwin-amd64.zip lumail
sleep 1
mv lumail $dir/lumail-darwin-amd64/lumail
