const { test, describe } = require('node:test');
const assert = require('node:assert');
const Logger = require('./logger');

describe('Logger', () => {
  test('should respect log levels', () => {
    const logger = new Logger('warn');
    assert.strictEqual(logger.level, 2);
  });
});
