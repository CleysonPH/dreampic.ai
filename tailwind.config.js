/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ", "./**/*.templ"],
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["dark"]
  }
}
