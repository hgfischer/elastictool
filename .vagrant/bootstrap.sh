#!/bin/bash

function disable_apt_interactive_mode() {
	sudo rm -v /etc/apt/apt.conf.d/70debconf
	sudo dpkg-reconfigure debconf -f noninteractive -p critical
}

function setup_apt_repositories() {
	wget -qO - https://packages.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
	echo "deb http://packages.elastic.co/elasticsearch/2.x/debian stable main" | sudo \
		tee -a /etc/apt/sources.list.d/elasticsearch-2.x.list
}

function update_upgrade_autoremove() {
	sudo apt-get update
	sudo apt-get -y upgrade
	sudo apt-get -y dist-upgrade
	sudo apt-get -y autoremove
}

function install_apt_dependencies() {
	sudo apt-get -y install openjdk-8-jdk elasticsearch
}

function configure_elasticsearch() {
	sudo cp -v /tmp/defaults /etc/default/elasticsearch
	sudo cp -v /tmp/limits.conf /etc/security/limits.d/elasticsearch.conf
	sudo cp -v /tmp/elasticsearch.yml /etc/elasticsearch/elasticsearch.yml
	
	if [ ! -e /usr/share/elasticsearch/plugins/kopf ]
	then
		/usr/share/elasticsearch/bin/plugin install lmenezes/elasticsearch-kopf/v2.1.2
	fi
	
	sudo /bin/systemctl daemon-reload
	sudo /bin/systemctl enable elasticsearch.service
	sudo /bin/systemctl start elasticsearch.service
}

disable_apt_interactive_mode
setup_apt_repositories
update_upgrade_autoremove
install_apt_dependencies
configure_elasticsearch
