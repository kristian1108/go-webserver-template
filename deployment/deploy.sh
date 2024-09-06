#!/bin/bash
# This file should be run from the project root
# Example usage: deployment/deploy.sh

set -e

main() {
  CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o deployment/bin/test main.go

  echo -n "Sudo password: "
  read -s password
  echo

  ansible-playbook deployment/deploy.yaml \
                    --extra-vars "ansible_sudo_pass=$password source_dir=$(pwd)" \
                    -i deployment/hosts.yaml
}

main "$@"