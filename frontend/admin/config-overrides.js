const { override, addWebpackAlias } = require('customize-cra')
const path = require('path')

module.exports = override(
  addWebpackAlias({
    ['@']: path.resolve(__dirname, 'src'),
    ['_v']: path.resolve(__dirname, 'src/views'),
    ['_c']: path.resolve(__dirname, 'src/components'),
    ['_a']: path.resolve(__dirname, 'src/assets')
  })
)
