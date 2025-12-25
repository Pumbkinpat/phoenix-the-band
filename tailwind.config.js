module.exports = {
  content: [
    "./templates/**/*.templ",
    "./templates/**/*.go",
    "./**/*.{html, js, mjs, ejs}"
  ],
  theme: {
    extend: {
      
    }
  },
  plugins: [require('./daisyUI/daisyui')], // make sure to match the path of daisyui mjs files
}

// Tailwind v4 don't need configuration file 
// but the vs code Tailwind CSS IntelliSense 
// need it to function normally, DAMN IT!
