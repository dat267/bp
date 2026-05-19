const readline = require('readline');

class ConfigCommand {
  constructor(config) {
    this.name = 'config';
    this.description = 'Interactively setup configuration';
    this.config = config;
  }

  async execute() {
    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout
    });

    const prompt = (query, current) => new Promise((resolve) => {
      rl.question(`${query} [${current}]: `, (answer) => {
        resolve(answer || current);
      });
    });

    console.log('--- Interactive Configuration Setup ---');

    this.config.appEnv = await prompt('App Environment', this.config.appEnv);
    this.config.port = await prompt('Port', this.config.port);
    this.config.apiKey = await prompt('API Key (required)', this.config.apiKey);

    const confirm = await prompt('\nSave changes to config.json? (y/n)', 'n');
    if (confirm.toLowerCase() === 'y') {
      try {
        this.config.save();
        console.log('Configuration saved successfully.');
      } catch (err) {
        console.error('Failed to save config:', err.message);
      }
    } else {
      console.log('Changes discarded.');
    }

    rl.close();
  }
}

module.exports = ConfigCommand;
