/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')
module.exports = {
  darkMode: 'class',
  content: ["./template/**/*.{html,js,templ,go}"],  
  theme: {
    extend: {
        colors:{
          ...colors,
        }
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
