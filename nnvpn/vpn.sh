#!/bin/bash
export TNCC_CERTS="/Users/nick/<cert>.pem"
export TNCC_DEVICE_ID="<mac>"
export TNCC_FUNK="1"
export TNCC_HOSTNAME="<laptop>"
export TNCC_PLATFORM="Windows ??"
openconnect <server> --useragent 'Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36' -u <user> --protocol=nc  --csd-wrapper=./tncc.py  --servercert pin-sha256:<pinned_server_cert> -s ./split-tunnel-v2.sh
