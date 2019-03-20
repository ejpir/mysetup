#!/bin/bash
NN__NIC=$(ifconfig | grep -B1 '\tinet 10\.93' | grep -oE "utun\d+")
NN__IP=$(ifconfig | grep -B1 '\tinet 10\.93' | grep -oE "\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b" |head -1)
GW=$(netstat -rn |grep -i default |awk '{ print $2}' |tail -n 2 |head -n 1)

route delete default
route -n flush
route delete -net 0.0.0.0 $NN__IP -ifp $NN__NIC
route add -net 10.0.0.0 $NN__IP -ifp $NN__NIC
route add -net 0.0.0.0 $GW -ifp en0

scutil << EOF
  open
  d.init
  get State:/Network/Service/$NN__NIC/DNS
  d.remove ServerAddresses
  d.add ServerAddresses 127.0.0.1 *
  set State:/Network/Service/$NN__NIC/DNS
  quit
EOF
