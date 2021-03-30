from typing import Any, Tuple, Optional

def handle_error_by_throwing_exception() -> None:
    raise Exception("Exception!")


def handle_error_by_returning_none(input_data: Any) -> Optional[int]:
    try:
        return int(input_data)
    except ValueError:
        return None


def handle_error_by_returning_tuple(input_data) -> Tuple[bool, Optional[int]]:
    try:
        return True, int(input_data)
    except ValueError:
        return False, None


def filelike_objects_are_closed_on_exception(filelike_object: Any) -> None:
    with filelike_object as file:
        file.do_something()
