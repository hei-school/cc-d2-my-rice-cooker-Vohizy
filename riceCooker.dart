import 'dart:io';

class RiceCooker {
  bool power = false;

  bool checkPower() {
    return power;
  }

  void turnOn() {
    power = true;
    print("Hello!! You are on the rice cooker. The rice cooker is turned on.");
  }

  void turnOff() {
    power = false;
    print("The rice cooker is turned off.");
  }

  void chooseFunction(String userInput) {
    switch (userInput) {
      case '1':
        turnOn();
        try {
          int action = getValidNumber(
              'What do you want to do? (1 to cook / 2 to boil): ', [1, 2]);
          switch (action) {
            case 1:
              int duration = getValidNumber('How many minutes do you want to cook? ', null);
              int temperature = getValidNumber('At what temperature do you want to cook? ', null);
              cook(duration, temperature);
              break;
            case 2:
              boil();
              break;
            default:
              print('Unrecognized action.');
          }
        } catch (error) {
          print('Error: $error');
        }
        break;
      case '2':
        turnOff();
        break;
      default:
        print('Unrecognized command.');
    }
  }

  int getValidNumber(String promptMessage, List<int>? validChoices) {
    while (true) {
      stdout.write(promptMessage);
      String userInput = stdin.readLineSync()!;
      try {
        int number = int.parse(userInput);
        if (validChoices != null && !validChoices.contains(number)) {
          throw FormatException('Please enter a valid number among the specified choices.');
        }
        return number;
      } catch (e) {
        print('Error: Please enter a valid number.');
      }
    }
  }

  void cook(int duration, int temperature) {
    try {
      print('Cooking in progress for $duration minute(s) at $temperature°C...');
      sleep(duration);
      print('Cooking completed!');
    } catch (error) {
      print('Error: $error');
    }
  }

  void boil() {
    try {
      print('Boiling water in progress...');
      const boilingTime = 10; // Simulation duration (10 minutes for example)
      int currentTemperature = 0;

      for (int i = 1; i <= boilingTime; i++) {
        currentTemperature += 10;
        sleep(1);
        print('Water temperature: $currentTemperature°C - Remaining time: ${boilingTime - i} minutes');
      }

      print('Water has reached 100°C. Automatic shutdown.');
    } catch (error) {
      print('Error: $error');
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

  stdout.write('Hello! You are on the rice cooker. Do you want to turn it on? (1 for yes / 2 for no): ');
  int userDecision = myRiceCooker.getValidNumber('', [1, 2]);
  myRiceCooker.chooseFunction(userDecision.toString());
}
