# Server-kun
Server-kun はGoogle Compute Engine(GCE)インスタンスをいい感じにDiscord経由で管理してくれるシンプルなterraformテンプレートです。

## About this project
クラウドを利用してゲームサーバーを立てると便利ですが、けっこうなスペックが要求されるために一瞬でお金がなくなります（なくなりました）。

GCEは時間単位の課金体系なので、利用する時間だけサーバーを起動することによって費用を節約することができますが、毎回コンソールやコマンドラインを叩くのは本当にめんどくさいし、高確率でやるのを忘れます（忘れました）。

server-kunはDiscordのチャンネル上からサーバーを起動・停止できるようにすることで、手間を最小限にしながら起動時間の節約をサポートします。

server-kunには以下の特徴があります。

- **フレキシブル**。server-kunはゲームサーバー自体は用意せず、ゲームサーバーのインスタンス名さえあれば動作できます。なので、ゲームサーバーのOSやスペックは自由に選べますし、既に存在するゲームサーバーにserver-kunを追加したり外したりすることもできます。
- **安い**。[GCPの無料利用枠](https://cloud.google.com/free)を利用することで、通常の使用であれば**ほとんど無料**で導入することができます。
- ~~**導入がかんたん**。~~ まだ少しめんどくさいです。あなたのPRを待っています。

## Requirements
server-kunはゲームサーバー自体は用意しません。まずはGCPでプロジェクトを作成し、以下の条件を満たすゲームサーバーを作成しておく必要があります。

- 永続化されたパブリックIPアドレスを持っている。
- サーバーを起動するだけで目的のソフトウェアが立ち上がる状態になっている(例: systemctlを使用)。

また、Discordにログインした状態で[Discord Developer Portal](https://discord.com/developers/applications)からアプリケーションのユーザーを用意する必要があります。作成後、`PUBLIC KEY`の項目をメモしておきます。

## Getting Started

まず、このレポジトリをcloneします。

$ git clone https://github.com/kypkyp/server-kun
$ cd server-kun

次に、[GCP Console](https://console.cloud.google.com/iam-admin/serviceaccounts/)か[コマンドラインツール](https://cloud.google.com/sdk/gcloud/reference/iam/service-accounts/create)からTerraform用のサービスアカウントを作成します。サービスアカウントはGCE, Cloud Functionに対する管理者権限を持つ必要があります。

サービスアカウントを作成したら、鍵をダウンロードし、`infra/credentials/key.json`に保存します。

$ mkdir infra/credentials
$ cp {鍵のパス} infra/credentials/key.json

次に、terraformを動かすのに必要な環境変数をセットします。

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

最後に、planが正しく通ることを確認した上で、terraformを実行します。

$ terraform plan
$ terraform apply

## Contribution

コントリビューションはいつでも大歓迎です！PRがある場合には、このレポジトリをforkしてPRを送ってください。

## License

Distributed under the MIT License. See LICENSE for more information.
