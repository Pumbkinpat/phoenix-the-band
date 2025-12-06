## How to run 

- Install latest tailwindcss standalone cli's executable file that suit your os
```shell
  curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download tailwindcss-macos-arm64
  chmod +x tailwindcss-macos-arm64
  mv tailwindcss-macos-arm64 tailwindcss
``` 
Check tailwind document for [tailwindcss standalone cli](https://tailwindcss.com/blog/standalone-cli)

- Gernerate css and go files `make watch -j2`
- Build and run `make start`