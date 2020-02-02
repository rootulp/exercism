import re

def convert_headers(line):
    if re.match('###### (.*)', line) is not None:
        line = '<h6>' + line[7:] + '</h6>'
    elif re.match('## (.*)', line) is not None:
        line = '<h2>' + line[3:] + '</h2>'
    elif re.match('# (.*)', line) is not None:
        line = '<h1>' + line[2:] + '</h1>'
    return line

def convert_bold(item):
    is_bold = re.match('(.*)__(.*)__(.*)', item)
    if is_bold:
        return is_bold.group(1) + '<strong>' + \
            is_bold.group(2) + '</strong>' + is_bold.group(3)
    return item

def convert_italic(item):
    is_italic = re.match('(.*)_(.*)_(.*)', item)
    if is_italic:
        return is_italic.group(1) + '<em>' + is_italic.group(2) + \
            '</em>' + is_italic.group(3)
    return item

def parse_markdown(markdown):
    result = ''
    in_list = False
    for line in markdown.split('\n'):
        line = convert_headers(line)
        list_item = re.match(r'\* (.*)', line)
        if list_item:
            item = list_item.group(1)
            if not in_list:
                in_list = True
                item = convert_bold(item)
                item = convert_italic(item)
                line = '<ul><li>' + item + '</li>'
            else:
                item = convert_bold(item)
                item = convert_italic(item)
                line = '<li>' + item + '</li>'
        else:
            if in_list:
                line = '</ul>+i'
                in_list = False

        list_item = re.match('<h|<ul|<p|<li', line)
        if not list_item:
            line = '<p>' + line + '</p>'
        list_item = re.match('(.*)__(.*)__(.*)', line)
        if list_item:
            line = list_item.group(1) + '<strong>' + list_item.group(2) + '</strong>' + list_item.group(3)
        list_item = re.match('(.*)_(.*)_(.*)', line)
        if list_item:
            line = list_item.group(1) + '<em>' + list_item.group(2) + '</em>' + list_item.group(3)
        result += line

    if in_list:
        result += '</ul>'
    return result
