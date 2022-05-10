use std::default;

use enum_iterator::IntoEnumIterator;
use int_enum::IntEnum;

#[repr(usize)]
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
    return _color.int_value()
}

pub fn value_to_color_string(value: usize) -> String {
    match value {
        0 => "Black".to_string(),
        1 => "Brown".to_string(),
        2 => "Red".to_string(),
        3 => "Orange".to_string(),
        4 => "Yellow".to_string(),
        5 => "Green".to_string(),
        6 => "Blue".to_string(),
        7 => "Violet".to_string(),
        8 => "Grey".to_string(),
        9 => "White".to_string(),
        _ => "value out of range".to_string()
    }
}

pub fn colors() -> Vec<ResistorColor> {
    unimplemented!("return a list of all the colors ordered by resistance")
}
