import re

def convert_headers(line):
    if re.match('###### (.*)', line) is not None:
        line = '<h6>' + line[7:] + '</h6>'
    elif re.match('## (.*)', line) is not None:
        line = '<h2>' + line[3:] + '</h2>'
    elif re.match('# (.*)', line) is not None:
        line = '<h1>' + line[2:] + '</h1>'
    return line


def parse_markdown(markdown):
    result = ''
    in_list = False
    for line in markdown.split('\n'):
        line = convert_headers(line)
        m = re.match(r'\* (.*)', line)
        if m:
            is_bold = False
            is_italic = False
            curr = m.group(1)
            m1 = re.match('(.*)__(.*)__(.*)', curr)
            if not in_list:
                in_list = True
                if m1:
                    is_bold = True
                    curr = m1.group(1) + '<strong>' + \
                        m1.group(2) + '</strong>' + m1.group(3)
                m1 = re.match('(.*)_(.*)_(.*)', curr)
                if m1:
                    curr = m1.group(1) + '<em>' + m1.group(2) + \
                        '</em>' + m1.group(3)
                    is_italic = True
                line = '<ul><li>' + curr + '</li>'
            else:
                if m1:
                    is_bold = True
                m1 = re.match('(.*)_(.*)_(.*)', curr)
                if m1:
                    is_italic = True
                if is_bold:
                    curr = m1.group(1) + '<strong>' + \
                        m1.group(2) + '</strong>' + m1.group(3)
                if is_italic:
                    curr = m1.group(1) + '<em>' + m1.group(2) + \
                        '</em>' + m1.group(3)
                line = '<li>' + curr + '</li>'
        else:
            if in_list:
                line = '</ul>+i'
                in_list = False

        m = re.match('<h|<ul|<p|<li', line)
        if not m:
            line = '<p>' + line + '</p>'
        m = re.match('(.*)__(.*)__(.*)', line)
        if m:
            line = m.group(1) + '<strong>' + m.group(2) + '</strong>' + m.group(3)
        m = re.match('(.*)_(.*)_(.*)', line)
        if m:
            line = m.group(1) + '<em>' + m.group(2) + '</em>' + m.group(3)
        result += line

    if in_list:
        result += '</ul>'
    return result
