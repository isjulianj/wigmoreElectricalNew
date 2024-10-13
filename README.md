# Mutuo-go
***
My go project starter template

## Getting started
***

### Prerequisites
The Project requires [Go](https://go.dev/) version [1.21](https://go.dev/doc/devel/release#go1.21.0) or above.

#### UI templating
The Project requires the latest version of [templ](https://templ.guide/quick-start/installation) 



We're also using tailwind + DasiyUI, conventionally they both require node to work but tailwind releases a standalone executable. 
Using the DaisyUI library also requires a little work because the tailwind executable does not support 3rd party plugins. 

lukily [dobicinaitis](https://github.com/dobicinaitis/tailwind-cli-extra/commits?author=dobicinaitis) has created a tailwind-extra executable release that has DaisyUI baked in.

To get this working:

Download the latest executable from
the [releases](https://github.com/dobicinaitis/tailwind-cli-extra/releases/latest) page.

```bash
# Mac arm64 example
curl -sLO https://github.com/dobicinaitis/tailwind-cli-extra/releases/latest/download/tailwindcss-extra-macos-arm64 -o /tailwindcss-extra
```

Make the file executable:

```bash
chmod +x /tailwindcss-extra
```

Following these steps will allow the make file to pick up the executable and build your css.
