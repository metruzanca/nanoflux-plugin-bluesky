# Bluesky plugin notes

Implementation invariants for maintainers. The user-facing documentation is
`bluesky/readme.md`.

- The plugin replaces the profile RSS feed (`bsky.app/profile/{handle}/rss`),
  which is text-only: no item titles, no media, no author, and it drops an
  image/video post down to its caption. It reads the raw `app.bsky.feed.post`
  records through `com.atproto.repo.listRecords` and rebuilds entries from them.
  The stored feed URL stays the `/rss` shape (what discovery has always
  produced) and is never parsed; its handle/DID only selects the repository to
  read.
- Two origins are involved and they are **not** interchangeable:
  `public.api.bsky.app` answers `app.bsky.actor.getProfile` but returns **501**
  for `com.atproto.repo.listRecords`; `bsky.social` answers `com.atproto.repo`
  calls but **401**s `getProfile`. The plugin's `apiBase`, `repoBase`, `imgBase`
  and `videoBase` vars encode this so a test can point them at a mock host.
- Media: images use `cdn.bsky.app/img/feed_fullsize/plain/{did}/{cid}` (direct,
  cacheable). Video is attached as an **HLS** enclosure
  (`application/vnd.apple.mpegurl`, URL
  `video.bsky.app/watch/{did}/{cid}/playlist.m3u8`); the raw
  `com.atproto.sync.getBlob` MP4 is not used because it returns the whole file
  and ignores `Range`, so seeking is limited. HLS segments are CORS-open and
  Range-capable, and the stored master playlist is fetched fresh at play time
  (its session-scoped rendition URLs need no storage). The video's poster frame
  (`video.bsky.app/watch/.../thumbnail.jpg`) is the item thumbnail; it is served
  as `application/octet-stream`, so the host image proxy sniffs the bytes.
  The enclosure sets `Kind: pluginapi.EnclosureKindHLS`; the host resolves it to
  `hls` and a lazily-loaded hls.js player renders it.
- Replies are skipped, matching the RSS feed's original-posts-only behavior.
  Quotes are rendered as a link to the quoted post (not expanded: hydrating each
  quoted record would cost one request per quote). Facets (mentions, links,
  tags) are rendered as anchors from the record's byte offsets.
- Existing stored bsky feeds are adopted by `ReconcileFeeds` at startup (the URL
  shape matches), and the poller stores enclosures on **every** poll, not only on
  insert, so an already-stored feed gains its media on the next poll. The
  plugin's `Discover` supplies the candidate and preview metadata; the host has
  no site-specific discovery rule for it.
