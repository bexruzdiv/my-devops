#!/bin/bash

# Run Vault server with the specified configuration and log level
vault server -config=/etc/vault.d/vault.hcl -log-level=trace >/dev/null 2>&1 &
