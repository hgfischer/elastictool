# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

  config.vm.box = 'ubuntu/xenial64'
  config.vm.provider 'virtualbox' do |vb|
    vb.gui = false
    vb.cpus = 1
    vb.memory = 512
  end
  
  config.vm.provision 'file', source: '.vagrant/defaults', destination: '/tmp/defaults'
  config.vm.provision 'file', source: '.vagrant/limits.conf', destination: '/tmp/limits.conf'
  config.vm.provision 'file', source: '.vagrant/elasticsearch.yml', destination: '/tmp/elasticsearch.yml'
  config.vm.provision 'shell', path: '.vagrant/bootstrap.sh'

  config.vm.define 'elasticsearch01', autostart: true do |host|
    host.vm.hostname = 'elasticsearch01'
    host.vm.network 'private_network', ip: '172.20.192.11'
  end

  config.vm.define 'elasticsearch02', autostart: true do |host|
    host.vm.hostname = 'elasticsearch02'
    host.vm.network 'private_network', ip: '172.20.192.12'
  end

  config.vm.define 'elasticsearch03', autostart: true do |host|
    host.vm.hostname = 'elasticsearch03'
    host.vm.network 'private_network', ip: '172.20.192.13'
  end
end
