var booleanFromEnv = require('./config/booleanFromEnv')

const SAUCELABS = booleanFromEnv('CONTINUOUS_INTEGRATION', false)

process.env.WEBPACK_MINIMIZE = 'false'
process.env.WEBPACK_DEVTOOL = 'inline-source-map'
process.env.WEBPACK_INLINE = 'true'

var webpackConf = require('./webpack.config.js')
webpackConf.node = {fs: 'empty'}
webpackConf.module.postLoaders = [{
  test: /\.js(x)?$/,
  include: /app\//,
  loader: 'istanbul-instrumenter'
}]

const customLaunchers = {
  sl_chrome: {
    base: 'SauceLabs',
    browserName: 'chrome',
    version: 'latest'
  }
}

module.exports = function (config) {
  config.set({
    browsers: SAUCELABS ? Object.keys(customLaunchers) : ['Chrome'],
    frameworks: ['source-map-support', 'mocha'],
    files: [
      'test/unit/**/*.js'
    ],
    customLaunchers: SAUCELABS ? customLaunchers : {},
    preprocessors: {
      'test/unit/**/*.js': ['webpack']
    },
    startConnect: false,
    reporters: [
      'mocha',
      'coverage'
    ].concat(SAUCELABS ? ['saucelabs'] : []),
    webpack: webpackConf,
    webpackServer: {
      noInfo: true
    },
    coverageReporter: {
      type: 'text',
      dir: 'coverage/'
    },
    client: {
      mocha: {
        reporter: 'html', // change Karma's debug.html to the mocha web reporter
        ui: 'bdd'
      }
    }
  })
}
