#!/bin/sh
set -e

# Generate Caddyfile and index.html
CURRENT_DIR="$(pwd)"
TMP_DIR="/tmp/caddy-local"
mkdir -p "$TMP_DIR"
cd "$TMP_DIR"
/usr/bin/local-caddy-generate
if [ ! -f /var/www/index.html ]; then
  mkdir -p /var/www
  cp "$TMP_DIR/index.html" /var/www/index.html
fi
if [ ! -f /var/www/app502.html ]; then
  mkdir -p /var/www
  cp "$TMP_DIR/app502.html" /var/www/app502.html
fi
if [ ! -f /config/caddy/Caddyfile ]; then
  mkdir -p /config/caddy
  cp "$TMP_DIR/Caddyfile" /config/caddy/Caddyfile
fi
cd "$CURRENT_DIR"
rm -rf /tmp/caddy-local

# Start Caddy
exec "$@"
