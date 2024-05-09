const { defineConfig } = require('@pandacss/dev');

module.exports = defineConfig({
  plugins: [require('@pandacss/dev/postcss')()],
});
