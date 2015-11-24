#!/bin/bash

BIN_DIR=/usr/bin
DATA_DIR=/var/lib/influxdb
LOG_DIR=/var/log/influxdb
SCRIPT_DIR=/usr/lib/influxdb/scripts
LOGROTATE_DIR=/etc/logrotate.d

function install_init {
    cp -f $SCRIPT_DIR/init.sh /etc/init.d/influxdb
    chmod +x /etc/init.d/influxdb
}

function install_systemd {
    cp -f $SCRIPT_DIR/influxdb.service /lib/systemd/system/influxdb.service
    systemctl enable influxdb
}

function install_update_rcd {
    update-rc.d influxdb defaults
}

function install_chkconfig {
    chkconfig --add influxdb
}

id influxdb &>/dev/null
if [[ $? -ne 0 ]]; then
    useradd --system -U -M influxdb -s /bin/false -d $DATA_DIR
fi

chmod a+rX $BIN_DIR/influx*

test -d $LOG_DIR || mkdir -p $LOG_DIR
chown -R -L influxdb:influxdb $LOG_DIR

test -d $LOG_DIR || mkdir -p $DATA_DIR
chown -R -L influxdb:influxdb $DATA_DIR

# Add defaults file, if it doesn't exist
test -f /etc/default/influxdb || touch /etc/default/influxdb

# Remove legacy logrotate file, if it exists
test -f $LOGROTATE_DIR/influxd && rm -f $LOGROTATE_DIR/influxd

# Remove legacy symlink, if it exists
test -h /etc/init.d/influxdb && rm -f /etc/init.d/influxdb

# Distribution-specific logic
if [[ -f /etc/redhat-release ]]; then
    # RHEL-variant logic
    which systemctl &>/dev/null
    if [[ $? -eq 0 ]]; then
	install_systemd
    else
	# Assuming sysv
	install_init
	install_chkconfig
    fi
elif [[ -f /etc/lsb-release ]]; then
    # Debian/Ubuntu logic
    which systemctl &>/dev/null
    if [[ $? -eq 0 ]]; then
	install_systemd
    else
	# Assuming sysv
	install_init
	install_update_rcd
    fi
fi
