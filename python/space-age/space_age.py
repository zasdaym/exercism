def secondToYear(seconds):
    return seconds / 31557600.0

class SpaceAge:
    def __init__(self, seconds):
        self.age = seconds

    def on_earth(self):
        result = secondToYear(self.age) / 1.0
        return float(format(result, '.2f'))

    def on_mercury(self):
        result = secondToYear(self.age) / 0.2408467
        return float(format(result, '.2f'))

    def on_venus(self):
        result = secondToYear(self.age) / 0.61519726
        return float(format(result, '.2f'))

    def on_mars(self):
        result = secondToYear(self.age) / 1.8808158
        return float(format(result, '.2f'))

    def on_jupiter(self):
        result = secondToYear(self.age) / 11.862615
        return float(format(result, '.2f'))

    def on_saturn(self):
        result = secondToYear(self.age) / 29.447498
        return float(format(result, '.2f'))

    def on_uranus(self):
        result = secondToYear(self.age) / 84.016846
        return float(format(result, '.2f'))

    def on_neptune(self):
        result = secondToYear(self.age) / 164.79132
        return float(format(result, '.2f'))

