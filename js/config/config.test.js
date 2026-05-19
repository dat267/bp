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

  test('should auto-generate config if missing', () => {
    const tempPath = path.join(process.cwd(), 'config.json');
    if (fs.existsSync(tempPath)) fs.unlinkSync(tempPath);
    
    new Config();
    assert.ok(fs.existsSync(tempPath), 'Config file should have been generated');
    fs.unlinkSync(tempPath);
  });

  test('should save changes correctly', () => {
    const tempPath = path.join(process.cwd(), 'save_test.json');
    const config = new Config(tempPath);
    config.port = '7777';
    config.save(tempPath);

    const saved = JSON.parse(fs.readFileSync(tempPath, 'utf8'));
    assert.strictEqual(saved.port, '7777');
    fs.unlinkSync(tempPath);
  });

  test('should validate port correctly', () => {
    assert.strictEqual(Config.validatePort('8080'), null);
    assert.ok(Config.validatePort('abc'));
    assert.ok(Config.validatePort('70000'));
  });

  test('should validate not empty correctly', () => {
    assert.strictEqual(Config.validateNotEmpty('val'), null);
    assert.ok(Config.validateNotEmpty(''));
    assert.ok(Config.validateNotEmpty('  '));
  });

  test('should load environment variables with precedence', () => {
    process.env.BP_PORT = '9999';
    const config = new Config();
    assert.strictEqual(config.port, '9999');
    delete process.env.BP_PORT;
  });
});
