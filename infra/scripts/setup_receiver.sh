#!/bin/sh

# Create .env file used by receiver exec
cat << EOF > /.env
DISCORD_TOKEN=${discord_token}
DISCORD_CHANNEL=${discord_channel}
START_HOOK=${start_hook}
STOP_HOOK=${stop_hook}
START_MESSAGE=${start_message}
STOP_MESSAGE=${stop_message}
EOF

# Create systemctl file
cat << EOF > /etc/systemd/system/server-kun.service
[Unit]
Description = server-kun
After = network-online.target

[Service]
ExecStart = /receiver
WorkingDirectory = /

User = root
Restart = always

Type = simple

[Install]
WantedBy = multi-user.target
EOF

apt install -y unzip
curl -L ${zip_url} > /receiver.zip
unzip /receiver.zip
sudo systemctl enable server-kun
sudo systemctl start server-kun
