const fs = require('fs');
const path = require('path');

const config = require('../config/config');

class CLI {
  constructor() {
    this.commands = new Map();
    this.loadCommands();
  }

  loadCommands() {
    const commandsPath = path.join(__dirname, 'commands');
    const files = fs.readdirSync(commandsPath);

    for (const file of files) {
      if (file.endsWith('.js')) {
        const CommandClass = require(path.join(commandsPath, file));
        const cmd = new CommandClass(config);
        this.commands.set(cmd.name, cmd);
      }
    }
  }

  run(args) {
    const [commandName, ...rest] = args;

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
