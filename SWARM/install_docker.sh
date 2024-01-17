#!/bin/bash

# Docker CE o'rnatish uchun avtomatlashtirilgan skript

# 1. yum-utils paketini o'rnatish
sudo yum install -y yum-utils

# 2. Docker repo-ni qo'shish
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

# 3. Docker CE o'rnatish
sudo yum install -y docker-ce docker-ce-cli containerd.io

# 4. Docker ishga tushirish
sudo systemctl start docker

# 5. Foydalanuvchi guruhiga qo'shish
sudo usermod -aG docker $USER

# 6. O'zingizni kirishingiz uchun sistemaning qayta yuklanishini bajarish
su - $USER

# 7. Docker-versiyasini tekshirish
docker --version

