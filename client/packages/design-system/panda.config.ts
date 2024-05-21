import { defineConfig } from '@pandacss/dev';

export const keyframes = {
  loadingBar: {
    '0%': {
      left: '0',
      right: '100%',
    },
    '50%': {
      left: '25%',
      right: '0',
    },
    '100%': {
      left: '100%',
      right: '0',
    },
  },
};

export default defineConfig({
  // Whether to use css reset
  preflight: true,

  // Where to look for your css declarations
  include: ['./src/**/*.{ts,tsx}'],

  // Files to exclude
  exclude: [],

  // Useful for theme customization
  theme: {
    extend: {
      keyframes,
    },
  },

  // The output directory for your css system
  outdir: 'styled-system',
});
