# `gradient`

Package gradient provides capabilities for generating gradients of various types.

```go
package main

import (
    "image"
    "image/png"
    "os"

    "github.com/shelepuginivan/color"
    "github.com/shelepuginivan/color/gradient"
)

func main() {
    g, _ := gradient.NewLinear(
        gradient.WithAngle(45),
        gradient.InOklch(gradient.LongerHue),
        gradient.WithColorStop(color.Must(color.ParseHex("#a29c6b")), 0),
        gradient.WithColorStop(color.Must(color.ParseHex("#99a286")), 1),
    )

    img := image.NewRGBA(image.Rect(0, 0, 3840, 2160))

    g.Render(img)

    output, err := os.Create("gradient.png")
    if err != nil {
        panic(err)
    }
    defer output.Close()

    if err := png.Encode(output, img); err != nil {
        panic(err)
    }
}
```

## Gradient types

The following gradient types are implemented:

- Linear — colors transition along a straight line
- Radial — colors transition from center to edges, forming circular shapes
- Conic — colors transition rotating around the center
- Diamond — colors transition from center to edges, forming diamond shapes

| Gradient type |                                                                Sample                                                                |
| :------------ | :----------------------------------------------------------------------------------------------------------------------------------: |
| Linear        | <img width="512" height="128" alt="linear" src="https://github.com/user-attachments/assets/73aa89fe-86be-4dcf-a95e-767ee2e56cb1" />  |
| Radial        | <img width="512" height="128" alt="radial" src="https://github.com/user-attachments/assets/d303ec69-75a9-4374-a84b-dda408b0fb0d" />  |
| Conic         | <img width="512" height="128" alt="conic" src="https://github.com/user-attachments/assets/e1b52697-9fb9-4e94-bd2f-a3db0480462b" />   |
| Diamond       | <img width="512" height="128" alt="diamong" src="https://github.com/user-attachments/assets/88df6d72-b4fd-481f-9817-67a7c454a9e4" /> |

## Colorspaces

The following colorspaces are supported

- RGB
- HSL (with hue interpolation method)
- HSV (with hue interpolation method)
- XYZ (with reference white)
- Lab (with reference white)
- Lch (with reference white and hue interpolation method)
- Oklab
- Oklch

| Colorspace          |                                                                      Sample                                                                      |
| :------------------ | :----------------------------------------------------------------------------------------------------------------------------------------------: |
| RGB                 | <img width="512" height="128" alt="rgb" src="https://github.com/user-attachments/assets/2c38f1a3-1de4-4a23-b111-2768668e10ed" />                 |
| HSL shorter hue     | <img width="512" height="128" alt="hsl-shorter-hue" src="https://github.com/user-attachments/assets/acc771ec-b008-4ed9-9edb-6b1b885c12b8" />     |
| HSL longer hue      | <img width="512" height="128" alt="hsl-longer-hue" src="https://github.com/user-attachments/assets/e094b21d-74d3-4956-a373-2da8912495c4" />      |
| HSV shorter hue     | <img width="512" height="128" alt="hsv-shorter-hue" src="https://github.com/user-attachments/assets/c927c4da-cde6-43bf-a9d1-6fedab8255d5" />     |
| HSV longer hue      | <img width="512" height="128" alt="hsv-longer-hue" src="https://github.com/user-attachments/assets/33cdd604-e6e1-431b-9b0a-903d106e4559" />      |
| XYZ D50             | <img width="512" height="128" alt="xyz-d50" src="https://github.com/user-attachments/assets/c860af1a-cc49-4c4f-96d6-7876aab84bce" />             |
| XYZ D65             | <img width="512" height="128" alt="xyz-d65" src="https://github.com/user-attachments/assets/6ce8519c-9f28-4f2a-995c-0700d7b7c20f" />             |
| Lab D50             | <img width="512" height="128" alt="lab-d50" src="https://github.com/user-attachments/assets/f0e82e8d-91c6-47f3-8b07-a42e57667352" />             |
| Lab D65             | <img width="512" height="128" alt="lab-d65" src="https://github.com/user-attachments/assets/7de03aa5-b1f0-4dc7-aa5c-35f6b517cf61" />             |
| Lch shorter hue D50 | <img width="512" height="128" alt="lch-shorter-hue-d50" src="https://github.com/user-attachments/assets/2185c2e4-a2ca-476b-b401-8573fc3c1777" /> |
| Lch longer hue D50  | <img width="512" height="128" alt="lch-longer-hue-d50" src="https://github.com/user-attachments/assets/c3a0d4fa-1a5b-47c2-b578-a491cfb62aab" />  |
| Lch shorter hue D65 | <img width="512" height="128" alt="lch-shorter-hue-d65" src="https://github.com/user-attachments/assets/2597e1b7-6ce1-4278-9a23-c824e4e443dd" /> |
| Lch longer hue D65  | <img width="512" height="128" alt="lch-longer-hue-d65" src="https://github.com/user-attachments/assets/4234fd73-79b3-400a-9c38-f605f393146c" />  |
| Oklab               | <img width="512" height="128" alt="oklab" src="https://github.com/user-attachments/assets/85e5446a-bb7b-48b1-987e-7dd4dc7d11a3" />               |
| Oklch shorter hue   | <img width="512" height="128" alt="oklch-shorter-hue" src="https://github.com/user-attachments/assets/620f4e3a-f301-4093-ac8d-4aa2c26796bc" />   |
| Oklch longer hue    | <img width="512" height="128" alt="oklch-longer-hue" src="https://github.com/user-attachments/assets/57e4a672-1224-49e6-9700-98d7b0e5294a" />    |
