# Server-kun
Server-kun is a simple terraform template that makes you handle Google Compute Engine(GCE) instance with Discord easily.

[日本語](https://github.com/kypkyp/server-kun/blob/master/README.ja.md)

## About this project
Cloud is super good environment for game server. However, most game needs some high-spec servers and it costs you lots of money.

GCE is almost free while the instance is stopped (except for incredibly expensive static IP cost). So you can save money if you stop your server properly, but it's super bothering. Most of the time you just forget to stop your server and pay extra money for nothing.

Server-kun helps you save up-time by allowing you to start and stop servers easily from your Discord channel.

Server-kun has following features:

- **Flexible.** Server-kun doesn't prepare the game server itself. Server-kun just needs the name and project of the instance to work, so you can set the specs of your instance freely, and you can even add and remove Server-kun with your existing game server.

- **(Almost) free.** By using [GCP free quota](https://cloud.google.com/free) freely, server-kun can be installed with almost no additional cost.

- ~~**Easy to install.**~~ Maybe not so much for now, but I want to get this app more easy to install. Your contribution is super appreciated! 

## Requirements
Server-kun doesn't prepare the game server itself. Before installing this, you have to create GCP project and game server that fulfills the conditions below:

- Have static public IP address.
- Must start the game automatically (e.g. using systemctl).

And you have to create Discord bot account. Log in Discord and access [Discord Developer Portal](https://discord.com/developers/applications). After creating your bot, please note the token shown as "PUBLIC KEY".

## Getting Started

Clone this repository.

```
$ git clone https://github.com/kypkyp/server-kun
$ cd server-kun
```

And create IAM Service Account from [GCP Console](https://console.cloud.google.com/iam-admin/serviceaccounts/) or [CLI tool](https://cloud.google.com/sdk/gcloud/reference/iam/service-accounts/create). The account must have admin rights for GCE and Cloud Function.

After creating account, download the authorization key and move it to `infra/credentials/key.json` (or optionally you can choose specific path by configuring Terraform variable).

```
$ mkdir infra/credentials
$ cp /path/to/key infra/credentials/key.json
```

Next you should set tfvars.

$ mv infra/variables.tfvars.example infra/variables.tfvars
$ vim infra/variables.tfvars

```
# GCP Project ID of target server.
project = "you-have-to-set-this-123456"

# GCP zone of target server.
target_zone = "asia-northeast1-b"

# Instance name of target server.
target_instance_name = "my-game-server"

# Discord token (called as "public key" as well) of your bot account.
discord_token = "tH1S.is.F4K3.D15CoRdt0KeN.AnD.13CAl1eD.asPubl1cK3YasW3ll"

# Discord channel ID. See how to get your channel ID from:
# https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-
discord_channel = 782548249026232340
```

And check validity by `terraform plan`. If the changeset seems legit, it's time to go!

$ terraform plan
$ terraform apply

## Contribution

Every contribution is welcomed! If you want to send PR, fork this repository and edit it instead of committing into this repository.

## License

Distributed under the MIT License. See LICENSE for more information.
