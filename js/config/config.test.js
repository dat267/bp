const { test, describe, before, after } = require('node:test');
const assert = require('node:assert');
const fs = require('node:fs');
const path = require('node:path');

describe('Config', () => {
  const configPath = path.join(process.cwd(), 'config.json');

  test('should load from config.json', () => {
    const configData = { port: '1234', appEnv: 'production' };
    fs.writeFileSync(configPath, JSON.stringify(configData));
    
    // Clear require cache to force re-instantiation
    delete require.cache[require.resolve('./config')];
    const config = require('./config');
    
    assert.strictEqual(config.port, '1234');
    assert.strictEqual(config.appEnv, 'production');
    
    fs.unlinkSync(configPath);
  });

  test('should load environment variables with precedence', () => {
    process.env.PORT = '9999';
    delete require.cache[require.resolve('./config')];
    const config = require('./config');
    assert.strictEqual(config.port, '9999');
    delete process.env.PORT;
  });
});
