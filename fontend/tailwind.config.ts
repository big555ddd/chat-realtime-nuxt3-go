import type { Config } from 'tailwindcss';

const config: Config = {
  content: [
    './components/**/*.{vue,js,ts}', // Adjust paths according to your project structure
    './pages/**/*.{vue,js,ts}',
    './layouts/**/*.{vue,js,ts}',
    './plugins/**/*.{js,ts}',
    './nuxt.config.{js,ts}',
  ],
  theme: {
    extend: {
      colors: {
        'custom-red': '#ed423a',
        "background": '#343a40',
        "cut-background" : '#1f2937',
        "gray-dark" : '#343a40',
        "success" :  '#28a745',
        "red" : '#dc3545',
        "light" : '#f8f9fa',
        
      },
    },
  },
  plugins: [],
};

export default config;
