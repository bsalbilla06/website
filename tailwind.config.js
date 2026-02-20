/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/templates/**/*.html"],
  theme: {
    extend: {
      fontFamily: {
        'handwritten': ['"Architects Daughter"', 'cursive'],
      }
    },
  },
  plugins: [],
}
