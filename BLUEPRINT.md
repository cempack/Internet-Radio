# Internet Radio Crate Digger

A full product and engineering blueprint for a keyboard-first terminal app that helps users discover, save, and listen to internet radio stations from around the world.

## Product vision

Internet Radio Crate Digger is a cozy, high-agency terminal application for music discovery. The product is not just a stream picker; it is a browsing environment built for serendipity, curation, and long-running sessions.

The app should feel like a late-night record shop inside the terminal: warm, tactile, fast, and deeply keyboard driven. Lip Gloss is the styling and layout layer, Bubbles provides reusable TUI components, and lazyspotify shows that a media app can feel elegant and premium in a terminal-first interface.[cite:1][cite:6][cite:13]

## Core promise

The app should let a user do four things exceptionally well:

- Discover radio stations by country, city, language, tag, codec, bitrate, and popularity.
- Drop into listening quickly with near-zero friction.
- Build a personal crate of favorites, presets, and saved sessions.
- Stay inside a rich “now playing” environment that is pleasant enough to keep open for hours.

## Experience principles

### 1. Discovery first

The product wins on browsing quality, not just playback. Every screen should help the user answer one of these questions:

- What should be played next?
- What kind of station is this?
- What similar stations exist nearby in the graph of genre, geography, and mood?
- What did the user like enough to save?

### 2. Cozy, not corporate

Avoid the look of a dashboard. The interface should feel like a crafted music tool, not a developer admin panel.

### 3. Fast paths everywhere

Every major action should be available in one to three keystrokes:

- Search.
- Tune.
- Favorite.
- Open queue/history.
- Jump to related stations.
- Toggle presets.
- Start random mode.

### 4. Designed for dwell time

A great version of this app stays open in a terminal tab all evening. That means the idle state matters: the “now playing” view, metadata updates, station health, listening history, and visual rhythm should all feel alive.

## User archetypes

### Crate digger

A user who wants obscure stations, weird geography, and niche tags.

Needs:
- Deep search and filtering.
- Good related-station recommendations.
- Random exploration modes.
- History and save tools.

### Focus listener

A user who wants background audio while working.

Needs:
- Reliable presets.
- Quick resume.
- Low-friction favorites.
- Stable playback and reconnection.

### World explorer

A user who surfs by country, language, and city.

Needs:
- Region browsing.
- Strong metadata display.
- Country and language facets.
- Easy backtracking.

### Taste builder

A user who wants to cultivate a personal library of stations and moments.

Needs:
- Tagging and notes.
- Saved collections.
- Listening journal.
- Import/export.

## Feature set

### MVP

- Station search.
- Genre/tag filters.
- Country/language filters.
- Station detail panel.
- Playback controls.
- Favorites.
- Recent history.
- Presets.
- Random station mode.
- Related stations.
- Stream health indicator.
- Configurable external player backend.

### V1

- Local library of saved crates.
- Smart “more like this” browsing.
- Keyboard palette.
- Session restore.
- Notes on stations.
- Sleep timer.
- Recording buffer or clip bookmarks.
- Mini mode for small terminals.

### V2

- Community station packs.
- Shared crates.
- Collaborative listening room.
- Daily discovery ritual.
- Reputation or taste profile.
- Lightweight social layer.

## Information architecture

The app should use a multi-view layout with one primary shell and focused subviews.

### Global views

- Home.
- Discover.
- Search results.
- Station detail.
- Now playing.
- Favorites.
- History.
- Presets.
- Crates.
- Settings.
- Help / shortcuts.

### Recommended shell layout

#### Standard desktop terminal layout

Use a three-column structure:

- Left column: navigation, filters, active facets, saved crates.
- Center column: station list or search results.
- Right column: now playing, station detail, metadata, related stations, shortcuts.

#### Narrow terminal layout

Collapse into stacked or tabbed panes:

- Top: current view title and mode.
- Middle: main list.
- Bottom: player strip and context actions.

#### Modal overlays

Use modal or drawer-style overlays for:

- Command palette.
- Advanced filters.
- Add note.
- Rename preset or crate.
- Confirm delete.
- Help cheatsheet.

## Navigation model

### Primary navigation

Suggested keys:

- `1` Home
- `2` Discover
- `3` Favorites
- `4` History
- `5` Presets
- `6` Crates
- `/` Search
- `?` Help
- `:` Command palette
- `q` Back or quit contextually

### List navigation

