#!/bin/bash

if [[ -f /etc/opt/influxdb/influxdb.conf ]]; then
    # Legacy configuration found
    
    # Create new configuration location (if not already there)
    test -d /etc/influxdb || mkdir -p /etc/influxdb
    
    if [[ ! -f /etc/influxdb/influxdb.conf ]]; then
	# New configuration does not exist, move legacy configuration to new location
	cp --backup --suffix=.$(date +%s).install-backup -a /etc/opt/influxdb/* /etc/influxdb/
    fi
fi
