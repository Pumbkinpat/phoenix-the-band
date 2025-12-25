# How to run

## Setup tools automatically with make
- Install go language itself at [download go](https://go.dev/doc/install)
- Intall make build tool
```shell
  # linux
  sudo apt update && sudo apt install make

  # macos
  xcode-select --install
```

```shell
  # verify make installation
  make --version

  # install tools
  make install

  # Development - runs all watchers in parallel 
  make dev

  # Alternative: Run separate commands in two terminal
  make dev-css
  make dev-server
```

## Install tools manually
- First of all, install go language itself at [download go](https://go.dev/doc/install)
- Install **daisyUI** including tailwindcss executable, daisyUI bundle js files, visit [daisyUI github](https://daisyui.com/docs/install/standalone)
```shell
  # linux/macos
  curl -sL daisyui.com/fast | bash

  # window
  powershell -c "irm daisyui.com/fast.ps1 | iex"
``` 

- Install **templ go** - render HTML with go code
```shell
go get -tool github.com/a-h/templ/cmd/templ@latest
```
- Install latest **gsap animation** bundled files
```shell
  # gsap library
  curl -O https://raw.githubusercontent.com/greensock/GSAP/refs/heads/master/dist/gsap.min.js

  # ScrollTrigger plugin
  curl -O https://raw.githubusercontent.com/greensock/GSAP/refs/heads/master/dist/ScrollTrigger.min.js
```
or install those files above directly at _[gsap website](https://gsap.com/docs/v3/Installation)_. Simply install the zip package, inside `minified` folder copy and paste those files into the project folder in `./web/static/js/gsap`

- Check if the gsap import line in `BaseLayout.templ` file is correct, make sure to use the correct path to those `js` files
```html
  <script src="./static/js/gsap.min.js"></script>
  <script src="./static/js/ScrollTrigger.min.js"></script>
```

- (Optional) Install **wgo** - watcher tool for golang, visit [wgo github](https://github.com/bokwoon95/wgo/releases/latest) to download suitable file for your os, the command below works for macos intel
```shell
curl -s https://api.github.com/repos/bokwoon95/wgo/releases/latest \
  | grep -o "https://[^\"]*wgo-macos[^\"]*" \
  | head -n 1 \
  | xargs curl -LO 
  && chmod +x wgo-macos
  && sudo mv wgo-macos /usr/local/bin/wgo
```

- Install **make** - build tool, depend on your os, after that run to start all watchers and run server
```shell
    # Development - runs all watchers in parallel 
    make dev

    # Alternative: Run separate commands in two terminal
    make dev-css
    make dev-server
```