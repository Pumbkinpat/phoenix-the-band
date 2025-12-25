module.exports = {
  content: [
    "./templates/**/*.templ",
    "./templates/**/*.go",
    "./**/*.{html, js, mjs, ejs}",
    "./web/public/icon/**/*.svg",
  ],
  theme: {
    extend: {
      content: {
        // Use double backslashes in JS strings to escape correctly
        "chevron-rb": '"\\uE970"',
      },
    }
  },
  plugins: [require('./daisyUI/daisyui')], // make sure to match the path of daisyui mjs files
}

// Tailwind v4 don't need configuration file 
// but the vs code Tailwind CSS IntelliSense 
// need it to function normally, DAMN IT!
