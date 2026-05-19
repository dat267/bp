const { test, describe } = require('node:test');
const assert = require('node:assert');
const { withRetry } = require('./retry');

describe('Retry Utility', () => {
  test('should succeed after multiple attempts', async () => {
    let attempts = 0;
    const result = await withRetry(async () => {
      attempts++;
      if (attempts < 2) throw new Error('fail');
      return 'success';
    }, { initialDelay: 1, maxAttempts: 3 });

    assert.strictEqual(result, 'success');
    assert.strictEqual(attempts, 2);
  });

  test('should fail after max attempts', async () => {
    let attempts = 0;
    try {
      await withRetry(async () => {
        attempts++;
        throw new Error('permanent fail');
      }, { initialDelay: 1, maxAttempts: 2 });
      assert.fail('Should have thrown an error');
    } catch (error) {
      assert.strictEqual(error.message, 'permanent fail');
      assert.strictEqual(attempts, 2);
    }
  });
});
