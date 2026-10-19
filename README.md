# nanoflux-plugin-bluesky

A nanoflux external plugin for Bluesky profile feeds. It reads the account's
posts through the AT Protocol and rebuilds titles, text, images, video, and
attribution that the profile RSS feed omits.

## Install

```bash
cd nanoflux/plugins/bluesky
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../nanoflux-plugin-bluesky .
```

Restart nanoflux after putting the binary in its plugins directory
(`NF_PLUGINS_DIR`, `/plugins` in the container).

## Usage

Add a profile RSS URL, for example `https://bsky.app/profile/handle/rss`.
