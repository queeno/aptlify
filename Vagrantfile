# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # boxes at https://atlas.hashicorp.com/search.
  config.vm.provider "virtualbox" do |vb|
    config.vm.box = "ubuntu/xenial64"
  end
  config.vm.provider "vmware" do
    config.vm.box = "geerlingguy/ubuntu1604"
  end
  config.vm.provision "shell", inline: <<-SHELL
  echo "deb http://repo.aptly.info/ squeeze main" >> /etc/apt/sources.list.d/aptly.list
  echo "deb http://ftp.uk.debian.org/debian/ stretch universe main" >> /etc/apt/sources.list.d/golang.list
  apt-key adv --keyserver keys.gnupg.net --recv-keys 9E3E53F19C7DE460 8B48AD6246925553
  apt-get update
  apt-get install -y aptly golang-1.7 git docker.io
  usermod -a -G docker vagrant
  ln -s /usr/lib/go-1.7/bin/go /usr/bin/go
  mkdir -p /home/vagrant/gowork/src/github.com/queeno
  ln -s  /home/vagrant/gowork/src/github.com/queeno/aptlify /aptlify
  GOPATH=/home/vagrant/gowork go get github.com/mattn/gom
  echo "export GOPATH=/home/vagrant/gowork" >> /etc/profile
  echo 'export PATH=$PATH:/home/vagrant/gowork/bin' >> /etc/profile
  chown -R vagrant: /home/vagrant || true
  SHELL
  config.vm.synced_folder ".", "/home/vagrant/gowork/src/github.com/queeno/aptlify"
end
