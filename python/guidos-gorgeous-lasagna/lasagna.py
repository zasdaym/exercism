"""Functions used in preparing Guido's gorgeus lasagna."""

EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

def bake_time_remaining(elapsed_bake_time):
    """Return remaining bake time."""
    return EXPECTED_BAKE_TIME - elapsed_bake_time


def preparation_time_in_minutes(layers):
    """Return total preparation time for given number of layers."""
    return PREPARATION_TIME * layers


def elapsed_time_in_minutes(layers, elapsed_bake_time):
    """
    Return elapsed cooking time.

    This function takes two numbers representing the number of layers & the time already spent
    baking and calculates the total elapsed minutes spent cooking the lasagna.
    """
    return preparation_time_in_minutes(layers) + elapsed_bake_time
