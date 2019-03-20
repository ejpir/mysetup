#!/bin/sh
# Initialize empty split tunnel list
export CISCO_SPLIT_INC=0

# Delete DNS info provided by VPN server to use internet DNS
# Comment following line to use DNS beyond VPN tunnel
unset INTERNAL_IP4_DNS

export CISCO_SPLIT_INC_${CISCO_SPLIT_INC}_ADDR=10.0.0.0
export CISCO_SPLIT_INC_${CISCO_SPLIT_INC}_MASK=255.0.0.0
export CISCO_SPLIT_INC_${CISCO_SPLIT_INC}_MASKLEN=8
export CISCO_SPLIT_INC=$(($CISCO_SPLIT_INC + 1))

# Execute default script
. /usr/local/etc/vpnc-script
