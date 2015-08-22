import 'must'
import webdriverio from 'webdriverio'

const DELETE = '\uE003'
// const TAB = '\uE004'
// const RETURN = '\uE006'
// const DOWN_ARROW = '\uE015'

var browsers = webdriverio.multiremote({
  first: {
      desiredCapabilities: {
        browserName: 'chrome'
      }
    },
  second: {
      desiredCapabilities: {
        browserName: 'chrome'
      }
    }
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
    yield browsers.url('http://localhost/')
  })


  describe('Live', function () {
    describe('Grandmaster', function () {
      describe('Level', function () {
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
  })
})
