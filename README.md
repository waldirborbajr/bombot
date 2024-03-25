<p align="center">
  <img width="256" height="256" src="./assets/bombot-logo.png" />
</p>

<h1 align="center">bombot - GOLang Telegram Bot - Just for fun</h1>

<p align="center">
  <a href="https://github.com/waldirborbajr/bombot/actions/workflows/ci-cd.yaml">
    <img alt="tests" src="https://github.com/waldirborbajr/bombot/actions/workflows/ci-cd.yaml/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/waldirborbajr/bombot">
    <img alt="goreport" src="https://goreportcard.com/badge/github.com/waldirborbajr/bombot" />
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

**BETA:** This project is in active development. Please check out the issues and contribute if you're interested in helping out.

# TODO: rewrite this document

## bombot
`tl;dr:` `bombot`, a.k.a GOlang Symbolic Link (symlink), is an open-source software built-in with the main aim of being a personal alternative to **GNU Stow**.

As `GNU Stow`, `bombot` is a symlink farm manager which takes distinct packages of software and/or data located in separate directories on the filesystem, and makes them appear to be installed in the same place.

With `bombot` it is eeasy to track and manage configuration files in the user's home directory, especially when coupled with version control systems.

## How to install

```sh
brew install waldirborbajr/bombot/bombot
```

### Go

Alternatively, you can install bombot using Go's go install command:

```sh
go install github.com/waldirborbajr/bombot@latest
```

This will download and install the latest version of bombot. Make sure that your Go environment is properly set up.

**Note:** Do you want this on another package manager? [Create an issue](https://github.com/waldirborbajr/bombot/issues/new) and let me know!

## How to use

The main goal of `bombot` is to be as simple as that, `easy peasy lemon squeezy`, with few commands and straight to the target.

```sh
# To create a link to $HOMR
bombot l

# To force overwrite existing link : **TODO** not implemented
bombot f -f

# To remove (kill) all symblinks : **TODO** not implemented
bombot k

# To remove a specific symblinks : **TODO** not implemented
bombot r symlink-name


# To print all symlink created : **TODO** not implemented
bombot p
```

## .bombot-ignore`

You can add files/directories to ignore list, so when execute `bombot` the content will no be linked.

```sh
touch .bombot-ignore
```

## Contributing to bombot

If you are interested in contributing to `bombot`, we would love to have your help! You can start by checking out the [ open issues ](https://github.com/waldirborbajr/bombot/issues) on our GitHub repository to see if there is anything you can help with. You can also suggest new features or feel free to create a new feature by opening a new issue.

To contribute code to `bombot`, you will need to fork the repository and create a new branch for your changes. Once you have made your changes, you can submit a pull request for them to be reviewed and merged into the main codebase.

## Contributors

<a href="https://github.com/waldirborbajr/bombot/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=waldirborbajr/bombot" />
</a>

Made with [contrib.rocks](https://contrib.rocks).


# BomBot

## Starting NGrOK

```sh
ngrok http 9090
```

## Setting WebHook

```sh
curl -F "url=https://f531-2804-d55-433d-5600-10b-8c9-7d95-bc9b.ngrok.io" https://api.telegram.org/bot5343272189:AAF5_yv9adxzqsNrYCqAY5jakgb4GqZFGBc/setWebhook
```

## Deleting WebHook

```sh
curl  https://api.telegram.org/bot5343272189:AAF5_yv9adxzqsNrYCqAY5jakgb4GqZFGBc/deleteWebhook
```
