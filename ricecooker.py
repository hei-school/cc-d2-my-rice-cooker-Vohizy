import time

class RiceCooker:
    def __init__(self):
        self.power = False

    def check_power(self):
        return self.power

    def turn_on(self):
        self.power = True
        print("Le rice cooker est allumé.")

    def turn_off(self):
        self.power = False
        print("Le rice cooker est éteint.")

    def choose_function(self, user_input):
        if user_input == "1":
            self.turn_on()
            try:
                action = self.get_valid_number("Que voulez-vous faire ? (1 pour cuire / 2 pour bouillir): ", [1, 2])
                if action == 1:
                    duration = self.get_valid_number("Combien de minutes souhaitez-vous cuire ? ")
                    temperature = self.get_valid_number("À quelle température souhaitez-vous cuire ? ")
                    self.cook(duration, temperature)
                elif action == 2:
                    self.boil()
                else:
                    print("Action non reconnue.")
            except ValueError as e:
                print(f"Erreur: {e}")
        elif user_input == "2":
            self.turn_off()
        else:
            print("Commande non reconnue.")

    def get_valid_number(self, prompt_message, valid_choices=None):
        while True:
            try:
                user_input = input(prompt_message)
                if not user_input.isdigit():
                    raise ValueError("Veuillez saisir un nombre.")
                user_input = int(user_input)

                if valid_choices is not None and user_input not in valid_choices:
                    raise ValueError("Veuillez saisir un nombre valide parmi les choix spécifiés.")
                    
                return user_input
            except ValueError as e:
                print(f"Erreur: {e}")
    def cook(self, duration, temperature):
        try:
            print(f"Cuisson en cours pendant {duration} minute(s) à {temperature}°C...")
            self.sleep(duration)
            print("Cuisson terminée!")
        except ValueError as e:
            print(f"Erreur: {e}")

    def boil(self):
        try:
            print("Bouillir de l'eau en cours...")
            boiling_time = 10 # Durée de la simulation (10 minutes par exemple)
            current_temperature = 0

            for i in range(1, boiling_time + 1):
                current_temperature += 10
                self.sleep(1)
                print(f"Température de l'eau: {current_temperature}°C - Durée restante: {boiling_time - i} minutes")

            print("L'eau a atteint 100°C. Arrêt automatique.")
        except ValueError as e:
            print(f"Erreur: {e}")

    def sleep(self, minutes):
        milliseconds = minutes * 60 * 1000
        start_time = int(time.time() * 1000)
        while int(time.time() * 1000) < start_time + milliseconds:
            pass

# Exemple d'utilisation
my_rice_cooker = RiceCooker()

try:
    user_decision = my_rice_cooker.get_valid_number("Salut !! Vous êtes sur le rice cooker. Voulez-vous l'allumer ? (1 pour oui / 2 pour non): ", [1, 2])
    my_rice_cooker.choose_function(str(user_decision))
except ValueError as e:
    print(f"Erreur: {e}")
