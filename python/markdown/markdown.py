import re


# TODO Extract logic for parsing unordered lists.
def parse_markdown(markdown):
    result = ''
    in_list = False
    for line in markdown.split('\n'):
        line = convert_headers(line)
        line = convert_item(line)

        list_item = re.match(r'\* (.*)', line)
        if list_item:
            item = list_item.group(1)
            line = convert_item(line)

            if in_list:
                line = '<li>' + item + '</li>'
            else:
                in_list = True
                line = '<ul><li>' + item + '</li>'
        else:
            if in_list:
                in_list = False
                line = '</ul>'

        line = maybe_wrap_line_in_paragraph(line)

        result += line

    if in_list:
        result += '</ul>'
    return result


def convert_headers(line):
    if re.match('###### (.*)', line) is not None:
        return '<h6>' + line[7:] + '</h6>'
    elif re.match('## (.*)', line) is not None:
        return '<h2>' + line[3:] + '</h2>'
    elif re.match('# (.*)', line) is not None:
        return '<h1>' + line[2:] + '</h1>'
    return line


def convert_item(item):
    return convert_italic(convert_bold(item))


def convert_bold(item):
    bold_item = re.match('(.*)__(.*)__(.*)', item)
    if bold_item:
        return bold_item.group(1) + '<strong>' + \
            bold_item.group(2) + '</strong>' + bold_item.group(3)
    return item


def convert_italic(item):
    italic_item = re.match('(.*)_(.*)_(.*)', item)
    if italic_item:
        return italic_item.group(1) + '<em>' + italic_item.group(2) + \
            '</em>' + italic_item.group(3)
    return item


def maybe_wrap_line_in_paragraph(line):
    any_surrounding_tag = re.match('<h|<ul|<p|<li', line)
    if any_surrounding_tag:
        return line
    return '<p>' + line + '</p>'
