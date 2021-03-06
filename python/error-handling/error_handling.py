def handle_error_by_throwing_exception():
    raise Exception("Failed while doing something")


def handle_error_by_returning_none(input_data):
    if (input_data == '1'):
        return 1
    return None


def handle_error_by_returning_tuple(input_data):
    if (input_data == '1'):
        return (True, 1)
    return (False, None)


def filelike_objects_are_closed_on_exception(filelike_object):
    with filelike_object as f:
        f.do_something()
