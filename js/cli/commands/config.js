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
      Config.schema.forEach((field, i) => {
        console.log(`${i + 1}) Edit ${field.label}`);
      });
      console.log('s) Save and Exit');
      console.log('q) Quit without saving');
    };

    while (true) {
      showMenu();
      const choice = await new Promise(resolve => rl.question('Choose option: ', resolve));
      
      if (choice === 'v') {
        console.log('\nCurrent Configuration:');
        Config.schema.forEach(field => {
          console.log(`  ${(field.label + ':').padEnd(16)} ${this.config[field.key]}`);
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
      if (!isNaN(idx) && idx > 0 && idx <= Config.schema.length) {
        const field = Config.schema[idx - 1];
        this.config[field.key] = await prompt(field.label, this.config[field.key], field.validator);
      } else {
        console.log('Invalid option.');
      }
    }
  }
}

module.exports = ConfigCommand;
