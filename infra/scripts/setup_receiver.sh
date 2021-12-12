#!/bin/sh

cat << EOF > /.env
DISCORD_TOKEN=${discord_token}
DISCORD_CHANNEL=${discord_channel}
START_HOOK=${start_hook}
STOP_HOOK=${stop_hook}
START_MESSAGE=${start_message}
STOP_MESSAGE=${stop_message}
EOF

apt install -y unzip
curl -L ${zip_url} > /receiver.zip
unzip /receiver.zip
/receiver