- `j` / `k` move
- `gg` top
- `G` bottom
- `enter` tune/open
- `space` preview or play/pause depending on context
- `f` favorite
- `p` save to preset
- `c` save to crate
- `r` related stations
- `o` open detail
- `y` copy station URL

### Player navigation

- `tab` cycle panes
- `[` previous station in history
- `]` next station in history
- `m` mute
- `-` volume down
- `=` volume up
- `s` sleep timer
- `x` stop stream

## Core flows

### Flow: first-time use

1. Launch app.
2. Show a welcoming home screen with a few curated discovery entry points.
3. Prompt for preferred player backend if none is configured.
4. Offer quick picks such as “browse by country,” “browse by tag,” “random station,” and “favorites.”
5. Let the user start listening within 10 seconds.

### Flow: search and tune

1. User presses `/`.
2. Search box appears with incremental results.
3. Filters can be applied inline: country, language, bitrate, codec, votes, tags.
4. User selects a station.
5. App starts playback, moves station into history, and opens the now playing side panel.

### Flow: discovery loop

1. User tunes to a station.
2. Right panel shows related stations.
3. User jumps by tag, language, country, or similar-name cluster.
4. User favorites or crates good finds.
5. Session gradually becomes a curated trail.

### Flow: focus mode

1. User opens presets or favorites.
2. Picks a reliable stream.
3. Minimizes into compact now playing mode.
4. Leaves app running while working.

## Differentiators

### 1. Radio graph browsing

Do not treat stations as flat search rows. Build lightweight relationships:

- Same country.
- Same language.
- Overlapping tags.
- Similar bitrate/codec.
- Similar popularity.
- Shared words in station names.

This allows “digging” rather than merely filtering.

### 2. Crates instead of playlists

A crate is a curated collection of stations around a theme, mood, region, or task.

Examples:
- Night driving.
- Japanese city pop.
- Francophone public radio.
- Deep ambient.
- Balkan oddities.
- Coding focus.

### 3. Ritualized discovery

Add small loops that create attachment:

- Station of the day.
- One random country prompt.
- “Keep digging” suggestions after 10 minutes.
- Weekly listening recap from local history.

### 4. Beautiful idle state

The now playing screen should feel good even when the user is not actively navigating. This is where the product identity lives.

## Visual design direction

### Tone

- Warm.
- Noisy in a tasteful way.
- Retro radio meets modern terminal craft.
- Slightly nocturnal.
- Human, not gamer RGB.

### Styling goals with Lip Gloss

Lip Gloss is the styling layer for spacing, borders, colors, alignment, and composed layouts.[cite:1]

Use it to create:

- Distinct pane boundaries without heavy box-drawing clutter.
- Soft color hierarchy with one accent color and muted neutrals.
- Dense but readable lists.
- Strong title blocks and status badges.
- Clear focus state for the active pane.

### Suggested palette

Dark mode first:

- Background: near-black with a warm tint.
- Surface: charcoal or deep brown-gray.
- Primary text: soft off-white.
- Muted text: dusty gray.
- Accent: amber, teal, or desaturated green.
- Status colors: subdued, never neon.

### UI motifs

- Big station title in the detail panel.
- Small badges for country, language, codec, bitrate, tags.
- Signal or health indicator for stream reliability.
- Textual sparkline or miniature stat row for popularity or uptime.
- Sticky player strip at bottom or right panel.

## Component blueprint

Bubbles provides reusable TUI components, which makes it a strong fit for the input and navigation primitives in this app.[cite:6]

### Likely component mapping

| Need | Candidate approach |
|------|--------------------|
| Search input | text input component |
| Large scrollable result list | list or custom viewport-based list |
| Help and keyboard cheatsheet | viewport or styled text view |
| Filter chips / toggles | custom component with Lip Gloss styling |
| Progress / loading state | spinner and status line |
| Confirmation overlays | custom modal shell |
| Tabs or pane switcher | custom tabs using Lip Gloss + key bindings |
| Long metadata text | viewport |

### Custom components to build

- Station row renderer.
- Pane header.
- Now playing card.
- Metadata grid.
- Filter chip bar.
- Preset card.
- Crate card.
- Toast/status line.
- Command palette.
- Mini player strip.

## Domain model

### Station

