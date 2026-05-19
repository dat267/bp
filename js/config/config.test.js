const { test, describe } = require('node:test');
const assert = require('node:assert');

describe('Config', () => {
  test('should load environment variables', () => {
    process.env.PORT = '9999';
    const config = require('./config');
    assert.strictEqual(config.port, '9999');
    delete process.env.PORT;
  });
});
