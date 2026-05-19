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

    const prompt = (query, current, validator) => new Promise((resolve) => {
      const ask = () => {
        rl.question(`${query} [${current}]: `, (answer) => {
          const value = answer || current;
          if (validator) {
            const error = validator(value);
            if (error) {
              console.log(`Error: ${error}`);
              return ask();
            }
          }
          resolve(value);
        });
      };
      ask();
    });

    const menu = () => new Promise((resolve) => {
      console.log('\n--- rclone-style Configuration Menu ---');
      console.log('1) View current configuration');
      console.log('2) Edit App Environment');
      console.log('3) Edit Port');
      console.log('4) Edit API Key');
      console.log('s) Save and Exit');
      console.log('q) Quit without saving');
      rl.question('Choose option: ', resolve);
    });

    const validatePort = (val) => {
      const port = parseInt(val, 10);
      if (isNaN(port) || port < 1 || port > 65535) return 'Port must be between 1 and 65535';
      return null;
    };

    const validateNotEmpty = (val) => {
      if (!val || val.trim() === '') return 'Value cannot be empty';
      return null;
    };

    while (true) {
      const choice = await menu();
      switch (choice) {
        case '1':
          console.log('\nCurrent Configuration:');
          console.log(`  AppEnv:  ${this.config.appEnv}`);
          console.log(`  Port:    ${this.config.port}`);
          console.log(`  APIKey:  ${this.config.apiKey}`);
          break;
        case '2':
          this.config.appEnv = await prompt('App Environment', this.config.appEnv);
          break;
        case '3':
          this.config.port = await prompt('Port', this.config.port, validatePort);
          break;
        case '4':
          this.config.apiKey = await prompt('API Key', this.config.apiKey, validateNotEmpty);
          break;
        case 's':
          this.config.save();
          console.log('Configuration saved.');
          rl.close();
          return;
        case 'q':
          console.log('Exiting without saving.');
          rl.close();
          return;
        default:
          console.log('Invalid option.');
      }
    }
  }
}

module.exports = ConfigCommand;
