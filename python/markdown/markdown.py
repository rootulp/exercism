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

def parse_markdown(markdown):
    result = ''
    in_list = False
    for line in markdown.split('\n'):
        line = convert_headers(line)
        line = convert_bold(line)
        line = convert_italic(line)
        list_item = re.match(r'\* (.*)', line)
        if list_item:
            item = list_item.group(1)
            item = convert_bold(item)
            item = convert_italic(item)
            if not in_list:
                in_list = True
                line = '<ul><li>' + item + '</li>'
            else:
                line = '<li>' + item + '</li>'
        else:
            if in_list:
                line = '</ul>+i'
                in_list = False

        list_item = re.match('<h|<ul|<p|<li', line)
        if not list_item:
            line = '<p>' + line + '</p>'
        result += line

    if in_list:
        result += '</ul>'
    return result
