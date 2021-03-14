#!/usr/bin/env bash

RASP_IP=192.168.1.40
INSTALL_DIR=/opt/homekit

ssh ubuntu@$RASP_IP 'sudo systemctl stop homekit'

scp ./output/homekit ubuntu@$RASP_IP:$INSTALL_DIR
echo "🍺 app copied into 192.168.1.40:/$INSTALL_DIR"

ssh ubuntu@$RASP_IP 'sudo systemctl restart homekit'

