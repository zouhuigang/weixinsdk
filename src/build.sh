#!/bin/bash
go build -o weixinsdk main.go 
chmod +x weixinsdk && nohup ./weixinsdk &