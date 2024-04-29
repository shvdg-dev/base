/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/**/**/**/**/**/*.go",
    "./views/**/**/**/**/**/*.go",
    "./views/**/**/**/**/*.go",
    "./views/**/**/**/*.go",
    "./views/**/**/*.go",
    "./views/**/*.go"
  ],
  theme: {
    extend: {
      colors: {
        'base-input': '#787878',
      }
    },
  },
  daisyui: {
    themes: [
      {
        'custom-theme': {
          'primary': '#fcbbe8',
          'primary-focus': '#4506cb',
          'primary-content': '#ffffff',
          'secondary': '#f000b8',
          'secondary-focus': '#bd0091',
          'secondary-content': '#ffffff',
          'accent': '#37cdbe',
          'accent-focus': '#2aa79b',
          'accent-content': '#ffffff',
          'neutral': '#3d4451',
          'neutral-focus': '#2a2e37',
          'neutral-content': '#ffffff',
          'base-100': '#797979',
          'base-200': '#474747',
          'base-300': '#2e2e2e',
          'base-content': '#dbdbdb',
          'info': '#2094f3',
          'success': '#009485',
          'warning': '#ff9900',
          'error': '#ff5724',
        },
      },
    ],
  },
  plugins: [require("daisyui")],
}