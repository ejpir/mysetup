
Setup tested on MacOS mojave!

- brew install openconnect dnsmasq                                                        
- wget  https://github.com/ejpir/juniper-vpn-py/blob/add-periodic-host-checker/tncc.py (periodic host checker)
- pip install -r https://github.com/ejpir/juniper-vpn-py/blob/master/requirements.txt (as root for missing mechanize)
- openssl s_client -showcerts -connect `<server>:<port>` </dev/null 2>/dev/null|openssl x509 -outform PEM >mycertfile.pem
- (optional) modify dnsmasq.conf
- (optional) use custom resolver from resolver/ and change the filename and contents to the vpn domain
- modify/run vpn.sh (sudo)
- (optional) run split-tunnel.sh (if not using v2)

