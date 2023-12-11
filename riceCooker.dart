import 'dart:io';

class RiceCooker {
  bool power = false;

  bool checkPower() {
    return power;
  }

  void turnOn() {
    power = true;
    print("Salut !! Vous êtes sur le rice cooker. Le rice cooker est allumé.");
  }

  void turnOff() {
    power = false;
    print("Le rice cooker est éteint.");
  }

  void chooseFunction(String userInput) {
    switch (userInput) {
      case '1':
        turnOn();
        try {
          int action = getValidNumber(
              'Que voulez-vous faire ? (1 pour cuire / 2 pour bouillir): ', [1, 2]);
          switch (action) {
            case 1:
              int duration = getValidNumber('Combien de minutes souhaitez-vous cuire ? ', null);
              int temperature = getValidNumber('À quelle température souhaitez-vous cuire ? ', null);
              cook(duration, temperature);
              break;
            case 2:
              boil();
              break;
            default:
              print('Action non reconnue.');
          }
        } catch (error) {
          print('Erreur: $error');
        }
        break;
      case '2':
        turnOff();
        break;
      default:
        print('Commande non reconnue.');
    }
  }

  int getValidNumber(String promptMessage, List<int>? validChoices) {
    while (true) {
      stdout.write(promptMessage);
      String userInput = stdin.readLineSync()!;
      try {
        int number = int.parse(userInput);
        if (validChoices != null && !validChoices.contains(number)) {
          throw FormatException('Veuillez saisir un nombre valide parmi les choix spécifiés.');
        }
        return number;
      } catch (e) {
        print('Erreur: Veuillez saisir un nombre valide.');
      }
    }
  }

  void cook(int duration, int temperature) {
    try {
      print('Cuisson en cours pendant $duration minute(s) à $temperature°C...');
      sleep(duration);
      print('Cuisson terminée!');
    } catch (error) {
      print('Erreur: $error');
    }
  }

  void boil() {
    try {
      print('Bouillir de l\'eau en cours...');
      const boilingTime = 10; // Durée de la simulation (10 minutes par exemple)
      int currentTemperature = 0;

      for (int i = 1; i <= boilingTime; i++) {
        currentTemperature += 10;
        sleep(1);
        print('Température de l\'eau: $currentTemperature°C - Durée restante: ${boilingTime - i} minutes');
      }

      print('L\'eau a atteint 100°C. Arrêt automatique.');
    } catch (error) {
      print('Erreur: $error');
    }
  }

  void sleep(int minutes) {
    final milliseconds = Duration(minutes: minutes).inMilliseconds;
    final startTime = DateTime.now().millisecondsSinceEpoch;
    while (DateTime.now().millisecondsSinceEpoch < startTime + milliseconds);
  }
}

void main() {
  final myRiceCooker = RiceCooker();

  stdout.write('Salut !! Vous êtes sur le rice cooker. Voulez-vous l\'allumer ? (1 pour oui / 2 pour non): ');
  int userDecision = myRiceCooker.getValidNumber('', [1, 2]);
  myRiceCooker.chooseFunction(userDecision.toString());
}
