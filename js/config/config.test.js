const { test, describe, before, after } = require('node:test');
const assert = require('node:assert');
const fs = require('node:fs');
const path = require('node:path');
const Config = require('./config');

describe('Config', () => {
  const configPath = path.join(process.cwd(), 'config.json');

  test('should load from config.json', () => {
    const configData = { port: '1234', appEnv: 'production' };
    fs.writeFileSync(configPath, JSON.stringify(configData));
    
    const config = new Config();
    
    assert.strictEqual(config.port, '1234');
    assert.strictEqual(config.appEnv, 'production');
    
    fs.unlinkSync(configPath);
  });

  test('should load from a custom path', () => {
    const customPath = path.join(process.cwd(), 'custom.json');
    const configData = { port: '5555' };
    fs.writeFileSync(customPath, JSON.stringify(configData));
    
    const config = new Config(customPath);
    
    assert.strictEqual(config.port, '5555');
    
    fs.unlinkSync(customPath);
  });

  test('should load environment variables with precedence', () => {
    process.env.BP_PORT = '9999';
    const config = new Config();
    assert.strictEqual(config.port, '9999');
    delete process.env.BP_PORT;
  });
});
