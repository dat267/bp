const readline = require('readline');

const Config = require('../../config/config');

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

    const showMenu = () => {
      console.log('\n--- rclone-style Configuration Menu ---');
      console.log('v) View current configuration');
      Object.keys(Config.schema).forEach((key, i) => {
        console.log(`${i + 1}) Edit ${Config.schema[key].label}`);
      });
      console.log('s) Save and Exit');
      console.log('q) Quit without saving');
    };

    while (true) {
      showMenu();
      const choice = await new Promise(resolve => rl.question('Choose option: ', resolve));
      
      const schemaKeys = Object.keys(Config.schema);

      if (choice === 'v') {
        console.log('\nCurrent Configuration:');
        schemaKeys.forEach(key => {
          console.log(`  ${(Config.schema[key].label + ':').padEnd(16)} ${this.config[key]}`);
        });
        continue;
      }
      if (choice === 's') {
        this.config.save();
        console.log('Configuration saved.');
        rl.close();
        return;
      }
      if (choice === 'q') {
        console.log('Exiting without saving.');
        rl.close();
        return;
      }

      const idx = parseInt(choice, 10);
      if (!isNaN(idx) && idx > 0 && idx <= schemaKeys.length) {
        const key = schemaKeys[idx - 1];
        const field = Config.schema[key];
        this.config[key] = await prompt(field.label, this.config[key], field.validator);
      } else {
        console.log('Invalid option.');
      }
    }
  }
}

module.exports = ConfigCommand;
