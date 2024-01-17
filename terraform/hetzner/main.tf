#################
# PRIVATE NETWORK
#################

resource "hcloud_network" "network" {
  name     = "network"
  ip_range = "10.0.0.0/16"
}

################
# SUBNETS
################


resource "hcloud_network_subnet" "network-subnet-1" {
  type         = "cloud"
  network_id   = hcloud_network.network.id
  network_zone = "eu-central"
  ip_range     = "10.0.1.0/24"
}






##############
#MASTER-1
##############



resource "hcloud_server" "master" {
  name       = "master1"
  image      = "ubuntu-22.04"
  server_type = "cpx11"
  location   = "hel1"
  ssh_keys   = ["bexruz@bexruz-HP-Laptop-15-dy2xxx"]
#  public_net {
#    ipv4_enabled = true
#    ipv6_enabled = true
#  }
  network {
    network_id = hcloud_network.network.id
    ip         = "10.0.1.5"

  }


  depends_on = [
    hcloud_network_subnet.network-subnet-1
  ]
}




##################
# WORKER-1
##################


resource "hcloud_server" "worker-1" {
  name       = "worker-1"
  image      = "ubuntu-22.04"
  server_type = "cpx11"
  location   = "hel1"
  ssh_keys   = ["bexruz@bexruz-HP-Laptop-15-dy2xxx"]
#  public_net {
#    ipv4_enabled = true
#    ipv6_enabled = true
#  }
  network {
    network_id = hcloud_network.network.id
    ip         = "10.0.1.4"


  }
  


  depends_on = [
    hcloud_network_subnet.network-subnet-1
  ]
}






###################
#   PUBLIC-SERVER
###################


resource "hcloud_server" "haproxy-server" {
  name       = "haproxy-server"
  image      = "ubuntu-22.04"
  server_type = "cpx11"
  location   = "hel1"
  ssh_keys   = ["bexruz@bexruz-HP-Laptop-15-dy2xxx"]

  network {
    network_id = hcloud_network.network.id
    ip         = "10.0.1.7"

  }


  depends_on = [
    hcloud_network_subnet.network-subnet-1
  ]
}


#S7F225taF86ntiNiMooVeW0EM8d7n2E7El4Zc9DY1hJNev2pfoD0xcsPDBB8UJPK
