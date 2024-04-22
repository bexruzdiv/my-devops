terraform {
  cloud {
    organization = "Proximaops"

    workspaces {
      name = "proxima_ansible"
    }
  }
}