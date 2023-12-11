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
    console.log(
      "Salut !! Vous êtes sur le rice cooker. Le rice cooker est allumé."
    );
  }

  turnOff() {
    this.power = false;
    console.log("Le rice cooker est éteint.");
  }

  chooseFunction(userInput) {
    switch (userInput) {
      case "1":
        this.turnOn();
        try {
          const action = this.getValidNumber(
            "Que voulez-vous faire ? (1 pour cuire / 2 pour bouillir): ",
            [1, 2]
          );
          switch (action) {
            case 1:
              const duration = this.getValidNumber(
                "Combien de minutes souhaitez-vous cuire ? "
              );
              const temperature = this.getValidNumber(
                "À quelle température souhaitez-vous cuire ? "
              );
              this.cook(duration, temperature);
              break;

            case 2:
              this.boil();
              break;

            default:
              console.log("Action non reconnue.");
          }
        } catch (error) {
          console.error("Erreur:", error.message);
        }
        break;

      case "2":
        this.turnOff();
        break;

      default:
        console.log("Commande non reconnue.");
    }
  }

  getValidNumber(promptMessage, validChoices = []) {
    let userInput;
    while (true) {
      try {
        userInput = parseInt(prompt(promptMessage), 10);
        if (
          isNaN(userInput) ||
          (validChoices.length > 0 && !validChoices.includes(userInput))
        ) {
          throw new Error(
            "Veuillez saisir un nombre valide parmi les choix spécifiés."
          );
        }
        break;
      } catch (error) {
        console.error("Erreur:", error.message);
      }
    }
    return userInput;
  }

  cook(duration, temperature) {
    try {
      console.log(
        `Cuisson en cours pendant ${duration} minute(s) à ${temperature}°C...`
      );
      this.sleep(duration);
      console.log("Cuisson terminée!");
    } catch (error) {
      console.error("Erreur:", error.message);
    }
  }

  boil() {
    try {
      console.log("Bouillir de l'eau en cours...");
      const boilingTime = 10; // Durée de la simulation (10 minutes par exemple)
      let currentTemperature = 0;

      for (let i = 1; i <= boilingTime; i++) {
        currentTemperature += 10;
        this.sleep(1);
        console.log(
          `Température de l'eau: ${currentTemperature}°C - Durée restante: ${
            boilingTime - i
          } minutes`
        );
      }

      console.log("L'eau a atteint 100°C. Arrêt automatique.");
    } catch (error) {
      console.error("Erreur:", error.message);
    }
  }

  sleep(minutes) {
    const milliseconds = minutes * 60 * 1000;
    const startTime = new Date().getTime();
    while (new Date().getTime() < startTime + milliseconds);
  }
}

// Exemple d'utilisation
const myRiceCooker = new RiceCooker();

try {
  const userDecision = myRiceCooker.getValidNumber(
    "Salut !! Vous êtes sur le rice cooker. Voulez-vous l'allumer ? (1 pour oui / 2 pour non): ",
    [1, 2]
  );
  myRiceCooker.chooseFunction(userDecision.toString());
} catch (error) {
  console.error("Erreur:", error.message);
}
