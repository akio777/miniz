module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: ['plugin:vue/essential', '@vue/airbnb'],
  parserOptions: {
    parser: 'babel-eslint',
  },
  rules: {
    'no-console':
        process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger':
        process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'max-len': 'off',
    // 'linebreak-style': ['error', 'windows'],
    'comma-dangle': 'off',
    quotes: ['error', 'single'],
    'no-unused-vars': 'error'
  },
};
