/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ["./template/**/*.{html,js,templ,go}"],  
  plugins: [require("@tailwindcss/forms")],
};
