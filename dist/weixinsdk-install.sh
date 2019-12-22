#!/bin/bash
systemctl stop weixinsdk
yum remove weixinsdk -y
yum localinstall -y weixinsdk-1.0.2-1.x86_64.rpm
cd  /usr/local/software/weixinsdk && chmod +x weixinsdk
systemctl daemon-reload
systemctl start weixinsdk
systemctl status weixinsdk
