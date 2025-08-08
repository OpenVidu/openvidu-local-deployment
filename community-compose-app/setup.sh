#!/bin/sh

# Prepare volumes
mkdir -p /minio/data
mkdir -p /mongo/data
mkdir -p /mongo/data/
mkdir -p /egress/home/egress
chown 1001:1001 /minio /minio/data
chown 1001:1001 /mongo /mongo/data
chown 1001:1001 /egress
chown 1001:1001 /egress/home
chown 1001:1001 /egress/home/egress
