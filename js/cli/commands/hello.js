const { parseArgs } = require('../args');

class HelloCommand {
  constructor(config) {
    this.name = 'hello';
    this.description = 'Greets a person';
    this.config = config;
  }

  execute(args) {
    const options = parseArgs(args);
    
    let defaultName = 'World';
    if (this.config.appEnv === 'production') {
      defaultName = 'Production User';
    }

    const name = options.name || defaultName;
    console.log(`Hello, ${name}!`);
  }
}

module.exports = HelloCommand;
