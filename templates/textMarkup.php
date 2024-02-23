<?php

trait __fRead_textMarkup{
    const ver = "1.0";

    const textAlignLeft =       ["::BL", "::EL"];
    const textAlignRight =      ["::BR", "::ER"];
    const textAlignCenter =     ["::BC", "::EC"];

    const textStyleThroughline =    ["::BS", "::ES"];
    const textStyleUnderline =      ["::BU", "::EU"];
    const textStyleItalic =         ["::BI", "::EI"];
    const textStyleBold =           ["::BB", "::EB"];

    const del = ["\r", "\0", "\t"];
}