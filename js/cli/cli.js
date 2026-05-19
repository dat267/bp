const fs = require('fs');
const path = require('path');

const Config = require('../config/config');

class CLI {
  constructor() {
    // 1. Calculate default config path (portable: same folder as this script)
    const entryDir = __dirname;
    const defaultConfigPath = path.join(entryDir, 'config.json');

    // 2. Parse global flags
    const configArg = process.argv.find(arg => arg.startsWith('--config='));
    const configPath = configArg ? configArg.split('=')[1] : defaultConfigPath;
    
    this.config = new Config(configPath);
    this.commands = new Map();
    this.loadCommands();
  }

  loadCommands() {
    const commandsPath = path.join(__dirname, 'commands');
    const files = fs.readdirSync(commandsPath);

    for (const file of files) {
      if (file.endsWith('.js')) {
        const CommandClass = require(path.join(commandsPath, file));
        const cmd = new CommandClass(this.config);
        this.commands.set(cmd.name, cmd);
      }
    }
  }

  run(args) {
    // 1. Filter out global flags
    const verbose = args.some(arg => arg === '--verbose' || arg === '-v');
    if (verbose) {
      this.config.verbose = true;
    }

    const filteredArgs = args.filter(arg => !arg.startsWith('--config=') && arg !== '--verbose' && arg !== '-v');
    const [commandName, ...rest] = filteredArgs;

    if (!commandName || !this.commands.has(commandName)) {
      this.showHelp();
      process.exit(1);
    }

    const command = this.commands.get(commandName);
    try {
      command.execute(rest);
    } catch (error) {
      console.error(`Error: ${error.message}`);
      process.exit(1);
    }
  }

  showHelp() {
    console.log('Usage: node cli.js <command> [arguments]\n');
    console.log('Available commands:');
    for (const cmd of this.commands.values()) {
      console.log(`  ${cmd.name.padEnd(10)} ${cmd.description}`);
    }
  }
}

if (require.main === module) {
  const cli = new CLI();
  cli.run(process.argv.slice(2));
}
