# PersonalSite

## Go module exploration

Use the `go_package_api` and `go_search` gopls MCP tools to explore external/dependency packages. Do not use the Read tool to open files under `~/go/pkg/mod/`.

## Styling conventions

### Font sizing
Always use Tailwind's named text size classes (`text-xs`, `text-sm`, `text-base`, `text-lg`, `text-xl`, `text-2xl`, etc.) for font sizes. Do not use arbitrary pixel or rem values (e.g. `text-[11px]`, `text-[1.3rem]`).

Current type scale on the home page:
- Name / page title: `text-3xl`
- Intro paragraphs (hero): `text-xl`
- Body / section prose: `text-base`
- Small supporting text (section labels, tags, footer): `text-xs`
