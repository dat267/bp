class HelloCommand {
  constructor() {
    this.name = 'hello';
    this.description = 'Greets a person';
  }

  execute(args) {
    const nameArg = args.find(arg => arg.startsWith('--name='));
    const name = nameArg ? nameArg.split('=')[1] : 'World';
    console.log(`Hello, ${name}!`);
  }
}

module.exports = HelloCommand;
