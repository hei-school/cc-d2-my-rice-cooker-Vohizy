import time


class RiceCooker:
    def __init__(self):
        self.power = False

    def check_power(self):
        return self.power

    def turn_on(self):
        self.power = True
        print("The rice cooker is turned on.")

    def turn_off(self):
        self.power = False
        print("The rice cooker is turned off.")

    def choose_function(self, user_input):
        if user_input == "1":
            self.turn_on()
            try:
                action = self.get_valid_number("What do you want to do? (1 to cook / 2 to boil): ", [1, 2])
                if action == 1:
                    duration = self.get_valid_number("How many minutes do you want to cook? ")
                    temperature = self.get_valid_number("At what temperature do you want to cook? ")
                    self.cook(duration, temperature)
                elif action == 2:
                    self.boil()
                else:
                    print("Unrecognized action.")
            except ValueError as e:
                print(f"Error: {e}")
        elif user_input == "2":
            self.turn_off()
        else:
            print("Unrecognized command.")

    def get_valid_number(self, prompt_message, valid_choices=None):
        while True:
            try:
                user_input = input(prompt_message)
                if not user_input.isdigit():
                    raise ValueError("Please enter a number.")
                user_input = int(user_input)

                if valid_choices is not None and user_input not in valid_choices:
                    raise ValueError("Please enter a valid number among the specified choices.")

                return user_input
            except ValueError as e:
                print(f"Error: {e}")

    def cook(self, duration, temperature):
        try:
            print(f"Cooking in progress for {duration} minute(s) at {temperature}°C...")
            self.sleep(duration)
            print("Cooking completed!")
        except ValueError as e:
            print(f"Error: {e}")

    def boil(self):
        try:
            print("Boiling water in progress...")
            boiling_time = 10  # Simulation duration (10 minutes, for example)
            current_temperature = 0

            for i in range(1, boiling_time + 1):
                current_temperature += 10
                self.sleep(1)
                print(f"Water temperature: {current_temperature}°C - Remaining time: {boiling_time - i} minutes")

            print("Water has reached 100°C. Automatic shutdown.")
        except ValueError as e:
            print(f"Error: {e}")

    def sleep(self, minutes):
        milliseconds = minutes * 60 * 1000
        start_time = int(time.time() * 1000)
        while int(time.time() * 1000) < start_time + milliseconds:
            pass


# Example usage
my_rice_cooker = RiceCooker()

try:
    user_decision = my_rice_cooker.get_valid_number("Hello! You are using the rice cooker. Do you want to turn it on? (1 for yes / 2 for no): ", [1, 2])
    my_rice_cooker.choose_function(str(user_decision))
except ValueError as e:
    print(f"Error: {e}")
