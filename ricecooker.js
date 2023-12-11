const prompt = require("prompt-sync")();

class RiceCooker {
  constructor() {
    this.power = false;
  }

  checkPower() {
    return this.power;
  }

  turnOn() {
    this.power = true;
    console.log("Hello! You are on the rice cooker. The rice cooker is turned on.");
  }

  turnOff() {
    this.power = false;
    console.log("The rice cooker is turned off.");
  }

  chooseFunction(userInput) {
    switch (userInput) {
      case "1":
        this.turnOn();
        try {
          const action = this.getValidNumber("What do you want to do? (1 to cook / 2 to boil): ", [1, 2]);
          switch (action) {
            case 1:
              const duration = this.getValidNumber("How many minutes do you want to cook? ");
              const temperature = this.getValidNumber("At what temperature do you want to cook? ");
              this.cook(duration, temperature);
              break;

            case 2:
              this.boil();
              break;

            default:
              console.log("Unrecognized action.");
          }
        } catch (error) {
          console.error("Error:", error.message);
        }
        break;

      case "2":
        this.turnOff();
        break;

      default:
        console.log("Unrecognized command.");
    }
  }

  getValidNumber(promptMessage, validChoices = []) {
    let userInput;
    while (true) {
      try {
        userInput = parseInt(prompt(promptMessage), 10);
        if (isNaN(userInput) || (validChoices.length > 0 && !validChoices.includes(userInput))) {
          throw new Error("Please enter a valid number among the specified choices.");
        }
        break;
      } catch (error) {
        console.error("Error:", error.message);
      }
    }
    return userInput;
  }

  cook(duration, temperature) {
    try {
      console.log(`Cooking in progress for ${duration} minute(s) at ${temperature}°C...`);
      this.sleep(duration);
      console.log("Cooking completed!");
    } catch (error) {
      console.error("Error:", error.message);
    }
  }

  boil() {
    try {
      console.log("Boiling water in progress...");
      const boilingTime = 10; // Simulation duration (10 minutes, for example)
      let currentTemperature = 0;

      for (let i = 1; i <= boilingTime; i++) {
        currentTemperature += 10;
        this.sleep(1);
        console.log(`Water temperature: ${currentTemperature}°C - Remaining time: ${boilingTime - i} minutes`);
      }

      console.log("Water has reached 100°C. Automatic shutdown.");
    } catch (error) {
      console.error("Error:", error.message);
    }
  }

  sleep(minutes) {
    const milliseconds = minutes * 60 * 1000;
    const startTime = new Date().getTime();
    while (new Date().getTime() < startTime + milliseconds);
  }
}

// Example usage
const myRiceCooker = new RiceCooker();

try {
  const userDecision = myRiceCooker.getValidNumber(
    "Hello! You are on the rice cooker. Do you want to turn it on? (1 for yes / 2 for no): ",
    [1, 2]
  );
  myRiceCooker.chooseFunction(userDecision.toString());
} catch (error) {
  console.error("Error:", error.message);
}
