/**
 * Minimal argument parser for CLI flags
 * Converts ['--name=World', '--port=8080'] into { name: 'World', port: '8080' }
 */
function parseArgs(args) {
  const options = {};
  for (const arg of args) {
    if (arg.startsWith('--')) {
      const [key, value] = arg.slice(2).split('=');
      options[key] = value || true;
    }
  }
  return options;
}

module.exports = { parseArgs };
