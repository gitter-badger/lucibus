import 'must'
import webdriverio from 'webdriverio'
import booleanFromEnv from '../caido/config/booleanFromEnv'

const DELETE = '\uE003'

var browserConfig

if (booleanFromEnv('CI', false)) {
  console.log('Using CI')
  browserConfig = {
    user: process.env.SAUCE_USERNAME,
    key: process.env.SAUCE_ACCESS_KEY,
    desiredCapabilities: {
      browserName: 'chrome',
      'tunnel-identifier': process.env.TRAVIS_JOB_NUMBER,
      'idle-timeout': 900,
      name: 'lucibus',
      build: process.env.TRAVIS_BUILD_NUMBER
    }
  }
} else {
  browserConfig = {
    desiredCapabilities: {
      browserName: 'chrome'
    }
  }
}

var browsers = webdriverio.multiremote({
  first: browserConfig,
  second: browserConfig
})

var firstBrowser = browsers.select('first')
var secondBrowser = browsers.select('second')

describe('App', function () {
  before(function *() {
    yield browsers.init()
  })
  after(function *() {
    yield browsers.end()
  })
  beforeEach('go to URL', function *() {
    yield browsers.url('http://localhost:8081/')
  })

  describe('Grandmaster level', function () {
    function * mustEqual (b, value) {
      yield b.getValue('#grandmaster .level input').must.eventually.equal(value.toString())
    }
    it('must start at 100', function *() {
      yield* mustEqual(firstBrowser, 100)
      yield* mustEqual(secondBrowser, 100)
    })
    function * click (b) {
      yield b.waitForVisible('#grandmaster .level input')
      yield b.click('#grandmaster .level input')
    }
    function * clickAndType (b, value) {
      yield* click(b)
      yield b.keys(DELETE + DELETE + DELETE + value.toString())
    }
    it('changing on one should change both', function *() {
      yield* clickAndType(firstBrowser, 10)
      yield* mustEqual(firstBrowser, 10)
      yield* mustEqual(secondBrowser, 10)
      yield* clickAndType(secondBrowser, 20)
      yield* mustEqual(secondBrowser, 20)
      yield* mustEqual(firstBrowser, 20)
    })
  })
})
