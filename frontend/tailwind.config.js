/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        'primary': '#14568E',
        'orange': {
          500: '#fe6a00',
          600: '#da4e00'
        },
        'highlight': '#253f8d',
      }
    },
  },
  plugins: [],
}

