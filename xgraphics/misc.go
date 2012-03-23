/*
    A collection of miscellaneous methods that generate images useful for
    window decoration.

    I think this will get moved into the Wingo repo.
*/
package xgraphics

import (
    "image"
    "image/color"
)

const (
    BorderTop = iota
    BorderRight
    BorderBottom
    BorderLeft
)

// ColorFromHex takes a hex number in the format of '0xrrggbb' and transforms 
// it to an RGBA color.
func ColorFromStr(clr int) color.RGBA {
    r := clr >> 16
    g := (clr - (r << 16)) >> 8
    b := clr - ((clr >> 8) << 8)
    return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}
}

func Border(borderType int, borderColor int, bgColor int,
            width, height int) image.Image {
    bgClr := ColorFromStr(bgColor)
    borderClr := ColorFromStr(borderColor)

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch {
            case borderType == BorderTop && y == 0:
                img.SetRGBA(x, y, borderClr)
            case borderType == BorderRight && x == width - 1:
                img.SetRGBA(x, y, borderClr)
            case borderType == BorderBottom && y == height - 1:
                img.SetRGBA(x, y, borderClr)
            case borderType == BorderLeft && x == 0:
                img.SetRGBA(x, y, borderClr)
            default:
                img.SetRGBA(x, y, bgClr)
            }
        }
    }

    return img
}
