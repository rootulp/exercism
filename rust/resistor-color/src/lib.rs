use enum_iterator::IntoEnumIterator;
use int_enum::IntEnum;

#[repr(u8)]
#[derive(Clone, Copy, Debug, IntoEnumIterator, IntEnum, PartialEq)]
pub enum ResistorColor {
    Black = 0,
    Brown = 1,
    Red = 2,
    Orange = 3,
    Yellow = 4,
    Green = 5,
    Blue = 6,
    Violet = 7,
    Grey = 8,
    White = 9,
}

pub fn color_to_value(_color: ResistorColor) -> usize {
    return _color.int_value().into()
}

pub fn value_to_color_string(value: usize) -> String {
    unimplemented!(
        "convert the value {} into a string representation of color",
        value
    )
}

pub fn colors() -> Vec<ResistorColor> {
    unimplemented!("return a list of all the colors ordered by resistance")
}
