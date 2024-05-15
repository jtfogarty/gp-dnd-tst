# README

## About

I created this wails app on gitpod using the below command;
```./sk.sh pnpm newapp brand```
This shells script is modified from here
https://wails.io/docs/guides/sveltekit/#sveltekit-wailssh to work on gitpod.

The script creates a temp directory, does it's magic and moves the files (Thanks ChatGPT 4o)

The commands I ran after the above are as follows;

```
cd frontend/
pnpm dlx svelte-add@latest tailwindcs
pnpm install
pnpm dlx shadcn-svelte@latest init
pnpm dlx shadcn-svelte@latest add alert-dialog
wails dev
```