```go
type Station struct {
    ID               string
    Name             string
    StreamURL        string
    HomepageURL      string
    FaviconURL       string
    Country          string
    CountryCode      string
    State            string
    Language         string
    Tags             []string
    Codec            string
    Bitrate          int
    Votes            int
    ClickCount       int
    LastCheckOK      bool
    LastCheckTime    time.Time
    LastCheckStatus  string
    ResolvedURL      string
}
```

### Crate

```go
type Crate struct {
    ID          string
    Name        string
    Description string
    Tags        []string
    StationIDs   []string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Preset

```go
type Preset struct {
    Key         string
    Name        string
    StationID   string
    CreatedAt   time.Time
}
```

### History item

```go
type HistoryItem struct {
    StationID    string
    StartedAt    time.Time
    EndedAt      time.Time
    Duration     time.Duration
    Context      string
}
```

### App config

```go
type Config struct {
    PlayerBackend       string
    VolumeStep          int
    Theme               string
    CompactMode         bool
    StartView           string
    AutoReconnect       bool
    ReconnectAttempts   int
    SaveHistory         bool
    SaveSession         bool
    CountryBias         []string
    TagBias             []string
}
```

## Suggested architecture

### Recommended stack

- Go.
- Bubble Tea as the main event loop and application model.
- Lip Gloss for layout and styling.[cite:1]
- Bubbles for common TUI primitives.[cite:6]
- External player backend such as `mpv`, `ffplay`, or `vlc` controlled through subprocesses.
- Local persistence via JSON, BoltDB, or SQLite.

### Module layout

```text
/cmd/radiodrift
/internal/app
/internal/ui
/internal/ui/components
/internal/ui/views
/internal/domain
/internal/store
/internal/radio
/internal/player
/internal/search
/internal/config
/internal/keymap
/internal/theme
/internal/session
```

### Recommended package roles

- `app`: top-level Bubble Tea model, routing, update loop.
- `ui/components`: reusable widgets and pane renderers.
- `ui/views`: home, discover, favorites, presets, station detail.
- `domain`: station, crate, preset, history, config types.
- `radio`: remote station lookup, normalization, health checks.
- `player`: process management and playback state.
- `store`: persistence layer.
- `search`: filtering, ranking, related-station logic.
- `keymap`: global and contextual key bindings.
- `theme`: Lip Gloss styles and palette tokens.
- `session`: restore open station, last view, active filters.

## State management

### Top-level app state

```go
type AppModel struct {
    Width             int
    Height            int
    Ready             bool

    ActiveView        ViewID
    FocusedPane       PaneID
    Theme             Theme

    Query             string
    Filters           Filters
    SearchResults     []Station
    SelectedIndex     int

    CurrentStation    *Station
    PlayerState       PlayerState
    RelatedStations   []Station

    Favorites         []Station
    History           []HistoryItem
    Presets           []Preset
    Crates            []Crate

    StatusMessage     string
    ErrorMessage      string
    Loading           bool
}
```

### Update model strategy

Use a message-driven design:

- UI messages for keys, resize, focus changes.
- Async messages for search results, metadata refresh, and player events.
- Explicit submodels for search, player, and side-panel state.

Keep playback state separate from list-rendering state. The player must survive view changes.

## Data source strategy

The app needs a station directory source and a playback engine.

### Station directory requirements

The source should provide:

- Searchable station records.
- Country and language metadata.
- Tags or genres.
- Stream URLs.
- Availability or health hints.

### Data handling rules

- Normalize country, language, and tag values.
- Cache results locally for faster browsing.
- Track “last working” streams locally.
- Fall back gracefully when a stream is dead.
- Keep recent queries cached.

## Search and ranking

This is one of the highest-leverage areas.

### Basic search modes

- Name match.
- Tag match.
- Country match.
- Language match.
- Advanced filters.

### Ranking signals

- Exact name match.
- Popularity or vote count.
- Health / last successful check.
- Tag overlap with current station.
- Country or language affinity from user history.
- Lower penalty for modest bitrate if station reliability is high.

### Related station algorithm

Start simple with weighted similarity:

- +3 same country
- +3 same language
- +2 per shared tag
- +2 similar name token
- +1 same codec
- +1 similar bitrate range
- +2 if both are in favorites-adjacent clusters

Then sort by similarity score multiplied by health confidence.

## Playback subsystem

### Player goals

- Fast startup.
- Reliable stop/start.
- Auto-reconnect option.
- Clear error reporting.
- Backend abstraction.

### Backend abstraction

```go
type Player interface {
    Play(url string) tea.Cmd
    Stop() tea.Cmd
    Pause() tea.Cmd
    Resume() tea.Cmd
    SetVolume(v int) tea.Cmd
    State() PlayerState
}
```

### Playback behaviors

- Start playback on station open or explicit tune.
- Show connecting state immediately.
- Detect process exits.
- Retry with backoff if enabled.
- Emit meaningful status messages, not raw subprocess noise.

## Persistence

### Local files to persist

- Config.
- Favorites.
- Presets.
- Crates.
- History.
- Session snapshot.
- Cache of station lookups.

### Directory suggestion

```text
~/.config/radiodrift/config.json
~/.local/share/radiodrift/favorites.json
~/.local/share/radiodrift/history.json
~/.local/share/radiodrift/crates.json
~/.cache/radiodrift/search-cache.json
```

If stronger querying is needed later, move history and station cache into SQLite.

## Screen blueprint

### Home

Purpose: offer fast entry points.

Sections:
- Continue listening.
- Presets.
- Random picks.
- Explore by country.
- Explore by tag.
- Recent discoveries.

### Discover

Purpose: browsing without typing.

Sections:
- Country clusters.
- Language groups.
- Popular tags.
- Curated prompts.
- Random station button.

### Search

Purpose: fast lookup.

Layout:
- Search bar on top.
- Active filters under it.
- Result list center.
- Station detail or now playing on the right.

### Station detail

Purpose: enrich the currently selected item.

Fields:
- Station name.
- Stream URL summary.
- Country / language / tags.
- Codec / bitrate.
- Votes / popularity.
- Health / last check.
- Related stations.
- Actions.

### Now playing

Purpose: idle-state home.

Fields:
- Large station title.
- Live status.
- Elapsed listening time.
- Small metadata badges.
- Related suggestions.
- History trail.
- Keyboard hints.

### Favorites

Purpose: reliable quick access.

Features:
- Sort by last played, alphabetical, country, tag.
- Quick presets.
- Remove or annotate.

### Crates

Purpose: curation.

Features:
- Create crate.
- Rename crate.
- Add/remove stations.
- Drag-like reorder through keys.
- Notes or descriptions.

## Keymap design

Create one central keymap package with contextual help text.

### Principles

- Vim-like movement where possible.
- Common actions remain stable across views.
- No overloaded key should surprise the user.
- Every screen should expose the 4 to 8 relevant actions in a footer.

### Suggested footer structure

```text
/ Search   Enter Tune   F Favorite   C Crate   R Related   ? Help   Q Back
```

## Theming system

### Theme tokens

Create semantic style tokens instead of ad hoc styles:

- Base background.
- Pane background.
- Active pane border.
- Primary text.
- Muted text.
- Accent.
- Success.
- Warning.
- Error.
- Badge background.
- Selected row background.
- Status line.

### Lip Gloss style groups

- `AppFrameStyle`
- `PaneStyle`
- `FocusedPaneStyle`
- `TitleStyle`
- `SubtleStyle`
- `BadgeStyle`
- `SelectedRowStyle`
- `NowPlayingTitleStyle`
- `StatusBarStyle`
- `ErrorStyle`

## Accessibility and resilience

Even in a TUI, accessibility matters.

- Ensure strong contrast in both dark and light themes.
- Never rely only on color to signal status.
- Provide compact and spacious modes.
- Make all primary features keyboard accessible.
- Keep text truncation graceful.
- Support very small terminal sizes with a fallback layout.

## Observability and debugging

Add a debug mode early.

### Debug panel or log file should show

- Current view.
- Selected station ID.
- Player backend and state.
- Last command sent to player.
- Last stream error.
- Cache hits or misses.
- Active filters.

This will save enormous time during playback and state bugs.

## Performance targets

- Initial launch under 300 ms with warm cache.
- Search interaction should feel instant on cached data.
- Result list scrolling must remain smooth on large datasets.
- Rendering should degrade gracefully on terminal resize.
- Playback state updates must not block UI rendering.

## Error handling

### Failure classes

- No network.
- Empty search results.
- Dead stream.
- Player backend missing.
- Metadata unavailable.
- Cache corrupted.

### UX responses

- Show friendly actionable messages.
- Offer fallback actions such as retry, choose another player backend, or open related stations.
- Never crash to raw stack traces in normal mode.

## Testing strategy

### Unit tests

- Search filters.
- Ranking logic.
- Related-station scoring.
- Persistence round trips.
- Keymap routing.

### Integration tests

- Player subprocess lifecycle.
- Session restore.
- Cache loading.
- Search-to-play flow.

### Manual test matrix

- macOS terminal.
- Linux terminal.
- Small terminal window.
- Missing `mpv`.
- Dead stream.
- Slow network.
- Large favorites library.

## Incremental build plan

### Phase 1: shell and navigation

Build:
- Bubble Tea root model.
- Basic view routing.
- Static mock data.
- Pane layout.
- Footer help.

Exit criteria:
- App feels structurally real before any network or playback integration.

### Phase 2: station search

Build:
- Search input.
- Result list.
- Filters.
- Station detail panel.
- Async fetch and loading states.

Exit criteria:
- User can find stations and inspect them smoothly.

### Phase 3: playback

Build:
- Player interface.
- `mpv` backend.
- Start/stop/reconnect.
- Status line.

Exit criteria:
- User can tune and keep listening reliably.

### Phase 4: favorites, history, presets

Build:
- Persistence.
- Favorite flow.
- History trail.
- Presets.

Exit criteria:
- App becomes personally useful across sessions.

### Phase 5: crates and related stations

Build:
- Crate model.
- Add/remove flows.
- Similar-station engine.
- Discovery loop improvements.

Exit criteria:
- App becomes differentiated, not just functional.

### Phase 6: polish

Build:
- Themes.
- Better empty states.
- Compact mode.
- Command palette.
- Better help screen.
- Session restore.

Exit criteria:
- App feels premium.

## MVP backlog

### Must-have

- App shell.
- Search.
- Tune station.
- Stop station.
- Favorites.
- History.
- Presets.
- Error states.
- Config file.

### Should-have

- Related stations.
- Crates.
- Session restore.
- Compact mode.
- Command palette.

### Nice-to-have

- Notes.
- Sleep timer.
- Station-of-the-day.
- Local recap view.
- Export/import.

## Suggested repo structure

```text
radiodrift/
├── cmd/
│   └── radiodrift/
│       └── main.go
├── internal/
│   ├── app/
│   ├── config/
│   ├── domain/
│   ├── keymap/
│   ├── player/
│   ├── radio/
│   ├── search/
│   ├── session/
│   ├── store/
│   ├── theme/
│   └── ui/
│       ├── components/
│       └── views/
├── testdata/
├── Makefile
├── go.mod
└── README.md
```

## Milestone definition

### Milestone A: playable prototype

A user can search, tune, favorite, and reopen the app with saved favorites.

### Milestone B: discovery product

A user can genuinely explore by related stations, country, language, and crates.

### Milestone C: keeper app

A user prefers leaving this open over opening a browser tab for internet radio.

## Product risks

- Radio streams can be flaky.
- Metadata quality can be inconsistent.
- Terminal audio control is backend-dependent.
- Overbuilding social features too early could dilute the core product.

## Risk mitigations

- Treat playback health as a first-class concept.
- Cache aggressively.
- Build on a stable external player instead of writing playback from scratch.
- Nail the solo listening loop before adding community features.

## What “great” looks like

The app should feel instantly useful in the first minute, memorable in the first session, and habit-forming by the third session.

A great version has these qualities:

- Search is effortless.
- Playback is reliable.
- The interface feels handcrafted.
- Discovery leads naturally to curation.
- Favorites and crates slowly become a personal map of taste.

## Recommended first sprint

Build this exact vertical slice first:

1. Bubble Tea app shell.
2. Search box.
3. Station result list.
4. Right-side detail pane.
5. `mpv` playback backend.
6. Favorite toggle.
7. Local favorites persistence.

If that slice feels good, the product is real. Everything else is expansion and refinement.

## Build checklist

- [ ] App shell renders correctly across common terminal sizes.
- [ ] Search returns and displays station results cleanly.
- [ ] User can tune and stop streams reliably.
- [ ] Current station persists visibly in now playing state.
- [ ] Favorites are saved and restored.
- [ ] Errors are friendly and actionable.
- [ ] Keymap feels consistent.
- [ ] Theme has a coherent identity.
- [ ] Idle state is pleasant enough to leave open.
- [ ] Related station loop feels genuinely interesting.

## Final recommendation

Use the internal codename **radiodrift** during development. It gives the project a distinct identity, matches the exploratory mood, and separates the product concept from the more generic descriptive phrase “internet radio crate digger.”
