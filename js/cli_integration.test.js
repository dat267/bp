const { test, describe } = require('node:test');
const assert = require('node:assert');
const { spawnSync } = require('node:child_process');
const path = require('node:path');

describe('CLI Integration', () => {
  const cliPath = path.join(__dirname, 'cli', 'cli.js');

  const runCLI = (args) => {
    const result = spawnSync('node', [cliPath, ...args], { encoding: 'utf8' });
    return result.stdout;
  };

  test('should handle global flag before subcommand', () => {
    const output = runCLI(['--verbose', 'info']);
    assert.ok(output.includes('Verbose:     true'));
  });

  test('should handle global flag after subcommand', () => {
    const output = runCLI(['info', '--verbose']);
    assert.ok(output.includes('Verbose:     true'));
  });

  test('should pass subcommand-specific flags correctly', () => {
    const output = runCLI(['--verbose', 'hello', '--name=Tester']);
    assert.ok(output.includes('Hello, Tester!'));
  });
});
