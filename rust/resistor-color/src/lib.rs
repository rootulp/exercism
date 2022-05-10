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
    for color in ResistorColor::into_enum_iter() {
        if value == color.int_value() {
            return label(color)
        }
    }

    return "value out of range".to_string()
}

pub fn colors() -> Vec<ResistorColor> {
    unimplemented!("return a list of all the colors ordered by resistance")
}

fn label(color: ResistorColor) -> String {
    match color {
        ResistorColor::Black => "Black".to_string(),
        ResistorColor::Brown => "Brown".to_string(),
        ResistorColor::Red => "Red".to_string(),
        ResistorColor::Orange => "Orange".to_string(),
        ResistorColor::Yellow => "Yellow".to_string(),
        ResistorColor::Green => "Green".to_string(),
        ResistorColor::Blue => "Blue".to_string(),
        ResistorColor::Violet => "Violet".to_string(),
        ResistorColor::Grey => "Grey".to_string(),
        ResistorColor::White => "White".to_string(),
    }
}
