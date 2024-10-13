/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './ui/**/*.templ',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    require('daisyui'),
  ],
}

