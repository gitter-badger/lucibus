var booleanFromEnv = require('./config/booleanFromEnv')

const SAUCELABS = booleanFromEnv('CI', false)

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
    platform: 'Windows 7',
    version: '35'
  },
  sl_firefox: {
    base: 'SauceLabs',
    browserName: 'firefox',
    version: '30'
  },
  sl_ios_safari: {
    base: 'SauceLabs',
    browserName: 'iphone',
    platform: 'OS X 10.9',
    version: '7.1'
  },
  sl_ie_11: {
    base: 'SauceLabs',
    browserName: 'internet explorer',
    platform: 'Windows 8.1',
    version: '11'
  }
}

module.exports = function (config) {
  config.set({
    browsers: SAUCELABS ? Object.keys(customLaunchers) : ['Chrome'],
    frameworks: ['mocha'],
    files: [
      'test/unit/**/*.js'
    ],
    customLaunchers: SAUCELABS ? customLaunchers : {},
    preprocessors: {
      'test/unit/**/*.js': ['webpack']
    },
    reporters: [
      'dots',
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
    logLevel: config.LOG_DEBUG,

    sauceLabs: {
      testName: 'lucibus/caido:unit',
      tunnelIdentifier: process.env.TRAVIS_JOB_NUMBER,
      username: process.env.SAUCE_USERNAME,
      accessKey: process.env.SAUCE_ACCESS_KEY,
      startConnect: false
    }
  })
}
